package tab

import (
	"github.com/dezhiShen/edit/pkg/textarea"
	"github.com/therecipe/qt/widgets"
)

type UITabWidget struct {
	QTabWidget *widgets.QTabWidget
	Tabs       []*widgets.QWidget
	ActiveTab  *Tab
}

type Tab struct {
	QWidget  *widgets.QWidget
	TextArea *textarea.UITextArea
}
