package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Ryszawy/fyne_load_balance/client"

	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func refreshTimer(table *widget.List) {
	table.Refresh()
}

func createFilesListPerClient(c *client.Client) *widget.List {
	return widget.NewList(
		func() int {
			return len(*c.Files)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			cnt := fmt.Sprintf("File %d: %.2f MB", (*c.Files)[lii].FileID, (*c.Files)[lii].Size)
			co.(*widget.Label).SetText(cnt)
		},
	)
}

func showFiles(a fyne.App, c *client.Client) {
	filesWindow := a.NewWindow(fmt.Sprintf("%v Files", c.ClientName))
	filesWindow.Resize(fyne.NewSize(700, 700))
	grid := container.NewGridWithColumns(1, createFilesListPerClient(c))
	filesWindow.SetContent(grid)
	filesWindow.Show()
}

func main() {
	a := app.New()
	w := a.NewWindow("Load balancer")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))

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
			filesList := widget.NewButton("Check Files", nil)
			elapsedTimeLabel := widget.NewLabel("")
			return container.NewGridWithColumns(4, clientNameLabel, filesList, elapsedTimeLabel, addFileBtn)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := (*clients)[lii]

			objects := co.(*fyne.Container).Objects
			label := objects[0]
			showBtn := objects[1]
			showBtn.(*widget.Button).OnTapped = func() {
				showFiles(a, &c)
			}
			elapsedTimeLabel := objects[2].(*widget.Label)
			fileBtn := objects[3]
			label.(*widget.Label).SetText(c.ClientName)

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

	filesGrid := container.NewGridWithColumns(2, fileSizeEntry)
	controlMenu := container.NewGridWithColumns(2, widget.NewButton("Start", nil), widget.NewButton("Stop", nil))
	serverMenu := widget.NewButton("Create New Server", nil)
	menu := container.NewGridWithRows(4, serverMenu, filesGrid, addClientBtn, controlMenu)
	infinite := widget.NewProgressBarInfinite()
	infinite1 := widget.NewProgressBarInfinite()
	infinite2 := widget.NewProgressBarInfinite()
	grid := container.NewGridWithColumns(2, container.NewGridWithColumns(3, infinite, infinite1, infinite2), menu)
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
