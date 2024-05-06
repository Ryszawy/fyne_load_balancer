package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Ryszawy/fyne_load_balance/client"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func refreshTimer(table *widget.List) {
	table.Refresh()
}

func main() {
	a := app.New()
	w := a.NewWindow("Load balancer")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))

	bars := canvas.NewText("bars", color.White)
	clients := client.CreateEmptyClintsArr()
	clinetCounter := client.IDCounter()
	var table *widget.List

	updateTable := func() {
		table.Refresh()
	}

	table = widget.NewList(
		func() int {
			return len(*clients)
		},
		func() fyne.CanvasObject {
			clientNameLabel := widget.NewLabel("")
			addFileBtn := widget.NewButton("Add New File", nil)
			filesListContainer := container.NewHBox(widget.NewLabel("Empty"))
			elapsedTimeLabel := widget.NewLabel("")
			return container.NewGridWithColumns(4, clientNameLabel, filesListContainer, elapsedTimeLabel, addFileBtn)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := (*clients)[lii]

			objects := co.(*fyne.Container).Objects
			label := objects[0]
			filesList := objects[1].(*fyne.Container)
			filesList.Objects = nil
			elapsedTimeLabel := objects[2].(*widget.Label)
			fileBtn := objects[3]

			label.(*widget.Label).SetText(c.ClientName)

			for _, file := range *c.Files {
				fileSizeLabel := widget.NewLabel(
					fmt.Sprintf("File %d: %.2f MB", file.FileID, file.Size))
				filesList.Add(fileSizeLabel)
			}

			elapsedTime := c.ElapsedTime()
			elapsedTimeLabel.SetText(fmt.Sprintf("Elapsed Time: %.2f seconds", elapsedTime))
			fileBtn.(*widget.Button).OnTapped = func() {
				log.Println(c.ClientName)
				newFile := client.NewFile(1, 43.2)
				*c.Files = append(*c.Files, newFile)
				updateTable()
			}
		},
	)

	addClientBtn := widget.NewButton("Add Client", func() {
		id := clinetCounter()
		cName := fmt.Sprintf("Client %d", id)
		c := client.NewClient(id, cName)

		*clients = append(*clients, c)
		table.Refresh()
	})

	// fileSize := binding.NewFloat()
	fileSizeEntry := widget.NewEntry()
	fileSizeEntry.SetPlaceHolder("File Size")
	filesGrid := container.NewGridWithColumns(1, fileSizeEntry)
	controlMenu := container.NewGridWithColumns(2, widget.NewButton("Start", nil), widget.NewButton("Stop", nil))
	serverMenu := widget.NewButton("Create New Server", nil)
	menu := container.NewGridWithRows(4, serverMenu, filesGrid, addClientBtn, controlMenu)

	grid := container.NewGridWithColumns(2, bars, menu)
	tableContainer := container.NewGridWithColumns(1, table)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				refreshTimer(table)
			}
		}
	}()
	w.SetContent(container.NewGridWithColumns(1, grid, tableContainer))
	w.ShowAndRun()
}
