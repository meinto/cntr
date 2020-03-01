package systemtray

import (
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/meinto/cntr/counter"
)

type Systemtray struct {
	counter *counter.Counter
}

func NewSystemtrayWidget(c *counter.Counter) *Systemtray {
	return &Systemtray{c}
}

func (s *Systemtray) Run() {
	systray.RunWithAppWindow("cntr", 556, 442, s.onReady, s.onExit)
}

func (s *Systemtray) onReady() {
	systray.SetTitle("Key count: 0")
	systray.AddSeparator()
	keysToday := systray.AddMenuItem("Key Today: ", "Key Today")
	clicksToday := systray.AddMenuItem("Clicks Today: ", "Clicks Today")
	systray.AddSeparator()
	keysYesterday := systray.AddMenuItem("Key Yesterday: ", "Key Yesterday")
	clicksYesterday := systray.AddMenuItem("Clicks Yesterday: ", "Clicks Yesterday")
	systray.AddSeparator()
	openAppWindow := systray.AddMenuItem("Open App", "Open App Window")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-quit.ClickedCh:
				systray.Quit()
			case <-openAppWindow.ClickedCh:
				systray.ShowAppWindow("http://localhost:5564")
			}
		}
	}()

	for {
		now := time.Now()
		year := now.Year()
		yearday := now.YearDay()
		systray.SetTitle("Keys: " + strconv.FormatInt(int64(s.counter.GetKeys(year, yearday)), 10))
		keysToday.SetTitle("Keys Today: " + strconv.FormatInt(int64(s.counter.GetKeys(year, yearday)), 10))
		clicksToday.SetTitle("Clicks Today: " + strconv.FormatInt(int64(s.counter.GetClicks(year, yearday)), 10))

		yesterday := now.Add(-1 * 24 * time.Hour)
		year = yesterday.Year()
		yearday = yesterday.YearDay()
		keysYesterday.SetTitle("Keys Yesterday: " + strconv.FormatInt(int64(s.counter.GetKeys(year, yearday)), 10))
		clicksYesterday.SetTitle("Clicks Yesterday: " + strconv.FormatInt(int64(s.counter.GetClicks(year, yearday)), 10))
		time.Sleep(time.Second)
	}
}

func (s *Systemtray) onExit() {
	// clean up here
}
