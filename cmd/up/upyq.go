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
	//roottest()
	//readTest()
	validateTest()
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
	result, _ := cmd.UpReadYmlStr(cmd.TestYmlStr, path, "vvv")
	fmt.Println(result)
}

func validateTest() {
	err := cmd.UpValidateYmlFile("./test/uptestdata.yml", "vvv")
	fmt.Println(err)
}
