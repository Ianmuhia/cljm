package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/models"
	"maranatha_web/internal/services"
	"maranatha_web/internal/utils"
	"maranatha_web/internal/utils/crypto_utils"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID                int    `json:"ID"`
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
		Username: user.UserName,
		FullName: user.FullName,
		Email:    user.Email,
		ID:       int(user.ID),
	}
}

//RegisterUser new user
func (r *Repository) RegisterUser(c *gin.Context) {

	var registerModel createUserRequest

	if err := c.ShouldBindJSON(&registerModel); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}

	user := models.User{
		UserName:     registerModel.Username,
		FullName:     registerModel.FullName,
		Email:        registerModel.Email,
		PasswordHash: registerModel.Password,
	}
	saveErr := r.userServices.CreateUser(user)

	if saveErr != nil {
		err := errors.NewBadRequestError("Could not save user.")
		c.JSON(err.Status, saveErr)
		c.Abort()
		return
	}
	log.Println("passed here")
	type mail services.Mail
	code := utils.GenerateRandomExpiryCode(user.Email)

	m := mail{
		To:      user.Email,
		From:    "me@here.com",
		Subject: "hello",
		Content: code,
	}
	err := services.MailService.SendMsg(services.Mail(m))
	//log.Println(&dc)

	if err != nil {
		log.Println(err)
		//logger.Info("could not send email ")
		return
	}
	message := fmt.Sprintf("Thank %s you for creating and account.Please verify your email %s code is %s", user.UserName, user.Email, code)

	response := registerUserResponse{
		Message:  message,
		EmailUrl: "localhost:8090/api/users/",
	}
	c.JSON(http.StatusOK, response)
}

//
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
func (r *Repository) Login(ctx *gin.Context) {

	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := r.userServices.GetUserByEmail(req.Email)
	if user == nil {
		//log.Println(user)
		data := errors.NewBadRequestError("The user does not exist.Please create an account to continue")
		ctx.JSON(data.Status, data)
		return
	}

	if !user.IsVerified {
		data := errors.NewBadRequestError("Please verify your email address to login")
		ctx.JSON(data.Status, data)
		return

	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ok := crypto_utils.CheckPasswordHash(req.Password, user.PasswordHash)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.NewBadRequestError("invalid email or password "))
		return
	}
	duration := 20 * time.Hour
	fmt.Println(duration.Hours())

	accessToken, erro := token.TokenService.CreateToken(user.Email, duration, user.ID)

	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, erro)
		return
	}
	duration = time.Duration(time.Now().Add(time.Hour * 20).Unix())

	refreshToken, erro := token.TokenService.CreateRefreshToken(user.Email, duration, user.ID)

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

func (r *Repository) VerifyEmailCode(ctx *gin.Context) {
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
		err := r.userServices.UpdateUserStatus(req.Email)

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

func TryAuthMiddlewareMiddleware(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello welcome")
}

func (r *Repository) GetAllUsers(ctx *gin.Context) {
	data, err := r.userServices.GetAllUsers()
	if err != nil {
		restErr := errors.NewBadRequestError("Error getting all users.")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (r *Repository) UpdateUserProfileImage(ctx *gin.Context) {
	data := r.GetPayloadFromContext(ctx)
	file, m, err := ctx.Request.FormFile("profile_image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	fmt.Println(file, m.Header, m.Filename, m.Size)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	log.Println(uploadFile)

	err = r.userServices.UpdateUserImage(data.Username, uploadFile.Key)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing upload profile image request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, "Profile image upload successful")

}
