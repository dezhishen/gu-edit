package menubar

import (
	"github.com/dezhiShen/edit/pkg/model"
	"github.com/therecipe/qt/widgets"
)

// Instance 全局唯一对象
var Instance *model.UIMenuBar

func Render(mainWindow *widgets.QMainWindow) (*model.UIMenuBar, error) {
	if Instance != nil {
		return Instance, nil
	}
	menuBar := widgets.NewQMenuBar(mainWindow)
	Instance = &model.UIMenuBar{
		QMenubar: menuBar,
	}
	// TODO 给Instanace赋值
	return Instance, nil
}

func AddMenu(menu *model.UIMenu) error {
	return Instance.AddMenu(menu)
}

func CreateMenu(name string, action string) (*model.UIMenu, error) {
	// TODO 梳理事件,完善事件逻辑
	menu := &model.UIMenu{
		QMenu:   widgets.NewQMenu(Instance.QMenubar),
		MenuBar: Instance,
	}
	menu.QMenu.SetObjectName(name)
	menu.QMenu.AddAction(action)
	return menu, nil
}
func CreateMenuAndAdd(name string, action string) (*model.UIMenu, error) {
	// TODO 梳理事件,完善事件逻辑
	menu := &model.UIMenu{
		QMenu:   widgets.NewQMenu(Instance.QMenubar),
		MenuBar: Instance,
	}
	AddMenu(menu)
	return menu, nil
}
