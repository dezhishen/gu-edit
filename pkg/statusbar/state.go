package statusbar

import (
	"github.com/therecipe/qt/widgets"
)

var Instance *UIStatusBar

func Render(mainWindow *widgets.QMainWindow) (*UIStatusBar, error) {
	if Instance == nil {
		Instance = &UIStatusBar{}
		Instance.QStatusBar = widgets.NewQStatusBar(mainWindow)
		label := widgets.NewQLabel(mainWindow, 0)
		label.SetText("我是你爹")
		Instance.QStatusBar.AddWidget(label, 0)
		// TODO 初始化statusbar
	}
	return Instance, nil
}
