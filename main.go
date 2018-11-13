package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var dir = "./data/"

type Contact struct{ tel, name, sdName string }

func (c *Contact) toString() string {
	return fmt.Sprintf("%s\t%s\t%s", c.tel, c.name, c.sdName)
}

func (c *Contact) create(dir string) error {
	text := []byte(c.name + " " + c.sdName)
	fileName := dir + c.tel
	return ioutil.WriteFile(fileName, text, 0777)
}

func (c *Contact) find(dir string, tel string) error {
	c.tel = tel
	text, err := ioutil.ReadFile(dir + c.tel)
	if err != nil {
		return err
	}
	textAr := strings.Fields(string(text))
	c.name = textAr[0]
	c.sdName = textAr[1]
	return nil
}

func main() {
	add := flag.String("add", "", "Add new contact. Write tel number, first and last name in quotes with spaces between them")
	tel := flag.String("tel", "", "Find one of your contacts. Write 11 digits for correct search")
	flag.Parse()

	_, err := ioutil.ReadDir(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Ended with error", err)
			os.Exit(2)
		}
		if err = os.Mkdir("data", 0777); err != nil {
			fmt.Println("Ended with error", err)
			os.Exit(2)
		}
	}

	if *add != "" {
		addContact(*add)
	} else if *tel != "" {
		printContact(*tel)
	} else {
		printAll()
	}
}

func addContact(newContact string) {
	newContactArr := strings.Fields(newContact)
	c := Contact{newContactArr[0], newContactArr[1], newContactArr[2]}
	if err := c.create(dir); err != nil {
		fmt.Println("Ended with error while adding contact", err)
		os.Exit(2)
	}
	fmt.Println("New contact created")
}

func printContact(tel string) {
	if len(tel) != 11 {
		fmt.Println("Incorrect number format. It must contain 11 digits")
		os.Exit(2)
	}
	c := Contact{}
	if err := c.find(dir, tel); err != nil {
		fmt.Println("Ended with error while searching contact", err)
		os.Exit(2)
	}
	fmt.Println(c.toString())
}
func printAll() {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Ended with error while reading contacts list", err)
		os.Exit(2)
	}
	if len(files) == 0 {
		fmt.Println("Empty contact list")
		os.Exit(2)
	}
	c := Contact{}
	for _, file := range files {
		if err := c.find(dir, file.Name()); err != nil {
			fmt.Println("Ended with error while reading contact", err)
			os.Exit(2)
		}
		fmt.Println(c.toString())
	}
}
