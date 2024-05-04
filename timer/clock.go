package timer

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func UpdateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
