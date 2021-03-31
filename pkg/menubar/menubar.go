package menubar

import (
	"sync"

	"github.com/therecipe/qt/widgets"
)

type UIMenuBar struct {
	menubar *widgets.QMenuBar
	Menus   []UIMenu
}

type UIMenu struct {
	menu    *widgets.QMenu
	MenuBar *UIMenuBar
	Index   int
}

var menuBarMenusLock sync.Mutex

// 添加一个
func (menuBar *UIMenuBar) AddMenu(menu UIMenu) error {
	//加锁
	menuBarMenusLock.Lock()
	defer menuBarMenusLock.Unlock()
	// dosomethings
	menu.menu = widgets.NewQMenu(menuBar.menubar)
	menu.MenuBar = menuBar
	menu.Index = len(menuBar.Menus)
	menuBar.Menus = append(menuBar.Menus, menu)
	return nil
}

// 添加多个
func (menuBar *UIMenuBar) AddMenus(menus []UIMenu) error {
	for _, v := range menus {
		e := menuBar.AddMenu(v)
		if e != nil {
			return e
		}
	}
	return nil
}

// 移除菜单
func (menu *UIMenu) Remove() error {
	//加锁
	menuBarMenusLock.Lock()
	defer menuBarMenusLock.Unlock()
	menu.MenuBar.Menus = append(menu.MenuBar.Menus[:menu.Index], menu.MenuBar.Menus[menu.Index+1:]...)
	return nil
}
