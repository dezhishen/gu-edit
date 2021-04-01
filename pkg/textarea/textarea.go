package textarea

import (
	"github.com/dezhiShen/edit/internal/fileutil"
	"github.com/therecipe/qt/widgets"
)

type UITextArea struct {
	// 文件路径
	Path string
	// 编辑器对象
	QTextArea *widgets.QTextEdit
}

//初始化
func Init(path string) (*UITextArea, error) {
	return nil, nil
}

// 保存到文件
func (u *UITextArea) Save() error {
	fileutil.Save(u.Path, nil)
	return nil
}

// 重新从文件加载
func (u *UITextArea) Reload() error {
	_, e := fileutil.Read(u.Path)
	return e
}
