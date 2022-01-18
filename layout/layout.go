package layout

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	content = container.NewMax()
)

func NewLayout(window fyne.Window) *container.Split {
	generalButton := widget.NewButton("General", gerneral)
	proxiesButton := widget.NewButton("Proxies", proxies)
	profilesButton := widget.NewButton("Profiles", profiles)
	logsButton := widget.NewButton("Logs", logs)

	navgater := container.NewVBox(generalButton, layout.NewSpacer(), proxiesButton, layout.NewSpacer(), profilesButton, layout.NewSpacer(), logsButton)
	page := container.NewHSplit(navgater, content)
	page.SetOffset(0.2)
	return page
}

func gerneral() {
	newContent := generalScreen()
	content.Objects = []fyne.CanvasObject{newContent}
	content.Refresh()
}

func proxies() {
	newContent := proxiesScreen()
	content.Objects = []fyne.CanvasObject{newContent}
	content.Refresh()
	log.Print("proxiesButton clicked.")
}

func profiles() {
	newContent := profilesScreen()
	content.Objects = []fyne.CanvasObject{newContent}
	content.Refresh()
	log.Print("profilesButton clicked.")
}

func logs() {
	newContent := logsScreen()
	content.Objects = []fyne.CanvasObject{newContent}
	content.Refresh()
	log.Print("logsButton clicked.")
}
