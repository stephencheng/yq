package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"
)

var (
	mock_delete_yml = `
tom:
  sex: male
  age: 23
jason:
  sex: male
  age: 35
emily:
  sex: female
  age: 32
`

	testfile = "localtest_delete.yml"
	filepath = path.Join(".", testfile)
)

func TestUpDeletePathFromFile001(t *testing.T) {
	ioutil.WriteFile(filepath, []byte(mock_delete_yml), 0644)
	result, err := UpDeletePathFromFile(filepath, "jason.sex", false, "v")
	fmt.Println(err)
	fmt.Printf("result:\n--------\n%s\n--------\n", result)

	content, err := ioutil.ReadFile(filepath)
	fmt.Println(string(content))
}

func TestUpDeletePathFromFile002(t *testing.T) {
	//in place delete: directly delete from file
	writeInplace = true
	ioutil.WriteFile(filepath, []byte(mock_delete_yml), 0644)
	result, err := UpDeletePathFromFile(filepath, "jason.sex", false, "v")
	fmt.Println(err)
	fmt.Printf("result:\n--------\n%s\n--------\n", result)

	content, err := ioutil.ReadFile(filepath)
	fmt.Printf("read file result:\n--------\n%s\n--------\n", string(content))
}

func TestUpDeletePathFromFile003(t *testing.T) {
	writeInplace = true
	ioutil.WriteFile(filepath, []byte(mock_delete_yml), 0644)
	result, err := UpDeletePathFromFile(filepath, "tom", false, "v")
	fmt.Println(err)
	fmt.Printf("result:\n--------\n%s\n--------\n", result)

	content, err := ioutil.ReadFile(filepath)
	fmt.Printf("read file result:\n--------\n%s\n--------\n", string(content))
}

func TestUpDeletePathFromFile004(t *testing.T) {
	//loglevel to vvvv
	ioutil.WriteFile(filepath, []byte(mock_delete_yml), 0644)
	result, err := UpDeletePathFromFile(filepath, "tom", true, "vvvvv")
	fmt.Println(err)
	fmt.Printf("result:\n--------\n%s\n--------\n", result)

	content, err := ioutil.ReadFile(filepath)
	fmt.Printf("read file result:\n--------\n%s\n--------\n", string(content))
}

//if it is in place, the result will be empty
func TestUpDeletePathFromFile005(t *testing.T) {
	ioutil.WriteFile(filepath, []byte(mock_delete_yml), 0644)
	result, err := UpDeletePathFromFile(filepath, "jason.sex", true, "v")
	fmt.Println(err)
	fmt.Printf("result:\n--------\n%s\n--------\n", result)

	content, err := ioutil.ReadFile(filepath)
	fmt.Println(string(content))
}
