package cmd

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	//working dir is proj/cmd
	//this init applies once for all in the same pkg
	os.Chdir("..")
}

func TestUpReadYml01(t *testing.T) {
	var path = "tasks.**.task"
	result, _ := UpReadYmlStr(TestYmlStr, path, "vvv", false)
	fmt.Println(result)
}

func TestUpReadYml02(t *testing.T) {
	var path = "tasks.[1]"
	result, _ := UpReadYmlStr(TestYmlStr, path, "vvv", false)
	fmt.Println(result)
}

func TestUpReadYml03(t *testing.T) {
	var path = "tasks.[*]"
	result, _ := UpReadYmlStr(TestYmlStr, path, "vvv", false)
	fmt.Println(result)
}

func TestUpReadYml04(t *testing.T) {
	var path = "tasks(name==task2)"
	result, _ := UpReadYmlStr(TestYmlStr, path, "vvv", false)
	fmt.Println(result)
}

func TestUpReadYml05(t *testing.T) {
	var path = "tasks(name==task*)"
	UpReadYmlStr(TestYmlStr, path, "vvvv", false)
}

func TestUpReadYml06(t *testing.T) {
	var path = "tasks(name==task*).desc"
	UpReadYmlStr(TestYmlStr, path, "vvvv", false)
}

func TestUpReadYml07(t *testing.T) {
	var path = "tasks(name==task*).task(func==cmd)"
	UpReadYmlStr(TestYmlStr, path, "vvvv", false)
}

//test read and collect into array
func TestUpReadYml08(t *testing.T) {
	var path = "tasks.**.desc"
	result, _ := UpReadYmlStr(TestYmlStr, path, "vvv", true)
	fmt.Println(result)
}

func TestUpReadYml09(t *testing.T) {
	UpReadYmlStr(TestYmlStr3, "nsw.sydney.[1].*(name==Emily)", "vvvv", false)
}

func TestUpReadYml10(t *testing.T) {
	UpReadYmlStr(TestYmlStr3, "nsw.sydney.**(name==Emily)", "vvvv", false)
}

func TestUpReadYml11(t *testing.T) {
	UpReadYmlStr(TestYmlStr3, "nsw.sydney.[*].kings(name==Emily)", "vvvv", false)
}

func TestUpReadYml12(t *testing.T) {
	UpReadYmlStr(TestYmlStr3, "nsw.sydney.[*].*(name==Grace)", "vvvv", false)
}

func TestUpReadYmlFile01(t *testing.T) {
	dir, _ := os.Getwd()
	p("pwd:", dir)
	var path = "tasks.**.task"
	result, _ := UpReadYmlFile("./test/uptestdata.yml", path, "vvvv", false)
	fmt.Println(result)
}
func TestUpReadYmlFile02(t *testing.T) {
	dir, _ := os.Getwd()
	p("pwd:", dir)
	var path = "tasks.**.task"
	result, _ := UpReadYmlFile("./test/uptestdata.yml", path, "vvvv", true)
	fmt.Println(result)
}
