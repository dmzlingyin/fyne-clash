package layout

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func generalScreen() fyne.CanvasObject {
	port := canvas.NewText("Port", color.White)
	portValue := canvas.NewText("7890", color.White)
	allowLAN := canvas.NewText("Allow LAN", color.White)
	allowLANValue := widget.NewSlider(0.0, 1.0)
	logLevel := canvas.NewText("Log Level", color.White)
	logLevelValue := widget.NewSelect([]string{"slient", "info", "warning", "error", "debug"}, setLevel)

	content := container.New(layout.NewFormLayout(), port, portValue, allowLAN, allowLANValue, logLevel, logLevelValue)
	return content
}

func setLevel(value string) {
	log.Print(value)
}
