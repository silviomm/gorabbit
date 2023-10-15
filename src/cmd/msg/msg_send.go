/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdmsg

import (
	"gorabbit/src/api"
	config "gorabbit/src/config"
	"log"

	"github.com/spf13/cobra"
)

// Flags
var qName string

// moveMsgsCmd represents the moveMsgs command
var sendMsgCmd = &cobra.Command{
	Use:   "send",
	Short: "Send msg to a queue",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You must provide a msg as argument!")
		}
		_, rCh := api.ConnectAndGetRabbitChannel(config.CurrentContext)
		if _, err := api.GetQueue(rCh, qName); err != nil {
			log.Fatal("Error connecting to queue: ", qName)
		}
		api.SendMsg(qName, args[0], rCh)
	},
}

func init() {
	msgCmd.AddCommand(sendMsgCmd)

	sendMsgCmd.Flags().StringVarP(&qName, "queue", "q", "", "(required) specify for which queue it will send the message")
	sendMsgCmd.MarkFlagRequired("queue")
}
