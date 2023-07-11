package cmd

import (
	"fmt"
	"golang_web/conf"
	"golang_web/router"
)

func Start() {
	fmt.Println("========Start========")
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean========")
}
