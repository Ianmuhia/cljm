package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"maranatha_web/internal/config"
)

type DAO interface {
	NewBookQuery() BookQuery
	// NewUserQuery() UserQuery
	// NewCourseQuery() CourseQuery
	// NewAnswerQuery() AnswerQuery
	// NewIndicatorQuery() IndicatorQuery
	// NewQuestionQuery() QuestionQuery
	// NewReviewQuery() ReviewQuery
	// NewScoreQuery() ScoreQuery
	// NewSectionQuery() SectionQuery
	// NewTransactionQuery() TransactionQuery
}

// type dao struct {
// 	App *config.AppConfig
// 	DB  *gorm.DB
// }

// func NewDAO(db *sql.DB) DAO {
// 	DB = db
// 	return &dao{}
// }

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *gorm.DB
}

func NewPostgresRepo(conn *gorm.DB, a *config.AppConfig) DAO {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// func (d *dao) NewTransactionQuery() TransactionQuery {
// 	return &transactionQuery{}
// }

// func (d *dao) NewSectionQuery() SectionQuery {
// 	return &sectionQuery{}
// }

func (d *postgresDBRepo) NewBookQuery() BookQuery {
	return &bookQuery{dbRepo: *d}
}

// func (d *dao) NewReviewQuery() ReviewQuery {
// 	return &reviewQuery{}
// }

// func (d *dao) NewQuestionQuery() QuestionQuery {
// 	return &questionQuery{}
// }

// func (d *dao) NewUserQuery() UserQuery {
// 	return &userQuery{}
// }

// func (d *dao) NewCourseQuery() CourseQuery {
// 	return &courseQuery{}
// }

// func (d *dao) NewAnswerQuery() AnswerQuery {
// 	return &answerQuery{}
// }

// func (d *dao) NewIndicatorQuery() IndicatorQuery {
// 	return &indicatorQuery{}
// }

const (
	DbUsername = "DB_USER"
	DbSchema   = "DB_NAME"
	DbHost     = "DB_HOST"
	DbPort     = "DB_PORT"
	DbPassword = "DB_PASSWORD"
)

var (
	Client *gorm.DB
)

//type users models.User
//type news models.News
//type churchPartner models.ChurchPartner
//type sermon models.Sermon
//type books models.Books
//type genre models.Genre
//
//type prayer models.Prayer
//type churchEvent models.ChurchEvent
//type testimonies models.Testimonies
//type churchJob models.ChurchJob
//type volunteerChurchJob models.VolunteerChurchJob

func GetDatabaseConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error lc c c c oading .env file")
	}

	username := os.Getenv(DbUsername)
	host := os.Getenv(DbHost)
	port := os.Getenv(DbPort)
	schema := os.Getenv(DbSchema)
	password := os.Getenv(DbPassword)

	//dsn := "host=localhost user=wise password=*20406005 dbname=clj port=5432 sslmode=disable "
	dc := fmt.Sprintf("host=%v user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, schema, port)

	database, _ := gorm.Open(postgres.Open(dc), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 logger.Default,
	})

	// MigrateTables(database)

	database.Logger.LogMode(logger.Silent)
	Client = database

	log.Println("database successfully configured")

	return Client

}

// func MigrateTables(db *gorm.DB) *errors.RestErr {

// 	db.Logger.LogMode(logger.Info)

// 	if err := db.AutoMigrate(
// 		&users{},
// 		&news{},
// 		&churchPartner{},
// 		&sermon{},
// 		&books{},
// 		genre{},
// 		&churchJob{},
// 		&churchEvent{},
// 		&volunteerChurchJob{},
// 		&prayer{},
// 		&testimonies{},
// 	); err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	} else {
// 		fmt.Println("New tables are  migrated successfully!")
// 	}
// 	// create database foreign key for user & credit_cards
// 	err := db.Migrator().CreateConstraint(&news{}, "News")
// 	if err != nil {
// 		panic(err)

// 	}
// 	err = db.Migrator().CreateConstraint(&news{}, "fk_users_credit_cards")
// 	if err != nil {
// 		panic(err)

// 	}
// 	return nil
// }
