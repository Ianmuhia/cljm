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

func init() {

	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	dsn := "host=localhost user=ianmuhia3 password=*#*Johnte2536 dbname=maranatha port=5432 sslmode=disable "

	database, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 logger.Default,
	})

	MigrateTables(database)

	database.Logger.LogMode(logger.Silent)
	// database.Logger.Trace()

	//err := Client.AutoMigrate(&models.User{})
	//
	//if err != nil {
	//	panic(err)
	//}
	Client = database

	log.Println("database successfully configured")

}

func MigrateTables(db *gorm.DB) *errors.RestErr {

	db.Logger.LogMode(logger.Info)

	if err := db.AutoMigrate(&users{}, &news{}, &churchPartner{}); err != nil {
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
