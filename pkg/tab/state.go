package tab

import "github.com/therecipe/qt/widgets"

// Instance 全局唯一对象
var Instance *UITabWidget

func Render(mainWindow *widgets.QMainWindow) error {
	tabWidget := widgets.NewQTabWidget(mainWindow)
	Instance = &UITabWidget{
		QTabWidget: tabWidget,
	}
	// TODO 给Instanace赋值
	return nil
}
