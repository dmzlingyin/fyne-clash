package widgets

import (
	"clashG/api"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MyButton 继承widget.Button用于个性化定制
type MyButton struct {
	widget.Button
}

// 重写Tapped方法, 捕获点击事件的同时,获取button的属性信息并执行一系列操作
func (b *MyButton) Tapped(*fyne.PointEvent) {
	if b.Disabled() {
		return
	}

	//使用goroutine进行代理延迟测试, 否则将阻塞事件循环, 严重会卡死窗口
	// go api.GetProxyDelayByName(b.Text)
	proxy := strings.Split(b.Text, "=")[0]
	delay := api.GetProxyDelayByName(proxy)
	b.Text = proxy + "=" + delay
	api.ChangeProxyByName(proxy)

	b.Refresh()
}

// MyButton的工厂方法
func NewButton(label, delay string) *MyButton {
	button := &MyButton{}
	button.Text = label + "=" + delay
	button.Alignment = 1

	button.ExtendBaseWidget(button)
	return button
}
