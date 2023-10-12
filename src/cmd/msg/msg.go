/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdmsg

import (
	root "gorabbit/src/cmd"

	"github.com/spf13/cobra"
)

// queuesCmd represents the queues command
var msgCmd = &cobra.Command{
	Use:   "msg",
	Short: "Cmds to move, consume and send msgs",
}

func init() {
	root.RootCmd.AddCommand(msgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
