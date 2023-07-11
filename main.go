package main

import "golang_web/cmd"

// @title Go-Web开发案例
// @version 0.0.1
// @description 前端学Golang Web实战开发记录
func main() {
	defer cmd.Clean()
	cmd.Start()
}
