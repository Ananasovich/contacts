package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var c = Contact{"79153423434", "Vasya", "Tupichkin"}

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

var dirTest = "./test_data/"

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
