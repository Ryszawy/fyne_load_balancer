package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Ryszawy/fyne_load_balance/timer"
	//"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	w.SetMaster()
	w.SetFullScreen(true)

	clock := widget.NewLabel("")
	timer.UpdateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			timer.UpdateTime(clock)
		}
	}()
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Show()

	w.ShowAndRun()
}
