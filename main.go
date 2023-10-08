/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"gorabbit/src/cmd"
	config "gorabbit/src/config"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
