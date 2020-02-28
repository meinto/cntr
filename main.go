package main

import (
	"github.com/meinto/cntr/counter"
	"github.com/meinto/cntr/statusbar"
)

func main() {
	go counter.Count()
	statusbar.Run()
}
