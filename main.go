package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	add := flag.String("add", "", "Use this option for add new contact. Write tel number, first and last name in quotes with spaces between them")
	tel := flag.String("tel", "", "Use this option fo find one of your contacts. Write 11 digits for correct search")
	flag.Parse()

	if *add != "" {
		addContact(*add)
	} else if *tel != "" {
		printContact(*tel)
	} else {
		printAll()
	}
}

type Contact struct{ tel, name, sdName string }

func addContact(add string) {
	addM := strings.Fields(add)
	c := Contact{addM[0], addM[1], addM[2]}
	text := []byte(c.name + " " + c.sdName)
	fileName := "./data/" + c.tel
	er := ioutil.WriteFile(fileName, text, 0777)
	if er != nil {
		fmt.Println("Ended with error while writing result to file", er)
		os.Exit(2)
	}
	fmt.Println("New contact created")
}
func printAll() {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		fmt.Println("Ended with error while reading contacts list", err)
		os.Exit(2)
	}
	if len(files) == 0 {
		fmt.Println("Empty contact list")
		os.Exit(2) 
	}
	for _, file := range files {
		fileName := "./data/" + file.Name()
		text, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Ended with error while reading contact", file.Name(), err)
			os.Exit(2)
		}
		textAr := strings.Fields(string(text))
		c := Contact{file.Name(), textAr[0], textAr[1]}
		fmt.Printf("%s\t%s\t%s", c.tel, c.name, c.sdName)
	}
}

func printContact(tel string) {
	if len(tel) != 11 {
		fmt.Println("Incorrect number format. It must contain 11 digits")
		os.Exit(2) 
	}
	files, _ := ioutil.ReadDir("./data")
	contEx := false
	for _, file := range files {
		if file.Name() == tel {
			contEx = true
		}
	}
	if !contEx {
		fmt.Println("You have no such contact")
		os.Exit(0)
	}
	fileName := "./data/" + tel
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ended with error while reading contact", err)
		os.Exit(2)
	}
	textAr := strings.Fields(string(text))
	c := Contact{tel, textAr[0], textAr[1]}
	fmt.Printf("%s\t%s\t%s", c.tel, c.name, c.sdName)

	fmt.Println()
}
