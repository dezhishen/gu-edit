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
	// centralwidget := widgets.NewQWidget(mainWindow, core.Qt__Widget)
	// centralwidget.SetObjectName("Centralwidget")
	gridLayout := widgets.NewQGridLayout(mainWindow)
	gridLayout.SetObjectName("gridLayout")
	gridLayout.SetContentsMargins(0, 0, 0, 0)
	gridLayout.SetSpacing(0)
	mainWindow.SetLayout(gridLayout)

	tabWidget := widgets.NewQTabWidget(mainWindow)
	tabWidget.SetObjectName("TabWidget")
	Tab := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	Tab.SetObjectName("Tab")
	tabWidget.AddTab(Tab, "tab1")
	Tab2 := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	Tab2.SetObjectName("Tab2")
	tabWidget.AddTab(Tab2, "tab2")
	//gridLayout.AddWidget3(tabWidget,1, 0, 1, 1, 0)
	mainWindow.SetCentralWidget(tabWidget)

	menubar := widgets.NewQMenuBar(mainWindow)
	menubar.SetObjectName("Menubar")
	menubar.Show()
	menubar.SetGeometry(core.NewQRect4(0, 0, 418, 26))
	menu := widgets.NewQMenu(menubar)
	menu.SetTitle("菜单1")
	menu.SetObjectName("file")
	menuAction := menu.AddAction("打开")
	menuAction.ConnectTriggered(func(checkBool bool) {
		if checkBool {
			msg := widgets.NewQMessageBox(mainWindow)
			msg.SetWindowTitle("yes comrade!")
			msg.SetText("im opening now!")
			msg.AddButton(widgets.NewQPushButton(msg),widgets.QMessageBox__AcceptRole)
			msg.Exec()

		}
		
	})
	menubar.AddMenu(menu)
	mainWindow.SetMenuBar(menubar)
	//gridLayout.SetMenuBar(menubar)
	//gridLayout.AddWidget(tabWidget)
	mainWindow.Show()
	widgets.QApplication_Exec()
}
