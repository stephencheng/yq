package main

import (
	"fmt"
	"github.com/stephencheng/yq/v3/cmd"
	"os"

	logging "gopkg.in/op/go-logging.v1"
)

var (
	p = fmt.Println
)

func main() {
	p("=======test yml file==========")
	//p(cmd.TestYmlStr)
	p("==============================")
	roottest()
	//readTest()
	//validateTest()
	//validateTest()
	p("==============================")
}

func roottest() {
	cmd := cmd.UpCmd()
	log := logging.MustGetLogger("yq")
	if err := cmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func readTest() {
	var path = "tasks.**.task"
	result, _ := cmd.UpReadYmlStr(cmd.TestYmlStr, path, "vvv", false)
	fmt.Println(result)
}

func validateTest() {
	//run it in ide
	err := cmd.UpValidateYmlFile("./test/uptestdata.yml", "vvv")
	//run it in command line
	//err := cmd.UpValidateYmlFile("../../test/uptestdata.yml", "vvv")
	fmt.Println(err)
}
