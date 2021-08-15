package main

func main() {
	etcaidCmd.AddCommand(initCmd, newCmd, backupCmd, restoreCmd, editCmd, listCmd, viewCmd)
	etcaidCmd.Execute()
}
