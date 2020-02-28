package main

import (
	"log"

	"github.com/meinto/cntr/counter"
	"github.com/meinto/cntr/database"
	"github.com/meinto/cntr/systemtray"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.Automigrate(db)
	database.Init(db)

	c := counter.NewCounter(db)
	c.Count()
	s := systemtray.NewSystemtrayWidget(c)
	s.Run()
}
