package utils

import (
	"regexp"

	"bitbucket.org/wycemiro/cljm.git/internal/models"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

// ValidateUser returns a slice of string of validation errors
func ValidateUser(user models.Register, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 6 {
		err = append(err, "Invalid password, Password should be more than 6 characters")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}

	return err
}

func ValidatePasswordReset(resetPassword models.ResetPassword)(bool,string){
	if len(resetPassword.Password) < 6{
		return false,"Invalid password, password should be more than 6 characters"
	}
	if resetPassword.Password != resetPassword.ConfirmPassword{
		return false,"Password reset failed, passwords must match"
	}
	return true,""
}