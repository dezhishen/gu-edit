package ui

import (
	"github.com/dezhiShen/edit/pkg/menubar"
	"github.com/dezhiShen/edit/pkg/window"
)

// 渲染页面
func Render() error {
	uiWindow, err := window.Render()
	if err != nil {
		return err
	}
	uiMenuBar, err := menubar.Render(uiWindow.QWindow)
	if err != nil {
		return err
	}
	uiWindow.MenuBar = uiMenuBar
	menu1, _ := menubar.CreateMenu("打开", "click")
	uiWindow.MenuBar.AddMenu(menu1)
	// uiMenuBar.AddMenu(&menubar.UIMenu{})
	// uiTab, err := tab.Render(uiWindow.QWindow)
	// if err != nil {
	// 	return err
	// }
	// tab1 := widgets.NewQWidget(uiTab.QTabWidget, core.Qt__Widget)
	// tab1.SetObjectName("tab1")
	// uiTab.QTabWidget.AddTab(tab1, "tab1")
	// tab2 := widgets.NewQWidget(uiTab.QTabWidget, core.Qt__Widget)
	// tab2.SetObjectName("tab2")
	// uiTab.QTabWidget.AddTab(tab2, "tab2")
	// uiWindow.Tabs = uiTab
	//tab.Instance.AddTab()
	// uiStatusbar, err := statusbar.Render(uiWindow.QWindow)
	// if err != nil {
	// 	return err
	// }
	// uiWindow.StatusBar = uiStatusbar
	// 显示
	uiWindow.QWindow.Show()
	return nil
}
