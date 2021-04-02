package ui

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// 渲染页面
func Render() {
	widgets.NewQApplication(len(os.Args), os.Args)
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetObjectName("MainWindow")
	mainWindow.SetGeometry(core.NewQRect4(0, 0, 418, 483))
	centralwidget := widgets.NewQWidget(mainWindow, core.Qt__Widget)
	centralwidget.SetObjectName("Centralwidget")
	gridLayout := widgets.NewQGridLayout(centralwidget)
	gridLayout.SetObjectName("gridLayout")
	gridLayout.SetContentsMargins(0, 0, 0, 0)
	gridLayout.SetSpacing(0)
	tabWidget := widgets.NewQTabWidget(centralwidget)
	tabWidget.SetObjectName("TabWidget")
	Tab := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	Tab.SetObjectName("Tab")
	tabWidget.AddTab(Tab, "")
	Tab2 := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	Tab2.SetObjectName("Tab2")
	tabWidget.AddTab(Tab2, "")
	gridLayout.AddWidget3(tabWidget, 0, 0, 1, 1, 0)
	mainWindow.SetCentralWidget(centralwidget)
	menubar := widgets.NewQMenuBar(mainWindow)
	menubar.SetObjectName("Menubar")
	menubar.SetGeometry(core.NewQRect4(0, 0, 418, 26))
	menu := widgets.NewQMenu(menubar)
	menu.SetObjectName("Menu")
	menu.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
	})
	menu2 := widgets.NewQMenu(menubar)
	menu2.SetObjectName("Menu2")
	menu3 := widgets.NewQMenu(menubar)
	menu3.SetObjectName("Menu3")
	menu4 := widgets.NewQMenu(menubar)
	menu4.SetObjectName("Menu4")
	mainWindow.SetMenuBar(menubar)
	statusbar := widgets.NewQStatusBar(mainWindow)
	statusbar.SetObjectName("Statusbar")
	mainWindow.SetStatusBar(statusbar)
	menubar.QWidget.AddAction(menu.MenuAction())
	// if menu.MenuAction() != nil {
	// }
	// if menu2.MenuAction() != nil {
	// 	menubar.QWidget.AddAction(menu2.MenuAction())
	// }
	// if menu3.MenuAction() != nil {
	// 	menubar.QWidget.AddAction(menu3.MenuAction())
	// }
	// if menu4.MenuAction() != nil {
	// 	menubar.QWidget.AddAction(menu4.MenuAction())
	// }
	_translate := core.QCoreApplication_Translate
	mainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	tabWidget.SetTabText(tabWidget.IndexOf(Tab), _translate("MainWindow", "Tab 1", "", -1))
	tabWidget.SetTabText(tabWidget.IndexOf(Tab2), _translate("MainWindow", "Tab 2", "", -1))
	menu.SetTitle(_translate("MainWindow", "文件", "", -1))
	menu2.SetTitle(_translate("MainWindow", "编辑", "", -1))
	menu3.SetTitle(_translate("MainWindow", "格式化", "", -1))
	menu4.SetTitle(_translate("MainWindow", "设置", "", -1))
	mainWindow.Show()
}
