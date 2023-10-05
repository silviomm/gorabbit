/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gorabbit/src/api"
	config "gorabbit/src/config"

	"github.com/spf13/cobra"
)

// deleteQueuesCmd represents the deleteQueues command
var deleteQueuesCmd = &cobra.Command{
	Use:   "delete-queues",
	Short: "Delete all queues on rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		_, ch := api.ConnectAndGetRabbitChannel(config.LocalRabbit())
		queues := api.GetQueues(config.LocalRabbit())
		api.DeleteQueues(ch, queues)
	},
}

func init() {
	rootCmd.AddCommand(deleteQueuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteQueuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteQueuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
