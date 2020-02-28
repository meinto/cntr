package counter

import (
	"time"

	"github.com/jinzhu/gorm"
	hook "github.com/robotn/gohook"
)

type Counter struct {
	db   *gorm.DB
	keys int
	time time.Time
}

func NewCounter(db *gorm.DB) *Counter {
	return &Counter{db, 0, time.Now()}
}

func (c *Counter) GetKeys() int {
	var client Client
	c.db.First(&client)
	var stats Stats
	c.db.Where(Stats{
		ClientUUID: client.UUID,
		Year:       time.Now().Year(),
		YearDay:    time.Now().YearDay(),
		Hour:       time.Now().Hour(),
	}).FirstOrCreate(&stats)
	return stats.Keys
}

func (c *Counter) increment() {
	var client Client
	c.db.First(&client)

	now := time.Now()

	var stats Stats
	c.db.Where(Stats{
		ClientUUID: client.UUID,
		Year:       now.Year(),
		YearDay:    now.YearDay(),
		Hour:       now.Hour(),
	}).FirstOrCreate(&stats)

	c.db.Model(&stats).Update("keys", stats.Keys+1)
}

func (c *Counter) Count() {
	go func() {
		evChan := hook.Start()
		defer hook.End()

		for ev := range evChan {
			if ev.Kind == hook.KeyUp {
				c.increment()
			}
		}
	}()
}
