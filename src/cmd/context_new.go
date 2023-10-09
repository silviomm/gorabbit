/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	config "gorabbit/src/config"
	prompts "gorabbit/src/prompt"
	"log"

	"github.com/spf13/cobra"
)

// currentContextCmd represents the currentContext command
var newContextCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new context",
	Run: func(cmd *cobra.Command, args []string) {
		result := prompts.PromptContext()
		if result == nil {
			log.Fatal("Error creating new context")
		}

		config.SetContext(result.Name, config.RabbitMqConfig{
			Host:      result.Host,
			Port:      result.Port,
			AdminPort: result.AdminPort,
			User:      result.User,
			Password:  result.Password})
	},
}

func init() {
	contextCmd.AddCommand(newContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
