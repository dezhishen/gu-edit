package main

import (
	"github.com/dezhiShen/edit/internal/setting"
	"github.com/dezhiShen/edit/internal/ui"
)

func init() {
	err := setting.LoadConf()
	if err != nil {
		panic(err)
	}
}

func main() {
	ui.Render()
}
