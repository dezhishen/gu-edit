package window

var Instance *UIMainWindow

func Render() error {
	// TODO 初始化
	Instance = &UIMainWindow{}
	return nil
}
