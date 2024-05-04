package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Ryszawy/fyne_load_balance/client"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Load balancer")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))

	bars := canvas.NewText("bars", color.White)
	clients := client.CreateEmptyClintsArr()
	clinetCounter := client.IDCounter()

	table := widget.NewList(
		func() int {
			return len(*clients)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := (*clients)[lii]
			label := co.(*widget.Label)
			label.SetText(c.ClientName)
		},
	)
	addClientBtn := widget.NewButton("Add Client", func() {
		id := clinetCounter()
		cName := fmt.Sprintf("Client %d", id)
		c := client.NewClient(id, cName)

		*clients = append(*clients, c)
		table.Refresh()
	})

	menuLabel := widget.NewLabel("Create Client")
	// fileSize := binding.NewFloat()
	fileSizeEntry := widget.NewEntry()
	fileSizeEntry.SetPlaceHolder("File Size")
	filesGrid := container.NewGridWithColumns(2, fileSizeEntry)

	menu := container.NewGridWithRows(4, menuLabel, filesGrid, bars, addClientBtn)

	grid := container.NewGridWithColumns(2, bars, menu)
	tableContainer := container.NewGridWithColumns(1, table)

	w.SetContent(container.NewGridWithColumns(1, grid, tableContainer))
	w.ShowAndRun()
}
