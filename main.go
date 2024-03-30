package main

import (
	"fmt"
	"flag"
)

func main() {
	var saveFlag = flag.String("save", "", "Generate and save a password for the provided name")
	var findFlag = flag.String("find", "", "Find names matching the provided substring")

	flag.Parse()

	if *saveFlag != "" {
		savePassword(*saveFlag)
	}

	if *findFlag != "" {
		findPasswords(*findFlag)
	}
}

func savePassword(name string) {
	fmt.Println("Save:", name)
}

func findPasswords(nameSubstring string) {
	fmt.Println("Find:", nameSubstring)
}

