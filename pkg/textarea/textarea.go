package textarea

import "github.com/therecipe/qt/widgets"

// New 创建一个
func New(parent widgets.QWidget_ITF) *widgets.QTextEdit {
	return widgets.NewQTextEdit(parent)
}
