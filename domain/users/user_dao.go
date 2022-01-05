package users

import (
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"

	"gorm.io/gorm"
)

type userResponse struct {
	Username          string `json:"username"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PasswordChangedAt string `json:"password_changed_at"`
	CreatedAt         string `json:"created_at"`
}

func newUserResponse(user models.User) userResponse {
	return userResponse{
		Username:          user.UserName,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.UpdatedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func GetUser(param string) (*errors.RestErr, *models.User) {
	var user models.User
	err := postgresql_db.Client.Debug().Where("email = ?", param).First(&user).Scan(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.NewBadRequestError("User does not exit"), &user
		// if user.Email == "" {
		// 	return nil, &user
		// }

		// return nil, &user
	}
	return nil, &user
}

func Create(user *models.User) *errors.RestErr {

	err := postgresql_db.Client.Debug().Create(&user).Error

	if err != nil {
		logger.Error("error when trying to save user", err)

		return errors.NewInternalServerError("database error")
	}
	return nil
}

func FindByEmailAndPassword() *errors.RestErr {
	// 	if err := users_db.Client.Ping(); err != nil {
	// 		panic(err)
	// 	}
	// 	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	// 	if err != nil {
	// 		logger.Error("error when trying to prepare get user by email and password statement", err)
	// 		return errors.NewInternalServerError("database error")
	// 	}
	// 	defer func(stmt *sql.Stmt) {
	// 		err := stmt.Close()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}(stmt)
	// 	result := stmt.QueryRow(user.Password, user.Email, StatusActive)
	// 	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
	// 		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
	// 			return errors.NewNotFoundError("invalid user credentials")
	// 		}
	// 		logger.Error("error when trying to get user by email and password", getErr)
	// 		return errors.NewInternalServerError("database error")
	// 		//return mysql_utils.ParseError(getErr)

	// 	}
	// 	return nil
	// }

	// func (user *User) Save() *errors.RestErr {
	// 	stmt, err := users_db.Client.Prepare(queryInsertUser)
	// 	if err != nil {
	// 		logger.Error("error when trying to prepare statement", err)
	// 		return errors.NewInternalServerError("database error")

	// 	}
	// 	defer func(stmt *sql.Stmt) {
	// 		err := stmt.Close()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}(stmt)

	// 	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	// 	if saveErr != nil {
	// 		logger.Error("error when trying to save user", saveErr)
	// 		return errors.NewInternalServerError("database error")

	// 		//return mysql_utils.ParseError(saveErr)

	// 	}
	// 	userId, err := insertResult.LastInsertId()
	// 	if err != nil {
	// 		logger.Error("error when trying to get last insert id after creating a new user", err)
	// 		return errors.NewInternalServerError("database error")

	// 		//return mysql_utils.ParseError(err)

	// 	}
	// user.Id = userId

	return nil
}

func Update() *errors.RestErr {
	// 	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	// 	if err != nil {
	// 		logger.Error("error when trying to prepare update user statement", err)
	// 		return errors.NewInternalServerError("database error")

	// 		//return errors.NewInternalServerError(err.Error())
	// 	}
	// 	defer func(stmt *sql.Stmt) {
	// 		err := stmt.Close()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}(stmt)
	// 	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	// 	if err != nil {
	// 		logger.Error("error when trying to update user details", err)
	// 		return errors.NewInternalServerError("database error")
	// 		//
	// 		//return mysql_utils.ParseError(err)
	// 	}
	// 	return nil
	// }

	// func (user *User) Delete() *errors.RestErr {
	// 	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	// 	if err != nil {
	// 		logger.Error("error when trying to prepare delete user statement", err)
	// 		return errors.NewInternalServerError("database error")

	// 		//return errors.NewInternalServerError(err.Error())
	// 	}
	// 	defer func(stmt *sql.Stmt) {
	// 		err := stmt.Close()
	// 		if err != nil {

	// 		}
	// 	}(stmt)
	// 	_, err = stmt.Exec(user.Id)
	// 	if err != nil {
	// 		logger.Error("error when trying to delete user", err)
	// 		return errors.NewInternalServerError("database error")

	// 		//return mysql_utils.ParseError(err)
	// 	}
	return nil

}

func Search(status string) ([]models.User, *errors.RestErr) {

	// stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	// if err != nil {
	// 	logger.Error("error when trying to prepare find users by status statement", err)
	// 	return nil, errors.NewInternalServerError("database error")

	// }
	// defer func(stmt *sql.Stmt) {
	// 	err := stmt.Close()
	// 	if err != nil {

	// 	}
	// }(stmt)
	// rows, err := stmt.Query(status)
	// if err != nil {
	// 	logger.Error("error when trying to find users by status", err)
	// 	return nil, errors.NewInternalServerError("database error")

	// }
	// defer func(rows *sql.Rows) {
	// 	err := rows.Close()
	// 	if err != nil {

	// 	}
	// }(rows)
	results := make([]models.User, 0)
	// for rows.Next() {
	// 	var user User
	// 	if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
	// 		logger.Error("error when trying to scan row into user struct", err)
	// 		return nil, errors.NewInternalServerError("database error")

	// 	}
	// 	results = append(results, user)
	// }
	// if len(results) == 0 {
	// 	return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	// }
	return results, nil
}
