package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	CmdLineArgs := os.Args[1:]
	FileName := CmdLineArgs[0] + ".txt"

	fmt.Println(CmdLineArgs)

	if len(CmdLineArgs) < 2 {
		if _, err := os.Stat("./General"); os.IsNotExist(err) {
			os.Mkdir("General", os.ModePerm)
			createFileInGeneral(FileName)
		} else {
			createFileInGeneral(FileName)
		}
	}
}

func createFileInGeneral(FileName string) {
	GeneralPath, _ := filepath.Abs("./General/" + FileName)
	os.Create(GeneralPath)
}
