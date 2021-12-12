package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/wycemiro/cljm.git/internal/config"
	"bitbucket.org/wycemiro/cljm.git/internal/errors"
	"bitbucket.org/wycemiro/cljm.git/internal/models"
	"bitbucket.org/wycemiro/cljm.git/internal/repository/dbrepo"
	"bitbucket.org/wycemiro/cljm.git/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")
var m *dbrepo.PostgresDBRepo

//Claims jwt claims struct
type Claims struct {
	models.Users
	jwt.StandardClaims
}

// AuthMiddleware checks that token is valid, see https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func AuthMiddleware(c *gin.Context, jwtKey []byte) (jwt.MapClaims, bool) {
	//obtain session token from the requests cookies
	ck, err := c.Request.Cookie("token")
	fmt.Println(ck, "coookie")
	if err != nil {
		fmt.Print(err)
		return nil, false
	}

	// Get the JWT string from the cookie
	tokenString := ck.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}

//Initiate Password reset email with reset url
func InitiatePasswordReset(c *gin.Context) {
	var createReset models.CreateReset
	c.Bind(&createReset)
	if id, ok := checkAndRetrieveUserIDViaEmail(createReset); ok {
		link := fmt.Sprintf("%s/reset/%d", config.CLIENT_URL, id)
		//Reset link is returned in json response for testing purposes since no email service is integrated
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Successfully sent reset mail to " + createReset.Email, "link": link})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "errors": "No user found for email: " + createReset.Email})
	}
}

func ResetPassword(c *gin.Context) {
	var resetPassword models.ResetPassword
	c.Bind(&resetPassword)
	if ok, errStr := utils.ValidatePasswordReset(resetPassword); ok {
		password := models.CreateHashedPassword(resetPassword.Password)
		_, err := m.DB.Query(dbrepo.UpdateUserPasswordQuery, resetPassword.ID, password)
		errors.HandleErr(c, err)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User password reset successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "errors": errStr})
	}

}

//Create new user
func Create(c *gin.Context) {
	var user models.Register
	c.Bind(&user)
	exists := checkUserExists(user)

	valErr := utils.ValidateUser(user, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "email already exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	models.HashPassword(&user)
	_, err := m.DB.Query(dbrepo.CreateUserQuery, user.Name, user.Password, user.Email)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created succesfully"})
}

// Session returns JSON of user info
func Session(c *gin.Context) {
	user, isAuthenticated := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

// Login controller
func Login(c *gin.Context) {
	var user models.Login
	c.Bind(&user)

	row := m.DB.QueryRow(dbrepo.LoginQuery, user.Email)

	var id int
	var name, email, password string
	var createdAt, updatedAt time.Time

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	match := models.CheckPasswordHash(user.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{

		Users: models.Users{
			UserName: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "logged in succesfully", "user": claims.Users, "token": tokenString})
}

func checkUserExists(user models.Register) bool {
	rows, err := m.DB.Query(dbrepo.CheckUserExists, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

//Returns -1 as ID if the user doesnt exist in the table
func checkAndRetrieveUserIDViaEmail(createReset models.CreateReset) (int, bool) {
	rows, err := m.DB.Query(dbrepo.CheckUserExists, createReset.Email)
	if err != nil {
		return -1, false
	}
	if !rows.Next() {
		return -1, false
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, false
	}
	return id, true
}
