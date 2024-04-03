package main

import (
	"flag"
	"fmt"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	var saveFlag = flag.String("save", "", "Generate and save a password for the provided name")
	var findFlag = flag.String("find", "", "Find names matching the provided substring")

	flag.Parse()
	//todo: enforce one command line arg (print usage)

	fmt.Print("Password: ")
	pwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Failed to read password:", err)
		return
	}

	kek, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to generate KEK:", err)
		return
	}

	//todo: hash 480 bit kek or switch to different KDF

	if *saveFlag != "" {
		savePassword(kek, *saveFlag)
	} else if *findFlag != "" {
		findPasswords(kek, *findFlag)
	} else {
		flag.PrintDefaults()
	}
}

func savePassword(kek []byte, name string) {
	fmt.Println("KEK:", kek, "Save:", name)
}

func findPasswords(kek []byte, nameSubstring string) {
	fmt.Println("KEK", kek, "Find:", nameSubstring)
}
