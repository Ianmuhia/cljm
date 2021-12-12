package main

import (
	"encoding/gob"
	"log"

	"bitbucket.org/wycemiro/cljm.git/internal/driver"
	"bitbucket.org/wycemiro/cljm.git/internal/models"
)

const porte = ":8000"

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	r := SetupRouter()
	r.Run(porte)
}
func run() (*driver.DB, error) {
	gob.Register(models.Users{})
	gob.Register(models.News{})
	gob.Register(models.NewsImages{})
	gob.Register(models.Blogs{})
	gob.Register(models.BlogImages{})
	gob.Register(models.Books{})
	gob.Register(models.Geners{})
	gob.Register(models.Gallery{})
	gob.Register(models.GalleryImages{})
	gob.Register(models.Sermon{})

	
	//connect to database
	log.Println("Connecting to database...🍃")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=api password=50451103aA")
	if err != nil {
		log.Fatal("Can't connect to the database ☠️")
	}
	log.Println("Connected to the database 🎉")
	return db, nil
}
