package main

func main() {
	etcaidCmd.AddCommand(initCmd, newCmd, backupCmd, restoreCmd, editCmd, listCmd)
	etcaidCmd.Execute()
}
