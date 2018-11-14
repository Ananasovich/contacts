package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var dirTest = "./test_data/"
var c = Contact{"79153423434", "Vasya", "Tupichkin"}

func CreateTestEnvironment() {
	if err := os.Mkdir(dirTest, 0777); err != nil {
		fmt.Println("Ended with error", err)
		os.Exit(2)
	}
}

func DeleteTestEnvironment() {
	if err := os.Remove(dirTest); err != nil {
		fmt.Println("Ended with error", err)
		os.Exit(2)
	}
}
func TestMain(m *testing.M) {
	CreateTestEnvironment()
	defer DeleteTestEnvironment()
}

func TestToString(t *testing.T) {
	v := c.toString()
	r := "79153423434\tVasya\tTupichkin"
	if v != r {
		t.Error(
			"For", c,
			"Expected", r,
			"Got", v,
		)
	}
}

func TestCreate(t *testing.T) {
	c.create(dirTest)
	r := c.name + " " + c.sdName
	v, err := ioutil.ReadFile(dirTest + c.tel)
	if err != nil {
		t.Error("Can't read file with", err)
	} else if string(v) != r {
		t.Error(
			"For", c,
			"Expected", r,
			"Got", v,
		)
	}
	os.Remove(dirTest + c.tel)
}
func TestFind(t *testing.T) {
	c.create(dirTest)
	c1 := Contact{}
	c1.find(dirTest, c.tel)
	if c != c1 {
		t.Error(
			"Expected", c,
			"Got", c1,
		)
	}
	os.Remove(dirTest + c.tel)
}
