package main

import "golang_web/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
