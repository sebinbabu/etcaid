package main

import (
	"github.com/spf13/cobra"
)

var restoreCmd *cobra.Command = &cobra.Command{
	Use:   "restore",
	Short: "Restores all applications",
	Long:  `restore will read your application configurations and restore the paths from the etcaid directory.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		controller := buildController()
		if !controller.CheckInit() {
			cobra.CheckErr(errUninitialized)
		}

		err := controller.LoadApplications()
		cobra.CheckErr(err)

		controller.RestoreAll()
	},
}
