package main

import (
	"github.com/spf13/cobra"
)

var etcaidCmd *cobra.Command = &cobra.Command{
	Use:   "etcaid",
	Short: "Backup & manage your application configurations for fun and profit",
	Long: `
etcaid is a helps you backup your application configurations.
Backups can be synced externally using tools like git and restored across multiple devices.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
