package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd *cobra.Command = &cobra.Command{
	Use:   "view",
	Short: "Views an application config",
	Long:  `view accepts an application name, reads your application configurations available in your etcaid directory and prints it.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		controller := buildController()
		if !controller.CheckInit() {
			cobra.CheckErr(errUninitialized)
		}

		err := controller.LoadApplications()
		cobra.CheckErr(err)

		app, err := controller.Application(name)
		cobra.CheckErr(err)

		fmt.Println("Title: " + app.Title)

		if len(app.HomePaths) > 0 {
			fmt.Printf("Home configuration files - %v:\n", controller.HomePath())
			for _, p := range app.HomePaths {
				fmt.Printf(" - %v\n", p)
			}
		}

		if len(app.XDGConfigPaths) > 0 {
			fmt.Printf("\nXDG configuration files - %v:\n", controller.XDGConfigPath())
			for _, p := range app.XDGConfigPaths {
				fmt.Printf(" - %v\n", p)
			}
		}
	},
}
