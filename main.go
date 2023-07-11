package main

import (
	"fmt"
	"golang_web/cmd"
	"golang_web/utils"
)

// @title Go-Web开发案例
// @version 0.0.1
// @description 前端学Golang Web实战开发记录
func main() {
	defer cmd.Clean()
	cmd.Start()

	token, _ := utils.GenerateToken(1, "zs")
	fmt.Println(token)

	/*iJwtCostClaims, err := utils.ParseToken(token + "skfjksfjksjfk")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(iJwtCostClaims)*/

	fmt.Println(utils.IsTokenValid(token))
	fmt.Println(utils.IsTokenValid(token + "skfjksjfksjfk"))
}
