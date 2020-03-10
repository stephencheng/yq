package cmd

import (
	"fmt"
	"testing"
)

var (
	ymlstr = `
notes:
 goal:
   - to test using a simpler mode of data structure for tasks using array
 why:
   - array will allow you to put more attribute so that you can put desc

tasks:
 -
   name: task1
   desc: this is task1
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: cmd
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist
 -
   name: task2
   desc: this is task2
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: cmd
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist

`
)

func TestUpReadYml01(t *testing.T) {
	var path = "tasks.**.task"
	result, _ := UpReadYml(ymlstr, path, "vvv")
	fmt.Println(result)
}

func TestUpReadYml02(t *testing.T) {
	var path = "tasks.[1]"
	result, _ := UpReadYml(ymlstr, path, "vvv")
	fmt.Println(result)
}

func TestUpReadYml03(t *testing.T) {
	var path = "tasks.[*]"
	result, _ := UpReadYml(ymlstr, path, "vvv")
	fmt.Println(result)
}

func TestUpReadYml04(t *testing.T) {
	var path = "tasks(name==task2)"
	result, _ := UpReadYml(ymlstr, path, "vvv")
	fmt.Println(result)
}

func TestUpReadYml05(t *testing.T) {
	var path = "tasks(name==task*)"
	UpReadYml(ymlstr, path, "vvvv")
}

func TestUpReadYml06(t *testing.T) {
	var path = "tasks(name==task*).desc"
	UpReadYml(ymlstr, path, "vvvv")
}

func TestUpReadYml07(t *testing.T) {
	var path = "tasks(name==task*).task(func==cmd)"
	UpReadYml(ymlstr, path, "vvvv")
}
