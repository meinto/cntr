package systemtray

import (
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/meinto/cntr/counter"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Key count: 0")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	for {
		systray.SetTitle("Key count: " + strconv.FormatInt(int64(counter.ActiveCounter.GetKeys()), 10))
		time.Sleep(time.Second)
	}
}

func onExit() {
	// clean up here
}
