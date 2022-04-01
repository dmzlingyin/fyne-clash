package main

import (
	"clashG/layout"
	"clashG/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	myTheme := theme.MyTheme{}
	myTheme.SetFonts("./data/font.ttf", "")
	app.Settings().SetTheme(&myTheme)
	window := app.NewWindow("clashG")

	content := layout.NewLayout(window)
	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
	// window.SetFixedSize(true)
	window.CenterOnScreen()
	window.ShowAndRun()
}
