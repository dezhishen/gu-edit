package window

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var Instance *UIMainWindow

func Render() error {
	// TODO 初始化
	Instance = &UIMainWindow{}
	widgets.NewQApplication(len(os.Args), os.Args)
	qWindow := widgets.NewQMainWindow(nil, 0)
	qWindow.SetObjectName("MainWindow")
	qWindow.SetGeometry(core.NewQRect4(0, 0, 418, 483))
	Instance.QWindow = qWindow
	return nil
}
