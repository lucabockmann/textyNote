package main

import (
	"os"
	"os/exec"
)

const DefaultEditor = "code"

func main() {
	CmdLineArgs := os.Args[1:]

	if len(CmdLineArgs) > 0 {
		folderName := CmdLineArgs[1]
		fileName := CmdLineArgs[0] + ".txt"

		openFile(fileName, folderName)
	}
}

func openFile(fileName string, folderName string) error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = DefaultEditor
	}

	executable, err := exec.LookPath(editor)

	if err != nil {
		return err
	}

	cmd := exec.Command(executable, folderName+"/"+fileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
