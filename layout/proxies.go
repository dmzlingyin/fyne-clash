package layout

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"clashG/api"
	"clashG/widgets"
)

var (
	Delay          = container.NewMax()
	vbox           = container.NewVBox()
	contentProxies = container.NewCenter()
)

func proxiesScreen() fyne.CanvasObject {
	proxies := api.GetProxies().All
	buttons := make([]fyne.CanvasObject, len(proxies))
	delayButton := widget.NewButton("delay test", delayTest)
	for i := 0; i < len(proxies); i++ {
		button := widgets.NewButton(proxies[i], "Check")
		buttons[i] = button
	}
	Delay.Add(delayButton)
	vbox.Add(Delay)
	vbox.Add(container.NewGridWithColumns(2, buttons...))
	// lay := container.NewVBox(Delay, container.NewGridWithColumns(2, buttons...))
	// content := container.NewCenter(lay)
	// scroll := container.NewScroll(content)
	contentProxies.Add(vbox)
	return content
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
	// for proxy := range ch {
	// 	for k, v := range proxy {
	// 		fmt.Println(k, v)
	// 		// button := widgets.NewButton(k, v)
	// 		// buttons = append(buttons, button)
	// 	}
	// }
	select {
	case name := <-ch:
		for k, v := range name {
			button := widgets.NewButton(k, v)
			buttons = append(buttons, button)
		}
	default:
		contentProxies.Remove(vbox)
		vbox.Add(Delay)
		vbox.Add(container.NewGridWithColumns(2, buttons...))
		contentProxies.Add(vbox)
	}

	time.Sleep(2 * time.Second)
	Delay.Remove(progress)
}
