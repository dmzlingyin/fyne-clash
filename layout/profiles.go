package layout

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	input = widget.NewEntry()
)

func profilesScreen() fyne.CanvasObject {
	input.SetPlaceHolder("please input a url or copy a url here")
	download := widget.NewButton("Download", configDownload)
	topPanal := container.NewGridWithColumns(2, input, download)

	content := container.New(layout.NewBorderLayout(topPanal, nil, nil, nil), topPanal)
	return content
}

func configDownload() {
	fmt.Println(input.Text)
}
