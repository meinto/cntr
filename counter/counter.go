package counter

import (
	"time"

	hook "github.com/robotn/gohook"
)

type Counter struct {
	keys int
	time time.Time
}

func NewCounter() *Counter {
	return &Counter{0, time.Now()}
}

func (c *Counter) GetKeys() int {
	return c.keys
}

func (c *Counter) increment() {
	c.keys++
}

func (c *Counter) reset() {
	c.keys = 0
	c.time = time.Now()
}

func (c *Counter) Count() {
	go func() {
		evChan := hook.Start()
		defer hook.End()

		c.runResetCounterInterval()

		for ev := range evChan {
			if ev.Kind == hook.KeyUp {
				c.increment()
			}
		}
	}()
}

func (c *Counter) runResetCounterInterval() {
	go func() {
		for {
			if time.Now().Day() != c.time.Day() {
				c.reset()
			}
			time.Sleep(2 * time.Second)
		}
	}()
}
