/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdqueues

import (
	"fmt"
	"gorabbit/src/api"
	config "gorabbit/src/config"

	"github.com/spf13/cobra"
)

// getQueuesCmd represents the getQueues command
var getQueuesCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all queues from rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		queues := api.GetQueues(config.CurrentContext)
		if len(queues) == 0 {
			fmt.Println("No queues on '", config.CurrentContextName, "' context")
		} else {
			fmt.Println(queues)
		}
	},
}

func init() {
	queuesCmd.AddCommand(getQueuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getQueuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getQueuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
