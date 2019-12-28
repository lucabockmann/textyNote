/*
	Author: @lucabockmann

	Shell function:

	function note() {
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
	"path/filepath"
)

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
}
