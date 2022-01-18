package main

import (
	"clashG/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("clashG")

	content := layout.NewLayout(window)
	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.ShowAndRun()
}
