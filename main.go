package main

import (
	"log"
	// "fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Ryszawy/fyne_load_balance/client"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// var tableData = [][]string{{"top left", "top right"},
// 	{"bottom left", "bottom right"}}

func main() {
	a := app.New()
	w := a.NewWindow("Load balancer")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))

	// data := binding.BindStruct(
	// 	&[]client{},
	// )

	// list := widget.NewListWithData(data,
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i binding.DataItem, o fyne.CanvasObject) {
	// 		o.(*widget.Label).Bind(i.(binding.String))
	// 	})

	// add := widget.NewButton("Append", func() {
	// 	val := fmt.Sprintf("Item %d", data.Length()+1)
	// 	data.Append(val)
	// })
	// w.SetContent(container.NewBorder(nil, add, nil, nil, list))

	bars := canvas.NewText("bars", color.White)
	menu := canvas.NewText("menu", color.White)
	clients := []client.Client{}
	clientCounter := 1
	table := widget.NewButton("Test Client", func() {
		c := client.NewClient(clientCounter, "c1")
		clients = append(clients, c)
		clientCounter++
		log.Println(clients, clients[0].ElapsedTime())
	})
	// table := widget.NewTable(
	// 	func() (int, int) {
	// 		return len(tableData), len(tableData[0])
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("wide content")
	// 	},
	// 	func(i widget.TableCellID, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(tableData[i.Row][i.Col])
	// 	})

	grid := container.NewGridWithColumns(2, bars, menu)

	w.SetContent(container.NewGridWithColumns(1, grid, table))
	w.ShowAndRun()
}
