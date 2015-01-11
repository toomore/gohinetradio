package gohinetradio

import "testing"

func TestGetURL(*testing.T) {
	GetURL("232")
	GetURL("")
}

func TestPrintChannel(*testing.T) {
	PrintChannel()
}

func TestGetList(*testing.T) {
	GenList()
}
