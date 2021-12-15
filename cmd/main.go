package main

import (
	"bitbucket.org/wycemiro/cljm.git/internal/driver"
	"bitbucket.org/wycemiro/cljm.git/internal/models"
	"encoding/gob"
	"log"
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
	gob.Register(models.Genre{})
	gob.Register(models.Gallery{})
	gob.Register(models.GalleryImages{})
	gob.Register(models.Sermon{})

	//connect to database
	log.Println("Connecting to database...🍃")
	db, err := driver.ConnectSQL("host=postgres_clj port=5432 dbname=clj user=wise password=*20406005aA!")
	if err != nil {
		log.Fatal("Can't connect to the database ☠️")
	}
	log.Println("Connected to the database 🎉")
	return db, nil
}
