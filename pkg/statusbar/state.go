package statusbar

import "github.com/therecipe/qt/widgets"

var Instance *UIStatusBar

func Render(mainWindow *widgets.QMainWindow) error {
	Instance = &UIStatusBar{}
	// TODO 初始化statusbar
	return nil
}
