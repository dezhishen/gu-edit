package menubar

import "github.com/therecipe/qt/widgets"

// Instance 全局唯一对象
var Instance *UIMenuBar

func Render(mainWindow *widgets.QMainWindow) error {
	menuBar := widgets.NewQMenuBar(mainWindow)
	Instance = &UIMenuBar{
		QMenubar: menuBar,
	}
	// TODO 给Instanace赋值
	return nil
}

func AddMenu(menu *UIMenu) error {
	return Instance.AddMenu(menu)
}

func CreateMenu(name string, action string) (*UIMenu, error) {
	// TODO 梳理事件,完善事件逻辑
	menu := &UIMenu{
		QMenu:   widgets.NewQMenu(Instance.QMenubar),
		MenuBar: Instance,
	}
	menu.QMenu.SetObjectName(name)
	menu.QMenu.AddAction(action)
	return menu, nil
}
func CreateMenuAndAdd(name string, action string) (*UIMenu, error) {
	// TODO 梳理事件,完善事件逻辑
	menu := &UIMenu{
		QMenu:   widgets.NewQMenu(Instance.QMenubar),
		MenuBar: Instance,
	}
	AddMenu(menu)
	return menu, nil
}
