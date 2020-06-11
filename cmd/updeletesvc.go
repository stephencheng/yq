package cmd

import (
	"github.com/upcmd/yq/v3/pkg/yqlib"
)

func UpDeletePathFromFile(filepath, elepath string, inplace bool, logLevel string) (string, error) {
	writeInplace = inplace
	SetLogLevel(logLevel)
	var updateCommands []yqlib.UpdateCommand = make([]yqlib.UpdateCommand, 1)
	updateCommands[0] = yqlib.UpdateCommand{Command: "delete", Path: elepath}

	var buf YmlResultWriter
	err := updateDoc(filepath, updateCommands, &buf)
	return buf.Result, err
}
