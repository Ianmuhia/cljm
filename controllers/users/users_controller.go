package users_controller

import (
	"fmt"
	"net/http"
	"time"

	"maranatha_web/controllers/token"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils"
	"maranatha_web/utils/crypto_utils"
	"maranatha_web/utils/errors"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
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

type registerUserResponse struct {
	Message  string `json:"data"`
	EmailUrl string `json:"email-url"`
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

//RegisterUser new user
func RegisterUser(c *gin.Context) {

	var registerModel createUserRequest

	if err := c.ShouldBindJSON(&registerModel); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}

	user := models.User{
		UserName: registerModel.Username,
		FullName: registerModel.FullName,
		Email:    registerModel.Email,
		Password: registerModel.Password,
	}
	fmt.Println(registerModel)
	fmt.Println(user)

	result, saveErr := services.UsersService.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		c.Abort()
		return
	}
	type mail services.Mail
	code := utils.GenerateRandomExpiryCode(result.Email)

	m := mail{
		To:      result.Email,
		From:    "me@here.com",
		Subject: "hello",
		Content: code,
	}
	err := services.MailService.SendMsg(services.Mail(m))

	if err != nil {
		logger.Info("could not send email ")
		return
	}
	message := fmt.Sprintf("Thank %s you for creating and account.Please verify your email %s ", result.UserName, result.Email)

	response := registerUserResponse{
		Message:  message,
		EmailUrl: "localhost:8090/api/users/",
	}
	c.JSON(http.StatusOK, response)
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         userResponse `json:"user"`
}

// Login controller
func Login(ctx *gin.Context) {

	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := services.UsersService.GetUserByEmail(req.Email)

	if !user.IsVerified {
		data := errors.NewBadRequestError("Please verify your email address to login")
		ctx.JSON(http.StatusInternalServerError, data)
		return

	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ok := crypto_utils.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.NewBadRequestError("invalid email or password "))
		return
	}
	duration := 20 * time.Minute
	fmt.Println(duration.Minutes())

	accessToken, erro := token.TokenService.CreateToken(req.Email, duration)

	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, erro)
		return
	}
	duration = time.Duration(time.Now().Add(time.Minute * 20).Unix())

	refreshToken, erro := token.TokenService.CreateRefreshToken(req.Email, duration)

	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, erro)
		return
	}
	response := loginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}

type verifyEmailRequest struct {
	Code  string `json:"code" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type verifyEmailResponse struct {
	Message string `json:"message"`
}

func VerifyEmailCode(ctx *gin.Context) {
	var req verifyEmailRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	data := services.MailService.VerifyMailCode(req.Email)

	if req.Code == data {

		services.MailService.RemoveMailCode(req.Email)
		err := services.UsersService.UpdateUserStatus(req.Email)

		if err != nil {
			message := verifyEmailResponse{
				Message: "Error Updating user status",
			}
			ctx.JSON(http.StatusInternalServerError, message)
			ctx.Abort()
			return
		}
		message := verifyEmailResponse{
			Message: "Email has been verified you can now login to your account",
		}
		ctx.JSON(http.StatusOK, message)
		ctx.Abort()
		return
	}
	message := verifyEmailResponse{
		Message: "Email verification failed invalid code provided",
	}
	ctx.JSON(http.StatusBadRequest, message)

}

type resetEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type resetEmailResponse struct {
	Message string `json:"message"`
}

func ResetUserPassword(ctx *gin.Context) {

}

func TryAuthMiddlewareMiddleware(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello welcome")
}
