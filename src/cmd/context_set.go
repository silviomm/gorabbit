/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	config "gorabbit/src/config"
	"log"

	"github.com/spf13/cobra"
)

// currentContextCmd represents the currentContext command
var setContextCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current context",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Must provide context name")
		}
		if !config.ContextExists(args[0]) {
			log.Fatalf("Context '%s' not found", args[0])
		}
		config.SetCurrentContext(args[0])
	},
}

func init() {
	contextCmd.AddCommand(setContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
