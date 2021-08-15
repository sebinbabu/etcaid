package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd *cobra.Command = &cobra.Command{
	Use:   "list",
	Short: "Lists all applications available",
	Long:  `list will read your application configurations and list all applications available in your etcaid directory.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		controller := buildController()
		if !controller.CheckInit() {
			cobra.CheckErr(errUninitialized)
		}

		err := controller.LoadApplications()
		cobra.CheckErr(err)

		apps := controller.Applications()
		for _, a := range apps {
			fmt.Printf(" - %v\n", a.Name)
		}

		fmt.Println()
		fmt.Printf("%d applications are available\n", len(apps))
	},
}
