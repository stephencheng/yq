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

func UpReadYml(ymlstr, path, logLevel string) (string, error) {
	setLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return "", errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := cmReadYml(ymlstr, path, updateAll, docIndexInt)
	if errorReadingStream != nil {
		return "", errorReadingStream
	}
	var buf *YmlResult
	err := printResults(matchingNodes, buf)

	log.Infof("read result: \n%s", readResult)
	return readResult, err
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
