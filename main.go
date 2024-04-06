package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	saveFlag := flag.String("save", "", "Generate and save a password for the provided name")
	findFlag := flag.String("find", "", "Find names matching the provided substring")

	flag.Parse()
	if flag.NFlag() != 1 {
		fmt.Println("Must be called with one and only one valid input argument")
		flag.PrintDefaults()
		return
	}

	var op func([]byte, string, []string)
	var str string
	if *saveFlag != "" {
		op = savePassword
		str = *saveFlag
	} else {
		op = findPasswords
		str = *findFlag
	}

	images := make([]string, 0)
	for _, ext := range [3]string{"png", "jpg", "jpeg"} {
		names, err := filepath.Glob("*." + ext)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, names...)
	}

	fmt.Print("Password: ")
	pwd, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}

	kek, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	//todo: hash 480 bit kek or switch to different KDF

	op(kek, str, images)
}

func savePassword(kek []byte, name string, images []string) {
	fmt.Println("KEK:", kek)
	fmt.Println("Save:", name)
	fmt.Println("Images:", images)
}

func findPasswords(kek []byte, nameSubstring string, images []string) {
	fmt.Println("KEK:", kek)
	fmt.Println("Find:", nameSubstring)
	fmt.Println("Images:", images)
}
