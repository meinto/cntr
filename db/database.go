package db

import (
	"log"
	"os"
	"os/user"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDatabase() (*gorm.DB, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := usr.HomeDir + "/Library/ApplicationSupport/cntr/"
	os.MkdirAll(dbPath, os.ModePerm)
	dbName := "cntr.db"
	log.Printf("database at: %s", dbPath+dbName)
	return gorm.Open("sqlite3", dbPath+dbName)
}

func Automigrate(db *gorm.DB) {
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Stats{})

	var client Client
	db.First(&client)

	var stats []Stats
	db.Where(Stats{
		ClientUUID: client.UUID,
		Timeslot:   nil,
	}).Find(&stats)

	for _, s := range stats {
		current := time.Now()
		for s.Year != current.Year() || s.YearDay != current.YearDay() {
			current = current.Add(-1 * 24 * time.Hour)
		}

		timeslot := time.Date(s.Year, current.Month(), current.Day(), s.Hour, 0, 0, 0, current.Location())
		db.Model(&s).Update("timeslot", timeslot)
	}
}

func Init(db *gorm.DB) {
	var client Client
	db.First(&client)
	if client.ID == 0 {
		db.Create(Client{
			UUID: uuid.Must(uuid.NewV4()),
		})
	}
}
