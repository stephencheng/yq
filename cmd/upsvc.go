package cmd

import (
	"gopkg.in/op/go-logging.v1"
	"os"
)

func setLogLevel(level string) {
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
	)
	var backend = logging.AddModuleLevel(
		logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))

	//CRITICAL
	//ERROR v
	//WARNING vv
	//NOTICE vvv
	//INFO vvvv
	//DEBUG vvvvv

	//if verbose {
	//	backend.SetLevel(logging.DEBUG, "")
	//} else {
	//	backend.SetLevel(logging.ERROR, "")
	//}

	logLevel := logging.CRITICAL

	switch level {
	case "v":
		logLevel = logging.ERROR
	case "vv":
		logLevel = logging.WARNING
	case "vvv":
		logLevel = logging.NOTICE
	case "vvvv":
		logLevel = logging.INFO
	case "vvvvv":
		logLevel = logging.DEBUG
	}

	backend.SetLevel(logLevel, "")
	logging.SetBackend(backend)

}

func UpReadYml(logLevel string) error {
	var path = "tasks.**.task"

	setLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	var ymlstr = `
notes:
 goal:
   - to test using a simpler mode of data structure for tasks using array
 why:
   - array will allow you to put more attribute so that you can put desc

tasks:
 -
   name: task
   desc: this is task
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: shell
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist
 -
   name: task
   desc: this is task
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: shell
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist

`

	matchingNodes, errorReadingStream := cmReadYml(ymlstr, path, updateAll, docIndexInt)
	if errorReadingStream != nil {
		return errorReadingStream
	}
	var buf *YmlResult
	err := printResults(matchingNodes, buf)

	log.Infof("read result: \n%s", readResult)
	return err
}

var (
	readResult string
)

type YmlResult []byte

func (yml *YmlResult) Write(data []byte) (n int, err error) {
	ymlobj := YmlResult(data)
	yml = &ymlobj
	//p("write:", "[\n", string(data), "\n]")
	readResult = string(data)
	return len(data), nil
}
