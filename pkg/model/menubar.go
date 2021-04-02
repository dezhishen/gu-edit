package model

import (
	"sync"

	"github.com/therecipe/qt/widgets"
)

type UIMenuBar struct {
	QMenubar *widgets.QMenuBar
	Menus    []*UIMenu
}

type UIMenu struct {
	QMenu   *widgets.QMenu
	MenuBar *UIMenuBar
	Index   int
}

var menuBarMenusLock sync.Mutex

// AddMenu 添加一个
func (menuBar *UIMenuBar) AddMenu(menu *UIMenu) error {
	//加锁
	menuBarMenusLock.Lock()
	defer menuBarMenusLock.Unlock()
	// dosomethings
	menu.QMenu = widgets.NewQMenu(menuBar.QMenubar)
	menu.MenuBar = menuBar
	menu.Index = len(menuBar.Menus)
	menuBar.Menus = append(menuBar.Menus, menu)
	return nil
}

// AddMenus 添加多个
func (menuBar *UIMenuBar) AddMenus(menus []*UIMenu) error {
	for _, v := range menus {
		e := menuBar.AddMenu(v)
		if e != nil {
			return e
		}
	}
	return nil
}

// Remove 从Menubar 中移除本菜单
func (menu *UIMenu) Remove() error {
	//加锁
	menuBarMenusLock.Lock()
	defer menuBarMenusLock.Unlock()
	menu.MenuBar.Menus = append(menu.MenuBar.Menus[:menu.Index], menu.MenuBar.Menus[menu.Index+1:]...)
	return nil
}

// GetQMenu 获取 QMenu 对象
func (menu *UIMenu) GetQMenu() *widgets.QMenu {
	return menu.QMenu
}
