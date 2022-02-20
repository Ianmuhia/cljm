package repository

import (
	"fmt"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
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
	NewPartnersQuery() PartnersQuery
	NewEventsQuery() EventsQuery
	NewGenresQuery() GenresQuery
	NewJobsQuery() JobsQuery
	NewNewsQuery() NewsQuery
	NewSermonQuery() SermonQuery
	NewPrayerRequestQuery() PrayerRequestQuery
	NewUserQuery() UserQuery
	NewTestimonyQuery() TestimonyQuery
	NewVolunteerQuery() VolunteerQuery
	NewPodcastQuery() PodcastQuery
}

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

func (d *postgresDBRepo) NewBookQuery() BookQuery {
	return &bookQuery{dbRepo: *d}
}

func (d *postgresDBRepo) NewPartnersQuery() PartnersQuery {
	return &partnerQuery{dbRepo: *d}
}
func (d *postgresDBRepo) NewEventsQuery() EventsQuery {
	return &eventsQuery{dbRepo: *d}
}

func (d *postgresDBRepo) NewGenresQuery() GenresQuery {
	return &genresQuery{dbRepo: *d}
}

func (d *postgresDBRepo) NewJobsQuery() JobsQuery {
	return &jobsQuery{dbRepo: *d}
}
func (d *postgresDBRepo) NewNewsQuery() NewsQuery {
	return &newsQuery{dbRepo: *d}
}

func (d *postgresDBRepo) NewPrayerRequestQuery() PrayerRequestQuery {
	return &prayerRequestQuery{repo: *d}
}

func (d *postgresDBRepo) NewSermonQuery() SermonQuery {
	return &sermonQuery{repo: *d}
}

func (d *postgresDBRepo) NewTestimonyQuery() TestimonyQuery {
	return &testimonyQuery{repo: *d}
}

func (d *postgresDBRepo) NewUserQuery() UserQuery {
	return &userQuery{repo: *d}
}
func (d *postgresDBRepo) NewVolunteerQuery() VolunteerQuery {
	return &volunteerQuery{repo: *d}
}
func (d *postgresDBRepo) NewPodcastQuery() PodcastQuery {
	return &podcastQuery{dbRepo: *d}
}

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

type users models.User
type news models.News
type churchPartner models.ChurchPartner
type sermon models.Sermon
type books models.Books
type genre models.Genre

type prayer models.Prayer
type podcast models.Podcast

//type churchEvent models.ChurchEvent
type testimonies models.Testimonies

//type churchJob models.ChurchJob
//type volunteerChurchJob models.VolunteerChurchJob

func GetDatabaseConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
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

	database.Logger.LogMode(logger.Info)
	Client = database

	log.Println("database successfully configured")

	return Client

}

func MigrateTables(db *gorm.DB) *errors.RestErr {

	db.Logger.LogMode(logger.Info)

	if err := db.AutoMigrate(
		&users{},
		&news{},
		&churchPartner{},
		&sermon{},
		&books{},
		&genre{},
		//&churchJob{},
		//&churchEvent{},
		//&volunteerChurchJob{},
		&prayer{},
		&testimonies{},
		&podcast{},
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
