package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/dezhiShen/edit/internal/fileutil"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// Render 渲染页面
func Render() {
	widgets.NewQApplication(len(os.Args), os.Args)
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetObjectName("MainWindow")
	mainWindow.SetGeometry(core.NewQRect4(400, 200, 800, 600))
	gridLayout := widgets.NewQGridLayout(mainWindow)
	gridLayout.SetObjectName("gridLayout")
	gridLayout.SetContentsMargins(0, 0, 0, 0)
	gridLayout.SetSpacing(0)
	mainWindow.SetLayout(gridLayout)
	tabWidget := widgets.NewQTabWidget(mainWindow)
	tabWidget.SetObjectName("TabWidget")
	mainWindow.SetCentralWidget(tabWidget)
	menubar := widgets.NewQMenuBar(mainWindow)
	menubar.SetObjectName("Menubar")
	menubar.Show()
	menubar.SetGeometry(core.NewQRect4(0, 0, 418, 26))
	menu := widgets.NewQMenu(menubar)
	menu.SetTitle("文件")
	menu.SetObjectName("file")
	menuAction := menu.AddAction("打开")
	menuAction.ConnectTriggered(func(checkBool bool) {
		if !checkBool {
			path := widgets.QFileDialog_GetOpenFileName(mainWindow, "打开文本", "", "", "", widgets.QFileDialog__DontUseNativeDialog)
			index := strings.LastIndex(path, "/")
			fileName := path[index+1:]
			content, _ := fileutil.Read(path)
			textEdit := widgets.NewQPlainTextEdit2(string(content), tabWidget)
			textEdit.SetObjectName(fileName)
			textEdit.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
				if event.Modifiers() == 0 {
					switch event.Key() {
					case int(core.Qt__Key_Backspace):
						if textEdit.TextCursor().SelectedText() == "" {
							textEdit.TextCursor().DeletePreviousChar()
						} else {
							textEdit.TextCursor().RemoveSelectedText()
						}
					default:
						textEdit.InsertPlainText(event.Text())
					}
				} else if event.Modifiers() == core.Qt__ControlModifier && event.Key() == int(core.Qt__Key_S) {
					//presse
					fileutil.Save(path, []byte(textEdit.ToPlainText()))
				} else if event.Modifiers() == core.Qt__ShiftModifier {
					switch event.Key() {
					case int(core.Qt__Key_Backspace):
						if textEdit.TextCursor().SelectedText() == "" {
							textEdit.TextCursor().DeletePreviousChar()
						} else {
							textEdit.TextCursor().RemoveSelectedText()
						}
					default:
						textEdit.InsertPlainText(event.Text())
					}
				}
			})
			tabWidget.AddTab(textEdit, fileName)
			fmt.Println()
		}

	})
	menubar.AddMenu(menu)
	mainWindow.SetMenuBar(menubar)
	mainWindow.Show()
	widgets.QApplication_Exec()
}
