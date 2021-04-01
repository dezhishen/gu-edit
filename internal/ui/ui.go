package ui

import (
	"github.com/dezhiShen/edit/pkg/menubar"
	"github.com/dezhiShen/edit/pkg/statusbar"
	"github.com/dezhiShen/edit/pkg/tab"
	"github.com/dezhiShen/edit/pkg/window"
)

// 渲染页面
func Render() error {
	// TODO 渲染窗口
	err := window.Render()
	if err != nil {
		return err
	}
	// TODO 设置窗口
	err = menubar.Render(window.Instance.QWindow)
	if err != nil {
		return err
	}
	// TODO 创建一些菜单

	err = tab.Render(window.Instance.QWindow)
	if err != nil {
		return err
	}
	// TODO 初始化页签
	//tab.Instance.AddTab()
	err = statusbar.Render(window.Instance.QWindow)
	if err != nil {
		return err
	}
	// TODO 初始化状态栏
	// 显示
	window.Instance.QWindow.Show()
	return nil
}
