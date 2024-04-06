package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	var saveFlag = flag.String("save", "", "Generate and save a password for the provided name")
	var findFlag = flag.String("find", "", "Find names matching the provided substring")

	flag.Parse()
	if flag.NFlag() != 1 {
		log.Println("Must be called with one and only one valid input argument")
		flag.PrintDefaults()
		return
	}

	fmt.Print("Password: ")
	pwd, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatalln("Failed to read password:", err)
	}

	kek, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Failed to generate KEK:", err)
	}

	//todo: hash 480 bit kek or switch to different KDF

	if *saveFlag != "" {
		savePassword(kek, *saveFlag)
	} else {
		findPasswords(kek, *findFlag)
	}
}

func savePassword(kek []byte, name string) {
	fmt.Println("KEK:", kek, "Save:", name)
}

func findPasswords(kek []byte, nameSubstring string) {
	fmt.Println("KEK", kek, "Find:", nameSubstring)
}
