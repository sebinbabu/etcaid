package main

import (
	"github.com/spf13/cobra"
)

func init() {
}

var initCmd *cobra.Command = &cobra.Command{
	Use:   "init",
	Short: "Initializes etcaid",
	Long:  `Initializes the etcaid configuration and directories.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		controller := buildController()
		err := controller.Init()
		cobra.CheckErr(err)
	},
}
