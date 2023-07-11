package cmd

import (
	"fmt"
	"golang_web/conf"
)

func Start() {
	fmt.Println("========Start========")
	conf.InitConfig()
}

func Clean() {
	fmt.Println("========Clean========")
}
