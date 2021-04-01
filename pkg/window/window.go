package window

import (
	"github.com/dezhiShen/edit/pkg/menubar"
	"github.com/dezhiShen/edit/pkg/statusbar"
	"github.com/dezhiShen/edit/pkg/tab"

	"github.com/therecipe/qt/widgets"
)

type UIMainWindow struct {
	Tabs      *tab.UITabWidget
	MenuBar   *menubar.UIMenuBar
	StatusBar *statusbar.UIStatusBar
	QWindow   *widgets.QMainWindow
	// QCentralwidget *widgets.QWidget
	// QGridLayout    *widgets.QGridLayout
}
