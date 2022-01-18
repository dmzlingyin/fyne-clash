package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"clashG/api"
	"clashG/widgets"
)

func proxiesScreen() fyne.CanvasObject {
	proxies := api.GetProxies().All
	buttons := make([]fyne.CanvasObject, len(proxies))
	for i := 0; i < len(proxies); i++ {
		button := widgets.NewButton(proxies[i], "Check")
		buttons[i] = button
	}
	content := container.NewCenter(container.NewGridWithColumns(2, buttons...))
	scroll := container.NewScroll(content)

	return scroll
}
