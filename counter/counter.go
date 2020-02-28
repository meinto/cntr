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

func (c *Counter) GetKeys(daysPast int) int {
	var client Client
	c.db.First(&client)

	now := time.Now()
	yearday := now.YearDay() - daysPast
	if yearday < 1 {
		yearday = 1
	}

	type Result struct {
		Total int64
	}
	var result Result
	c.db.Table("stats").
		Select("sum(keys) as total").
		Where(Stats{
			ClientUUID: client.UUID,
			Year:       now.Year(),
			YearDay:    yearday,
		}).
		Group("year, year_day, hour").
		Scan(&result)

	return int(result.Total)
}

func (c *Counter) GetClicks(daysPast int) int {
	var client Client
	c.db.First(&client)

	now := time.Now()
	yearday := now.YearDay() - daysPast
	if yearday < 1 {
		yearday = 1
	}

	type Result struct {
		Total int64
	}
	var result Result
	c.db.Table("stats").
		Select("sum(clicks) as total").
		Where(Stats{
			ClientUUID: client.UUID,
			Year:       now.Year(),
			YearDay:    yearday,
		}).
		Group("year, year_day, hour").
		Scan(&result)

	return int(result.Total)
}

func (c *Counter) incrementKeys() {
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

func (c *Counter) incrementClicks() {
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

	c.db.Model(&stats).Update("clicks", stats.Clicks+1)
}

func (c *Counter) Count() {
	go func() {
		evChan := hook.Start()
		defer hook.End()

		for ev := range evChan {
			if ev.Kind == hook.KeyUp {
				c.incrementKeys()
			}
			if ev.Kind == hook.MouseUp {
				c.incrementClicks()
			}
		}
	}()
}
