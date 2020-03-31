package cmd

import (
	"fmt"
	"gopkg.in/op/go-logging.v1"
	"os"
)

var (
	p = fmt.Println
)

func SetLogLevel(level string) {
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

type YmlResultWriter struct {
	Buf    *[]byte
	Result string
}

func (yml *YmlResultWriter) Write(data []byte) (n int, err error) {
	yml.Buf = &data
	yml.Result = string(data)
	return len(data), nil
}
