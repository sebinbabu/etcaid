package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	editCmd *cobra.Command = &cobra.Command{
		Use:   "edit",
		Short: "Opens an editor to edit an existing application configuration",
		Long: `edit will open an existing application configuration in the etcaid directory with the name of the application specified.
It can be edited to add the paths of application configuration that will be backed up.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			controller := buildController()
			if !controller.CheckInit() {
				cobra.CheckErr(errUninitialized)
			}

			path, err := controller.ApplicationConfigPath(name)
			cobra.CheckErr(err)

			c := exec.Command(editor, path)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			err = c.Run()
			cobra.CheckErr(err)

			fmt.Println("Application config for", name, "edited at", path)
		},
	}
)

func init() {
	editCmd.Flags().StringVarP(&editor, "editor", "e", "vim", "command run when editing application config")
}
