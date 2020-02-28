package main

import (
	"github.com/meinto/cntr/counter"
	"github.com/meinto/cntr/systemtray"
)

func main() {
	c := counter.NewCounter()
	c.Count()
	s := systemtray.NewSystemtrayWidget(c)
	s.Run()
}
