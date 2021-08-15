package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	noedit bool
	editor string

	newCmd *cobra.Command = &cobra.Command{
		Use:   "new",
		Short: "Generates a new application configuration and opens it for editing",
		Long: `new will create an application configuration in the etcaid directory with the name of the application specified.
It can be edited to add the paths of application configuration that will be backed up.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			controller := buildController()
			path, err := controller.Create(name)
			cobra.CheckErr(err)

			if noedit == false {
				c := exec.Command(editor, path)
				c.Stdin = os.Stdin
				c.Stdout = os.Stdout
				err = c.Run()
				cobra.CheckErr(err)
			}

			fmt.Println("Application config for", name, "created at", path)
		},
	}
)

func init() {
	newCmd.Flags().BoolVarP(&noedit, "noedit", "n", false, "disables starting the editor after creating application config")
	newCmd.Flags().StringVarP(&editor, "editor", "e", "vim", "command run when editing application config")
}
