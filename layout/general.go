package layout

import (
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"clashG/api"
	"clashG/api/executor"
)

const (
	NoProxy = iota
	ManualProxy
	AutoConfig
	AutoDetect
	SystemProxy
)

func generalScreen() fyne.CanvasObject {
	logo := canvas.NewImageFromFile("/home/lingyin/go/my_src/clashG/data/logo.png")
	logo.FillMode = canvas.ImageFillOriginal
	logo.Resize(fyne.NewSize(200, 200))
	logoText := canvas.NewText("Clash Dashboard", color.White)
	logoPanal := container.NewGridWithColumns(2, logo, logoText)
	separator := widget.NewSeparator()

	config := api.NewConfig()
	port := canvas.NewText("Port", color.White)
	portValue := canvas.NewText(strconv.Itoa(config.Port), color.White)

	allowLAN := canvas.NewText("Allow LAN", color.White)
	allowLANValue := widget.NewCheck("", toggleLAN)
	allowLANValue.Checked = config.AllowLAN

	logLevel := canvas.NewText("Log Level", color.White)
	logLevelValue := widget.NewSelect([]string{"silent", "info", "warning", "error", "debug"}, setLevel)
	logLevelValue.SetSelected(config.LogLevel)
	logLevelValue.MinSize()

	systemProxy := canvas.NewText("System Proxy", color.White)
	systemProxyValue := widget.NewCheck("", toggleSystemProxy)
	systemProxyValue.Checked = executor.IsSystemProxy()

	startup := canvas.NewText("Start with Linux", color.White)
	startupValue := widget.NewCheck("", toggleStartup)
	startupValue.Checked = executor.IsStartUp()

	mainPanal := container.NewGridWithColumns(2, port, portValue, allowLAN, allowLANValue,
		systemProxy, systemProxyValue, startup, startupValue, logLevel, logLevelValue)
	content := container.NewCenter(container.NewVBox(logoPanal, separator, mainPanal))

	return content
}

func setLevel(value string) {
	log.Print(value)
}

func toggleLAN(checked bool) {
	log.Print(checked)
}

func toggleSystemProxy(checked bool) {
	var err error
	if checked {
		err = executor.SetProxy(SystemProxy)
	} else {
		err = executor.SetProxy(NoProxy)
	}

	if err != nil {
		log.Println(err)
	}
}

func toggleStartup(checked bool) {
	var err error
	if checked {
		err = executor.AutoStart(checked)
	} else {
		err = executor.AutoStart(checked)
	}

	if err != nil {
		log.Fatal(err)
	}
}
