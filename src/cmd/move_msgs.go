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

/*
todo:
- specify the rate
- dispose connections on ^C
- declare queue on IN rabbit if it doesn't exists
- better output to show messages consumed/sent
*/

// Flags
var contextNameIn string
var contextNameOut string

// moveMsgsCmd represents the moveMsgs command
var moveMsgsCmd = &cobra.Command{
	Use:   "move-msgs",
	Short: "Move msgs from one rabbitmq to another",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Moving messages. A SIGINT (^C) can be called to interrupt the program gracefully...")

		forever := make(chan bool)

		contextOut := config.GetContext(contextNameOut)
		_, fromCh := api.ConnectAndGetRabbitChannel(contextOut)

		contextIn := config.GetContext(contextNameIn)
		_, toCh := api.ConnectAndGetRabbitChannel(contextIn)

		go captureSigint()

		outQueues := api.GetQueues(contextOut)
		for _, q := range outQueues {
			go api.ConsumeAndSend(fromCh, toCh, q, q)
		}

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

	moveMsgsCmd.Flags().StringVarP(&contextNameOut, "context-out", "o", "", "(required) specify which context will have its messages WITHDRAWN")
	moveMsgsCmd.MarkFlagRequired("context-out")
	moveMsgsCmd.Flags().StringVarP(&contextNameIn, "context-in", "i", "", "(required) specify which context will RECEIVE the messages")
	moveMsgsCmd.MarkFlagRequired("context-in")
}
