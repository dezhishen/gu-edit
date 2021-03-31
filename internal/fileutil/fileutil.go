package fileutil

import (
	"io/ioutil"
	"os"
)

// Read 读取
func Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Save 保存
func Save(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}
