/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package queues

import (
	"fmt"
	"gorabbit/src/api"
	config "gorabbit/src/config"

	"github.com/spf13/cobra"
)

var queuesToDelete []string

// deleteQueuesCmd represents the deleteQueues command
var deleteQueuesCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete queues from rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		_, ch := api.ConnectAndGetRabbitChannel(config.CurrentContext)

		if len(queuesToDelete) == 0 {
			queuesToDelete = api.GetQueues(config.CurrentContext)
		}

		if len(queuesToDelete) <= 0 {
			fmt.Println("No queues to delete")
		} else {
			api.DeleteQueues(ch, queuesToDelete)
		}
	},
}

func init() {
	queuesCmd.AddCommand(deleteQueuesCmd)
	deleteQueuesCmd.Flags().StringSliceVarP(&queuesToDelete, "queues", "q", []string{}, "Queues that will be deleted")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteQueuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteQueuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
