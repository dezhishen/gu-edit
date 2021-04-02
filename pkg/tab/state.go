package tab

import (
	"github.com/dezhiShen/edit/pkg/model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

// Instance 全局唯一对象
var Instance *model.UITabWidget

func Render(mainWindow *widgets.QMainWindow) (*model.UITabWidget, error) {
	if Instance != nil {
		return Instance, nil
	}
	centralwidget := widgets.NewQWidget(mainWindow, core.Qt__Widget)
	centralwidget.SetObjectName("Centralwidget")
	gridLayout := widgets.NewQGridLayout(centralwidget)
	gridLayout.SetObjectName("gridLayout")
	gridLayout.SetContentsMargins(0, 0, 0, 0)
	gridLayout.SetSpacing(0)
	tabWidget := widgets.NewQTabWidget(centralwidget)
	Instance = &model.UITabWidget{
		QTabWidget: tabWidget,
	}
	return Instance, nil
}
