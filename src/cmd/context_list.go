/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	config "gorabbit/src/config"

	"github.com/spf13/cobra"
)

// currentContextCmd represents the currentContext command
var listContextCmd = &cobra.Command{
	Use:   "list",
	Short: "List contexts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.ListContexts())
	},
}

func init() {
	contextCmd.AddCommand(listContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
