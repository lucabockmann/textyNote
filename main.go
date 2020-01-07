/*
	Author: @lucabockmann

	Shell function:

	function note() {
    	cd
    	cd <path to go workspace>
    	go run main.go $1 $2
    	cd
	}

	function tn() {
    	cd
    	cd <path to go workspace>
    	go run main.go $1 $2
    	cd
	}
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const DefaultEditor = "code"

func main() {
	CmdLineArgs := os.Args[1:]
	FileName := CmdLineArgs[0] + ".txt"

	if len(CmdLineArgs) < 2 {
		createFolder(FileName, "General")
	} else {
		createFolder(FileName, CmdLineArgs[1])
	}
}

func createFolder(FileName string, FolderName string) {
	if _, err := os.Stat(FolderName); os.IsNotExist(err) {
		fmt.Println("created folder: " + FolderName)
		os.Mkdir(FolderName, os.ModePerm)
		createFileInFolder(FileName, FolderName)
	} else {
		fmt.Println("add File " + FileName + " to Folder: " + FolderName)
		createFileInFolder(FileName, FolderName)
	}
}

func createFileInFolder(FileName string, FolderName string) {
	FolderPath, _ := filepath.Abs("./" + FolderName + "/" + FileName)

	os.Create(FolderPath)

	openFileInEditor(FileName, FolderName)
}

func openFileInEditor(FileName string, FolderName string) error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = DefaultEditor
	}

	executable, err := exec.LookPath(editor)

	if err != nil {
		return err
	}

	cmd := exec.Command(executable, FolderName+"/"+FileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
