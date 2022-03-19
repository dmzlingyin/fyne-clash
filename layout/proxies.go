package layout

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"clashG/api"
	"clashG/widgets"
)

var (
	Delay          = container.NewMax()
	vbox           = container.NewVBox()
	data           = binding.BindStringList(&[]string{})
	contentProxies = container.NewCenter()
)

func proxiesScreen() fyne.CanvasObject {
	proxies := api.GetProxies().All
	data.Set(proxies)

	buttons := make([]fyne.CanvasObject, len(proxies))
	for i := 0; i < data.Length(); i++ {
		v, _ := data.GetValue(i)
		button := widgets.NewButton(v, "Check")
		buttons[i] = button
	}

	delayButton := widget.NewButton("delay test", delayTest)
	Delay.Add(delayButton)
	vbox.Add(Delay)
	vbox.Add(container.NewGridWithColumns(2, buttons...))
	// lay := container.NewVBox(Delay, container.NewGridWithColumns(2, buttons...))
	// content := container.NewCenter(lay)
	// scroll := container.NewScroll(content)
	contentProxies.Add(vbox)
	return contentProxies
}

func delayTest() {
	proxies := api.GetProxies().All
	progress := widget.NewProgressBar()
	progress.Min = 0.0
	progress.Max = float64(len(proxies))
	Delay.Add(progress)
	Delay.Refresh()

	ch := make(chan map[string]string, len(proxies))
	for i, proxy := range proxies {
		go api.GetProxyDelayByName(proxy, ch)
		progress.SetValue(float64(i + 1))
	}

	buttons := make([]fyne.CanvasObject, 0, len(proxies))
	num := 15
	for proxy := range ch {
		for k, v := range proxy {
			fmt.Println(k, v)
			button := widgets.NewButton(k, v)
			buttons = append(buttons, button)
		}
		num--
		if num <= 0 {
			break
		}
	}

	contentProxies.Remove(vbox)
	vbox.Add(Delay)
	vbox.Add(container.NewGridWithColumns(2, buttons...))
	contentProxies.Add(vbox)
	// contentProxies.Refresh()

	time.Sleep(2 * time.Second)
	Delay.Remove(progress)
}
