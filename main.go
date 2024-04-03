package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var saveFlag = flag.String("save", "", "Generate and save a password for the provided name")
	var findFlag = flag.String("find", "", "Find names matching the provided substring")

	flag.Parse()

	pwd := "password123" //todo: prompt user for password

	kek, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to generate KEK: %v\n", err)
	}

	//todo: hash 480 bit kek or switch to different KDF

	if *saveFlag != "" {
		savePassword(kek, *saveFlag)
	}

	if *findFlag != "" {
		findPasswords(kek, *findFlag)
	}
}

func savePassword(kek []byte, name string) {
	fmt.Println("KEK:", kek, "Save:", name)
}

func findPasswords(kek []byte, nameSubstring string) {
	fmt.Println("KEK", kek, "Find:", nameSubstring)
}
