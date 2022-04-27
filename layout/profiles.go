package layout

import (
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"clashG/api"
	"clashG/api/executor"
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
	url := input.Text
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Println("Invalid url.")
		return
	}
	if !strings.Contains(url, "clash") {
		log.Println("Invalid clash describe url.")
		return
	}

	go func() {
		err := executor.DownloadConfig(url)
		if err != nil {
			log.Println("Download failed, please make sure your network are working.")
		}
		// clash 热更新配置文件
		err = api.Reload()
		if err != nil {
			log.Println("Reload config failed.")
		}
	}()
}
