package layout

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"clashG/api"
	"clashG/widgets"
)

var Delay = container.NewMax()

func proxiesScreen() fyne.CanvasObject {
	proxies := api.GetProxies().All
	buttons := make([]fyne.CanvasObject, len(proxies))
	delayButton := widget.NewButton("delay test", delayTest)
	for i := 0; i < len(proxies); i++ {
		button := widgets.NewButton(proxies[i], "Check")
		buttons[i] = button
	}
	Delay.Add(delayButton)
	lay := container.NewVBox(Delay, container.NewGridWithColumns(2, buttons...))
	content := container.NewCenter(lay)
	// scroll := container.NewScroll(content)

	return content
}

func delayTest() {
	proxies := api.GetProxies().All
	progress := widget.NewProgressBar()
	progress.Min = 0.0
	progress.Max = float64(len(proxies))
	Delay.Add(progress)
	Delay.Refresh()

	for i, proxy := range proxies {
		go api.GetProxyDelayByName(proxy)
		go progress.SetValue(float64(i + 1))
	}
	time.Sleep(time.Second)
	Delay.Remove(progress)
}
