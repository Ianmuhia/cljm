package controllers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"maranatha_web/internal/controllers/token"
	redis_db "maranatha_web/internal/datasources/redis"
	"maranatha_web/internal/models"
	"maranatha_web/internal/services"
	"maranatha_web/internal/utils"
	"maranatha_web/internal/utils/crypto_utils"
	"maranatha_web/internal/utils/errors"
)

type createUserRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=6"`
	FullName        string `json:"full_name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
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
func (r *Repository) RegisterUser(ctx *gin.Context) {

	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	if req.Password != req.PasswordConfirm {
		r.App.ErrorLog.Error("password and password re-enter did not match")
		e := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return

	}
	newPassword := crypto_utils.Hash(req.Password)

	user := models.User{
		UserName:     req.Username,
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: newPassword,
	}
	saveErr := r.userServices.CreateUser(user)

	if saveErr != nil {
		err := errors.NewBadRequestError("Could not save user.")
		ctx.JSON(err.Status, saveErr)
		ctx.Abort()
		return
	}

	code := utils.GenerateRandomExpiryCode(user.Email)

	from := "me@here.com"
	to := user.Email
	subject := "Email Verification for Pass Change"
	//subject := "Email Verification for Pass Change"
	mailType := services.MailConfirmation
	mailData := &services.MailData{
		Username: user.UserName,
		Code:     code,
	}

	mailReq := r.mailService.NewMail(from, to, subject, mailType, mailData)

	err := r.mailService.SendMsg(mailReq)

	if err != nil {
		log.Println(err)
		return
	}
	message := fmt.Sprintf("Thank %s you for creating and account.Please verify your email %s code is %s", user.UserName, user.Email, code)

	response := registerUserResponse{
		Message:  message,
		EmailUrl: "localhost:8090/api/users/",
	}
	ctx.JSON(http.StatusOK, response)
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
	if err == gorm.ErrRecordNotFound {
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

	//t := r.App.TokenLifeTime
	duration := 10 * time.Hour
	//duration :=

	accessToken, err := token.TokenService.CreateToken(user.Email, duration, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	duration = time.Duration(time.Now().Add(time.Hour * 20).Unix())

	refreshToken, err := token.TokenService.CreateRefreshToken(user.Email, duration, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	//send fcm message
	err = r.messagingService.SendNotification()
	if err != nil {
		log.Println(err)
		//logger.Info("could not send email ")
		//return
	}
	response := loginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}

type verifyEmailRequest struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

type verifyEmailResponse struct {
	Message string `json:"message"`
}

func (r *Repository) VerifyEmailCode(ctx *gin.Context) {
	var req verifyEmailRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		log.Print(err)
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	data := r.mailService.VerifyMailCode(req.Email)

	if req.Code == data {

		r.mailService.RemoveMailCode(req.Email)
		err := r.userServices.UpdateUserStatus(req.Email)

		if err != nil {
			message := verifyEmailResponse{
				Message: "Error Updating user status",
			}
			ctx.JSON(http.StatusInternalServerError, message)
			ctx.Abort()
			return
		}
		resp := SuccessResponse{
			TimeStamp: time.Now(),
			Message:   "Email has been verified you can now login to your account",
			Status:    http.StatusOK,
			Data:      nil,
		}

		ctx.JSON(resp.Status, resp)
		ctx.Abort()
		return
	}

	resp := errors.NewBadRequestError("Email verification failed invalid code provided")
	log.Println(resp)
	ctx.JSON(resp.Status, resp)

}

type GetAllUsersResponse struct {
	Total int            `json:"total"`
	Users []*models.User `json:"users"`
}

func (r *Repository) GetAllUsers(ctx *gin.Context) {
	total, users, err := r.userServices.GetAllUsers()
	if err != nil {
		restErr := errors.NewBadRequestError("Error getting all users.")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	resp := GetAllUsersResponse{
		Total: total,
		Users: users,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (r *Repository) GetUser(ctx *gin.Context) {
	user := r.GetPayloadFromContext(ctx)

	data, err := r.userServices.GetUserByID(user.ID)
	if err != nil {
		restErr := errors.NewNotFoundError("user does not exits")
		log.Print(err)
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	res := NewStatusOkResponse("Successfully got user", data)
	ctx.JSON(res.Status, res)
}
func (r *Repository) DeleteUser(ctx *gin.Context) {
	r.GetPayloadFromContext(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := r.userServices.DeleteUser(uint(id))
	if err != nil {
		restErr := errors.NewNotFoundError("user does not exits")
		log.Print(err)
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	res := NewStatusOkResponse("Successfully got user", nil)
	ctx.JSON(res.Status, res)
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
	//TODO:Add response model
	ctx.JSON(http.StatusCreated, "Profile image upload successful")

}

type GetPasswordResetCode struct {
	Email string `json:"email"`
}

func (r *Repository) ForgotPassword(ctx *gin.Context) {
	var req GetPasswordResetCode

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	user, err := r.userServices.GetUserByEmail(req.Email)
	if err != nil {
		restErr := errors.NewBadRequestError("User with that email does not exits")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}

	// Send verification mail
	from := "me@gmail.com"
	to := user.Email
	subject := "Password Reset for User"
	//body := "Password Reset for User"
	mailType := services.PassReset
	mailData := &services.MailData{
		Username: user.UserName,
		Code:     utils.GenerateRandomExpiryCode(user.UserName),
	}

	//mailReq := r.mailService.NewMail(from, to, subject, body, mailType, mailData)
	mailReq := &services.Mail{
		From:     from,
		To:       to,
		Subject:  subject,
		Body:     mailData,
		MailType: mailType,
	}

	log.Println(mailReq)
	err = r.mailService.SendMsg(mailReq)
	if err != nil {
		restErr := errors.NewBadRequestError("Unable to send mail.")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	// store the password reset code to db
	verificationData := &services.VerificationData{
		Email:     user.Email,
		Code:      mailData.Code,
		Type:      string(rune(services.PassReset)),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(r.App.PasswordResetCodeExpiry)),
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(verificationData); err != nil {
		log.Println(err)
		restErr := errors.NewBadRequestError("Unable to send password reset code. Please try again later")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	err = redis_db.RedisClient.Set(ctx, verificationData.Email, b.Bytes(), time.Minute*time.Duration(r.App.PasswordResetCodeExpiry)).Err()
	if err != nil {
		log.Println(err)
		restErr := errors.NewBadRequestError("Unable to send password reset code. Please try again later")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	data := SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Password reset code sent successfully",
		Status:    http.StatusOK,
		Data:      mailData.Code,
	}

	ctx.JSON(http.StatusOK, data)

}

type PasswordResetCode struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

// VerifyPassWordResetCode   handles the password reset request
func (r *Repository) VerifyPassWordResetCode(ctx *gin.Context) {

	var req PasswordResetCode

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	data := r.userServices.VerifyPasswordResetCode(req.Email)

	if req.Code != data.Code {
		e := errors.NewBadRequestError("Invalid code provided.")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return
	}

	resp := &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Code verification successful",
		Status:    http.StatusOK,
		Data:      nil,
	}

	ctx.JSON(resp.Status, resp)

}

type PasswordResetReq struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Email           string `json:"email"`
	Code            string `json:"code"`
}

// ResetPassword  handles the password reset request
func (r *Repository) ResetPassword(ctx *gin.Context) {

	var req PasswordResetReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := r.userServices.GetUserByEmail(req.Email)
	if user == nil {
		log.Println(err)
		data := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(data.Status, data)
		return
	}

	data := r.userServices.VerifyPasswordResetCode(req.Email)
	if req.Code != data.Code {
		log.Println(data)
		e := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return
	}

	if req.Code != data.Code {
		r.App.ErrorLog.Error("verification code did not match even after verifying PassReset")
		e := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return
	}

	if req.Password != req.PasswordConfirm {
		r.App.ErrorLog.Error("password and password re-enter did not match")
		e := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return

	}
	newPassword := crypto_utils.Hash(req.Password)
	var pwd = models.User{
		PasswordHash: newPassword,
	}
	err = r.userServices.UpdateUserDetails(user.ID, pwd)
	if err != nil {
		r.App.ErrorLog.Error("update user failed")
		e := errors.NewBadRequestError("Unable to reset password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return

	}
	r.mailService.RemoveMailCode(req.Email)
	resp := &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Password reset successful",
		Status:    http.StatusOK,
		Data:      nil,
	}

	ctx.JSON(resp.Status, resp)

}

type UpdateUserPasswordRequest struct {
	NewPassword string `json:"new_password"`
}

type updateUserDetails struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (r *Repository) UpdateUserDetails(ctx *gin.Context) {

	user := r.GetPayloadFromContext(ctx)
	var req updateUserDetails
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		log.Println(err)
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	um := models.User{
		UserName: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
	}
	err := r.userServices.UpdateUserDetails(user.ID, um)
	if err != nil {
		restErr := errors.NewBadRequestError("Unable to update user details.")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	data := NewStatusOkResponse("Successfully updated profile", nil)

	ctx.JSON(data.Status, data)
}

func (r *Repository) UpdateUserPassword(ctx *gin.Context) {
	user := r.GetPayloadFromContext(ctx)

	var req UpdateUserPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	newPassword := crypto_utils.Hash(req.NewPassword)
	err := r.userServices.UpdateUserPassword(user.ID, newPassword)
	if err != nil {
		r.App.ErrorLog.Error("update user failed")
		e := errors.NewBadRequestError("Unable to update password. Please try again later")
		ctx.JSON(e.Status, e)
		ctx.Abort()
		return

	}

}
