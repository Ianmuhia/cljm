package postgresql_db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

const (
	DbUsername = "DbUsername"
	DbSchema   = "DbSchema"
	DbHost     = "DbHost"
	DbPassword = "DbPassword"
)

var (
	// Client *sql.DB
	Client *gorm.DB

	// username = os.Getenv(DbUsername)
	// host     = os.Getenv(DbHost)
	// schema   = os.Getenv(DbSchema)
	// password = os.Getenv(DbPassword)
)

type users models.User
type news models.News
type churchPartner models.ChurchPartner
type sermon models.Sermon
type books models.Books
type genre models.Genre

type prayer models.Prayer
type churchEvent models.ChurchEvent
type testimonies models.Testimonies
type churchJob models.ChurchJob
type volunteerChurchJob models.VolunteerChurchJob

func init() {

	dsn := "host=localhost user=wise password=*20406005 dbname=clj port=5432 sslmode=disable "

	database, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 logger.Default,
	})

	MigrateTables(database)

	database.Logger.LogMode(logger.Silent)
	Client = database

	log.Println("database successfully configured")

}

func MigrateTables(db *gorm.DB) *errors.RestErr {

	db.Logger.LogMode(logger.Info)

	if err := db.AutoMigrate(
		&users{},
		&news{},
		&churchPartner{},
		&sermon{},
		&books{},
		genre{},
		&churchJob{},
		&churchEvent{},
		&volunteerChurchJob{},
		&prayer{},
		&testimonies{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("New tables are  migrated successfully!")
	}
	// create database foreign key for user & credit_cards
	err := db.Migrator().CreateConstraint(&news{}, "News")
	if err != nil {
		panic(err)

	}
	err = db.Migrator().CreateConstraint(&news{}, "fk_users_credit_cards")
	if err != nil {
		panic(err)

	}
	return nil
}

func GetDB() *gorm.DB {
	return Client
}
