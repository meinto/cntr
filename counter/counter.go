package counter

import (
	"time"

	hook "github.com/robotn/gohook"
)

type Counter struct {
	keys int
	time time.Time
}

func (c *Counter) GetKeys() int {
	return c.keys
}

func (c *Counter) Increment() {
	c.keys++
}

func (c *Counter) Reset() {
	c.keys = 0
	c.time = time.Now()
}

var ActiveCounter = &Counter{0, time.Now()}

func Count() {
	evChan := hook.Start()
	defer hook.End()

	go shouldResetCounter()

	for ev := range evChan {
		if ev.Kind == hook.KeyUp {
			ActiveCounter.Increment()
		}
	}
}

func shouldResetCounter() {
	for {
		if time.Now().Day() != ActiveCounter.time.Day() {
			ActiveCounter.Reset()
		}
		time.Sleep(2 * time.Second)
	}
}
