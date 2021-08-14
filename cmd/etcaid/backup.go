package main

import (
	"github.com/spf13/cobra"
)

var backupCmd *cobra.Command = &cobra.Command{
	Use:   "backup",
	Short: "Backups all applications",
	Long: `
backup will read your application configurations and backup the paths mentioned the configuration to the etcaid directory.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		controller := buildController()

		err := controller.LoadApplications()
		cobra.CheckErr(err)

		controller.BackupAll()
	},
}
