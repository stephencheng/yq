package cmd

import (
	"fmt"
	"testing"
)

func TestWrite01(t *testing.T) {
	result, err := UpWriteNodeFromFile("../test/write01.yml", "jason.sex", "female", false, "v")
	fmt.Println("------------------")
	fmt.Println(result)
	fmt.Println("------------------")
	fmt.Println(err)
}

func TestWrite02(t *testing.T) {
	result, err := UpWriteNodeFromFile("../test/write01.yml", "jason.sex", "female", true, "vvvvv")
	fmt.Println("------------------")
	fmt.Println(result)
	fmt.Println("------------------")
	fmt.Println(err)
}

var (
	ymlstr = `
jason:
  sex: female
  age: 30
`

	wife = `
emily:
  sex: female
  age: 24
`
)

func TestWrite03(t *testing.T) {
	result, err := UpWriteNodeFromStrForSimpleValue(ymlstr, "jason.sex", "male", "v")
	fmt.Println("------------------")
	fmt.Println(result)
	fmt.Println("------------------")
	fmt.Println(err)
}

//write the node using complex sub node using string
func TestWrite04(t *testing.T) {
	result, err := UpWriteNodeFromStrForComplexValueFromYmlStr(ymlstr, "jason.wife", wife, "v")
	fmt.Println("------------------")
	fmt.Println(result)
	fmt.Println("------------------")
	fmt.Println(err)
}

//write the node using complex sub node using file
func TestWrite05(t *testing.T) {
	result, err := UpWriteNodeFromStrForComplexValueFromYmlFile(ymlstr, "jason.wife", "../test/write02.yml", "v")
	fmt.Println("------------------")
	fmt.Println(result)
	fmt.Println("------------------")
	fmt.Println(err)
}
