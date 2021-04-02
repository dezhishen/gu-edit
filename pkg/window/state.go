package window

import (
	"os"

	"github.com/dezhiShen/edit/pkg/model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var Instance *model.UIMainWindow

func Render() (*model.UIMainWindow, error) {
	// TODO 初始化
	if Instance != nil {
		return Instance, nil
	}
	Instance = &model.UIMainWindow{}
	widgets.NewQApplication(len(os.Args), os.Args)
	qWindow := widgets.NewQMainWindow(nil, 0)
	qWindow.SetObjectName("MainWindow")
	qWindow.SetWindowTitle("咕编辑器")
	qWindow.SetGeometry(core.NewQRect4(0, 0, 418, 483))
	Instance.QWindow = qWindow
	return Instance, nil
}
