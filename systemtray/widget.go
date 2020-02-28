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
	systray.Run(s.onReady, s.onExit)
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
	quit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	for {
		systray.SetTitle("Keys: " + strconv.FormatInt(int64(s.counter.GetKeys(0)), 10))
		keysToday.SetTitle("Keys Today: " + strconv.FormatInt(int64(s.counter.GetKeys(0)), 10))
		clicksToday.SetTitle("Clicks Today: " + strconv.FormatInt(int64(s.counter.GetClicks(0)), 10))
		keysYesterday.SetTitle("Keys Yesterday: " + strconv.FormatInt(int64(s.counter.GetKeys(1)), 10))
		clicksYesterday.SetTitle("Clicks Yesterday: " + strconv.FormatInt(int64(s.counter.GetClicks(1)), 10))
		time.Sleep(time.Second)
	}
}

func (s *Systemtray) onExit() {
	// clean up here
}
