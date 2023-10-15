/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdqueues

import (
	"fmt"
	"gorabbit/src/api"
	config "gorabbit/src/config"
	"log"

	"github.com/spf13/cobra"
)

// deleteQueuesCmd represents the deleteQueues command
var newQueuesCmd = &cobra.Command{
	Use:   "new",
	Short: "Create queue on rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Must provide queue name as argument")
		}
		qName := args[0]
		_, ch := api.ConnectAndGetRabbitChannel(config.CurrentContext)
		_, err := api.CreateQueue(ch, qName)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(qName, "created successfully")
	},
}

func init() {
	queuesCmd.AddCommand(newQueuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteQueuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteQueuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
