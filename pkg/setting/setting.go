package setting

import "github.com/spf13/viper"

var allConf = make(map[string]*viper.Viper)

const (
	path = "./conf"
)

// Load 加载配置文件
func Load(name string) error {
	temp := viper.New()
	temp.SetConfigName(name)
	temp.AddConfigPath(path)
	temp.SetConfigType("yaml")
	err := temp.ReadInConfig()
	if err != nil {
		return err
	}
	allConf[name] = temp
	return nil
}

// GetString GetString
func GetString(name, key string) string {
	v, ok := allConf[name]
	if !ok {
		return ""
	}
	return v.GetString(key)
}

// GetBool GetBool
func GetBool(name, key string) bool {
	v, ok := allConf[name]
	if !ok {
		return false
	}
	return v.GetBool(key)
}

// GetInt GetInt
func GetInt(name, key string) int {
	v, ok := allConf[name]
	if !ok {
		return 0
	}
	return v.GetInt(key)
}

// GetInt GetInt
func GetFloat64(name, key string) float64 {
	v, ok := allConf[name]
	if !ok {
		return 0
	}
	return v.GetFloat64(key)
}

// GetIntSlice GetIntSlice
func GetIntSlice(name, key string) []int {
	v, ok := allConf[name]
	if !ok {
		return nil
	}
	return v.GetIntSlice(key)
}

// GetIntSlice GetIntSlice
func GetStringSlice(name, key string) []string {
	v, ok := allConf[name]
	if !ok {
		return nil
	}
	return v.GetStringSlice(key)
}

// GetStringMapString GetStringMapString
func GetStringMapString(name, key string) map[string]string {
	v, ok := allConf[name]
	if !ok {
		return nil
	}
	return v.GetStringMapString(key)
}

func Set(name, key string, value interface{}) error {
	temp, ok := allConf[name]
	if !ok {
		temp = viper.New()
		temp.SetConfigName(name)
		temp.AddConfigPath(path)
		temp.SetConfigType("yaml")
		allConf[name] = temp
	}
	temp.Set(key, value)
	return temp.SafeWriteConfig()
}
