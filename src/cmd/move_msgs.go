/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gorabbit/src/api"
	config "gorabbit/src/config"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// moveMsgsCmd represents the moveMsgs command
var moveMsgsCmd = &cobra.Command{
	Use:   "move-msgs",
	Short: "Move msgs from one rabbitmq to another",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Moving messages. A SIGINT (^C) can be called to interrupt the program gracefully...")

		forever := make(chan bool)

		_, fromCh := api.ConnectAndGetRabbitChannel(config.LocalRabbit())
		// toCh := api.ConnectAndGetRabbitChannel(config.RabbitMqConfig{Host: "localhost", User: "admin", Password: "admin", Port: "5673", AdminPort: "15673"})
		_, toCh := api.ConnectAndGetRabbitChannel(config.LocalRabbit())

		go captureSigint()

		api.ConsumeAndSend(fromCh, toCh, "teste1", "teste2")

		<-forever
	},
}

func captureSigint() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for sig := range c {
		fmt.Println("")
		fmt.Println(sig, "captured. Gracefully stopping application.")
		os.Exit(0)
	}
}

func init() {
	rootCmd.AddCommand(moveMsgsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moveMsgsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moveMsgsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
