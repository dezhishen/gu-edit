package model

import (
	"github.com/therecipe/qt/widgets"
)

type UIMainWindow struct {
	Tabs      *UITabWidget
	MenuBar   *UIMenuBar
	StatusBar *UIStatusBar
	QWindow   *widgets.QMainWindow
}
