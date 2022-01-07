package postgresql_db

import (
	"fmt"
	"log"

	"maranatha_web/models"
	"maranatha_web/utils/errors"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func init() {

	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	dsn := "host=172.17.0.3 user=ianmuhia3 password=*#*Johnte2536 dbname=maranatha port=5432 sslmode=disable "

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 logger.Default,
	})

	MigrateTables(database)

	database.Logger.LogMode(logger.Info)

	//err = Client.AutoMigrate(models.Users{})

	if err != nil {
		panic(err)
	}
	Client = database

	log.Println("database successfully configured")

}

func MigrateTables(db *gorm.DB) *errors.RestErr {

	db.Logger.LogMode(logger.Info)

	if err := db.AutoMigrate(&users{}); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("New tables are  migrated successfully!")
	}
	return nil
}

func GetDB() *gorm.DB {
	return Client
}
