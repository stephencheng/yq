package cmd

import (
	"fmt"
	"testing"
)

func TestUpValidateYmlStr(t *testing.T) {
	err := UpValidateYmlStr(TestYmlStr, "vvv")
	fmt.Println(err)
}

func TestUpValidateYmlStr2(t *testing.T) {
	err := UpValidateYmlStr(TestYmlStr3, "vvv")
	fmt.Println(err)
}

func TestUpValidateYmlFile(t *testing.T) {
	err := UpValidateYmlFile("./test/uptestdata.yml", "vvv")
	fmt.Println(err)
}
