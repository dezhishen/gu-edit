package setting

import "github.com/dezhiShen/edit/pkg/setting"

// LoadConf 加载配置文件
func LoadConf() error {
	err := setting.Load("system")
	if err != nil {
		return err
	}
	return nil
}
