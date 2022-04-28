package layout

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"

	"clashG/api"
	"clashG/api/executor"
	"clashG/api/utils"
)

var (
	input = widget.NewEntry()
)

// 粘贴板初始化
func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func profilesScreen() fyne.CanvasObject {
	clip := string(clipboard.Read(clipboard.FmtText))
	if utils.IsUrlValid(clip) {
		input.SetText(clip)
	} else {
		input.SetPlaceHolder("please input a url or copy a url here")
	}

	download := widget.NewButton("Download", configDownload)
	topPanal := container.NewGridWithColumns(2, input, download)
	content := container.New(layout.NewBorderLayout(topPanal, nil, nil, nil), topPanal)
	return content
}

func configDownload() {
	url := input.Text
	if !utils.IsUrlValid(url) {
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
