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

	if len(str) > nameLen {
		fmt.Println("Names cannot be larger than", nameLen, "bytes")
		return
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
	//todo: find an image that does not already contain a password (for image in images, decrypt entry with kek and check magic)

	var e entry
	e.init(name)
	e.generate_password()

	fmt.Println(e) //todo: insert entry into image
}

func findPasswords(kek []byte, nameSubstring string, images []string) {
	//todo: for each image, decrypt entry, check magic, and check name contains nameSubstring
	//todo: print matching name/password pairs
}

type entry struct {
	magic    uint32
	name     [nameLen]byte
	password [20]byte
}

const magic = 0xDEADBEEF
const nameLen = 20

func (e *entry) init(name string) {
	if len(name) > len(e.name) {
		log.Fatal("Name too large")
	}

	e.magic = magic
	copy(e.name[:], name)
}

func (e *entry) generate_password() {
	copy(e.password[:], "password123") //todo: generate random password
}

func (e *entry) check_magic() {
	if e.magic != magic {
		log.Fatal("Invalid magic")
	}
}
