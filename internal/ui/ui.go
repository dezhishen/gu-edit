package ui

import (
	"os"

	"github.com/therecipe/qt/core"
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
	menu.SetObjectName("file")
	menuAction := menu.AddAction("打开")
	menuAction.ConnectTrigger(func() {
		println("我出来了")
	})
	menubar.AddMenu(menu)
	menubar.Show()
	mainWindow.SetMenuBar(menubar)
	mainWindow.Show()
}
