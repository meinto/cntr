package main

import (
	"github.com/meinto/cntr/counter"
	"github.com/meinto/cntr/systemtray"
)

func main() {
	go counter.Count()
	systemtray.Run()
}
