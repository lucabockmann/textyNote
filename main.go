package main

import (
	"fmt"
	"os"
)

func main() {
	CmdLineArgs := os.Args[1:]

	fmt.Println(CmdLineArgs)
}
