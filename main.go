/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"gorabbit/src/cmd"
	_ "gorabbit/src/cmd/context"
	_ "gorabbit/src/cmd/msg"
	_ "gorabbit/src/cmd/queues"
	config "gorabbit/src/config"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
