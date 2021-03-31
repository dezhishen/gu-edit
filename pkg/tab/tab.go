package tab

import (
	"github.com/dezhiShen/edit/pkg/textarea"
	"github.com/therecipe/qt/widgets"
)

type UITabWidget struct {
	TabWidget *widgets.QTabWidget
	Tabs      []*widgets.QWidget
	ActiveTab *Tab
}

type Tab struct {
	Widget   *widgets.QWidget
	TextArea *textarea.UITextArea
}
