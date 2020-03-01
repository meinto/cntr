package main

import (
	"log"

	"github.com/meinto/cntr/counter"
	"github.com/meinto/cntr/db"
	"github.com/meinto/cntr/server"
	"github.com/meinto/cntr/systemtray"
)

func main() {
	gormdb, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer gormdb.Close()

	db.Automigrate(gormdb)
	db.Init(gormdb)

	c := counter.NewCounter(gormdb)
	c.Count()
	s := server.NewServer(gormdb, c)
	s.Start()
	t := systemtray.NewSystemtrayWidget(c)
	t.Run()
}
