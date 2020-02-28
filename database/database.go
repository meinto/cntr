package database

import (
	"log"
	"os"
	"os/user"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/meinto/cntr/counter"
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
	db.AutoMigrate(&counter.Client{})
	db.AutoMigrate(&counter.Stats{})
}

func Init(db *gorm.DB) {
	var client counter.Client
	db.FirstOrInit(&client, counter.Client{
		UUID: uuid.Must(uuid.NewV4()),
	})
}
