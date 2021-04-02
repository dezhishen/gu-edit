package model

import (
	"github.com/therecipe/qt/widgets"
)

type UITabWidget struct {
	QTabWidget *widgets.QTabWidget
	Tabs       []*widgets.QWidget
	ActiveTab  *Tab
}

type Tab struct {
	QWidget  *widgets.QWidget
	TextArea *UITextArea
}
