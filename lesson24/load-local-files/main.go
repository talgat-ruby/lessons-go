package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
)

//go:embed files/username-password-recovery-code.csv
var recoveryCodes string

//go:embed files/*.csv
var folder embed.FS

func main() {
	showLocalFiles()
	showEmbeddedFiles()
}

func showLocalFiles() {
	// Open the file
	file, err := os.Open("./files/username.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the contents of the file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Print the contents
	fmt.Println(string(data))
	fmt.Println("-------------------------------")
}

func showEmbeddedFiles() {
	fmt.Println(recoveryCodes)
	fmt.Println("-------------------------------")

	usernameContent, err := folder.ReadFile("files/username.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usernameContent))
	fmt.Println("-------------------------------")
}
