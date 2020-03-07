package counter

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/meinto/cntr/db"
	hook "github.com/robotn/gohook"
)

type Counter struct {
	db *gorm.DB
}

func NewCounter(db *gorm.DB) *Counter {
	return &Counter{db}
}

func (c *Counter) GetKeys(from, to time.Time) int {
	var client db.Client
	c.db.First(&client)

	toMinusOneSec := to.Add(-1 * time.Second)

	type Result struct {
		Total int64
	}
	var result Result
	c.db.Table("stats").
		Select("sum(keys) as total").
		Where(db.Stats{ClientUUID: client.UUID}).
		Where("timeslot BETWEEN ? AND ?", from, toMinusOneSec).
		Group("year, year_day").
		Scan(&result)

	return int(result.Total)
}

func (c *Counter) GetClicks(from, to time.Time) int {
	var client db.Client
	c.db.First(&client)

	toMinusOneSec := to.Add(-1 * time.Second)

	type Result struct {
		Total int64
	}
	var result Result
	c.db.Table("stats").
		Select("sum(clicks) as total").
		Where(db.Stats{ClientUUID: client.UUID}).
		Where("timeslot BETWEEN ? AND ?", from, toMinusOneSec).
		Group("year, year_day").
		Scan(&result)

	return int(result.Total)
}

func (c *Counter) incrementKeys() {
	var client db.Client
	c.db.First(&client)

	now := time.Now()
	timeslot := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())

	var stats db.Stats
	c.db.Where(db.Stats{
		ClientUUID: client.UUID,
		Timeslot:   &timeslot,
		Year:       now.Year(),
		YearDay:    now.YearDay(),
		Hour:       now.Hour(),
	}).FirstOrCreate(&stats)

	c.db.Model(&stats).Update("keys", stats.Keys+1)
}

func (c *Counter) incrementClicks() {
	var client db.Client
	c.db.First(&client)

	now := time.Now()
	timeslot := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())

	var stats db.Stats
	c.db.Where(db.Stats{
		ClientUUID: client.UUID,
		Year:       now.Year(),
		Timeslot:   &timeslot,
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
