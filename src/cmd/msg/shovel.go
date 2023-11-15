/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdmsg

import (
	"fmt"
	"gorabbit/src/api"
	config "gorabbit/src/config"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// Flags
var moveQueues []string
var contextNameIn string
var contextNameOut string

// moveMsgsCmd represents the moveMsgs command
var shovelCmd = &cobra.Command{
	Use:   "shovel",
	Short: "Move msgs from one rabbitmq to another",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Moving messages. A SIGINT (^C) can be called to interrupt the program gracefully...")

		forever := make(chan bool)

		contextOut := config.GetContext(contextNameOut)
		_, fromCh := api.ConnectAndGetRabbitChannel(contextOut)

		contextIn := config.GetContext(contextNameIn)
		_, toCh := api.ConnectAndGetRabbitChannel(contextIn)

		go captureSigint()

		if len(moveQueues) == 0 {
			moveQueues = api.GetQueues(contextOut)
		}
		for _, q := range moveQueues {
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
	msgCmd.AddCommand(shovelCmd)

	shovelCmd.Flags().StringVarP(&contextNameOut, "context-out", "o", "", "(required) specify which context will have its messages WITHDRAWN")
	shovelCmd.MarkFlagRequired("context-out")
	shovelCmd.Flags().StringVarP(&contextNameIn, "context-in", "i", "", "(required) specify which context will RECEIVE the messages")
	shovelCmd.MarkFlagRequired("context-in")
	shovelCmd.Flags().StringSliceVarP(&moveQueues, "queues", "q", []string{}, "Queues that will be shoveled")
}
