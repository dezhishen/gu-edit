package window

import (
	"os"

	"github.com/dezhiShen/edit/pkg/menubar"
	"github.com/dezhiShen/edit/pkg/statusbar"
	"github.com/dezhiShen/edit/pkg/tab"

	"github.com/therecipe/qt/widgets"
)

var wind *widgets.QMainWindow

type UIMainWindow struct {
	UITabs    *tab.UITabWidget
	MenuBar   *menubar.UIMenuBar
	StatusBar *statusbar.UIStatusBar
}

func InitWindow() (*widgets.QMainWindow, error) {
	if wind == nil {
		widgets.NewQApplication(len(os.Args), os.Args)
		window := widgets.NewQMainWindow(nil, 0)
		gltout := widgets.NewQGridLayout(window)
		gltout.SetContentsMargins(10, 20, 10, 20)
		window.SetLayout(gltout)
		wind = widgets.NewQMainWindow(nil, 0)
	}
	return wind, nil

}
