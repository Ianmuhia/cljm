package users_controller

import (
	"fmt"
	"log"
	"maranatha_web/config"
	"maranatha_web/controllers/token"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils"
	"maranatha_web/utils/crypto_utils"
	"maranatha_web/utils/errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string `json:"username"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PasswordChangedAt string `json:"password_changed_at"`
	CreatedAt         string `json:"created_at"`
}

func newUserResponse(user *models.User) userResponse {
	return userResponse{
		Username:          user.UserName,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.UpdatedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func getUserId(userIdParams string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParams, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userId, nil
}

var jwtKey = []byte("secret")

// var m *dbrepo.PostgresDBRepo

//Claims jwt claims struct
type Claims struct {
	models.User
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
		// /password := models.CreateHashedPassword(resetPassword.Password)
		// _, err := m.DB.Query(dbrepo.UpdateUserPasswordQuery, resetPassword.ID, password)
		// errors.NewNotFoundError(err)
		// errors.HandleErr(c, err)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User password reset successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "errors": errStr})
	}

}

//RegisterUser new user
func RegisterUser(c *gin.Context) {

	var register_model models.User

	if err := c.ShouldBindJSON(&register_model); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(register_model)

	result, saveErr := services.UsersService.CreateUser(register_model)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)

}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

// Login controller
func Login(ctx *gin.Context) {

	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := services.UsersService.GetUser(req.Email)

	if err != nil {
		fmt.Println(err)
	}
	ok := crypto_utils.CheckPasswordHash(req.Password, user.Password)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.NewBadRequestError("email or password error"))
		return
	}

	var duration time.Duration

	var accessToken, erro = token.TokenService.CreateToken(req.Email, duration)
	if erro != nil {
		log.Println(erro)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(accessToken)

	response := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}

func checkUserExists(user models.Register) bool {
	// rows, err := m.DB.Query(dbrepo.CheckUserExists, user.Email)
	// if err != nil {
	// 	return false
	// }
	// if !rows.Next() {
	// 	return false
	// }
	return true
}

//Returns -1 as ID if the user doesnt exist in the table
func checkAndRetrieveUserIDViaEmail(createReset models.CreateReset) (int, bool) {
	// rows, err := m.DB.Query(dbrepo.CheckUserExists, createReset.Email)
	// if err != nil {
	// 	return -1, false
	// }
	// if !rows.Next() {
	// 	return -1, false
	// }
	// var id int
	// err = rows.Scan(&id)
	// if err != nil {
	// 	return -1, false
	// }
	return 1, true
}
