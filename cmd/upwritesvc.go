package cmd

import (
	"bufio"
	errors "github.com/pkg/errors"
	"github.com/stephencheng/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v3"
	"io"
)

func UpWriteNodeFromFile(filepath, elepath, value string, inplace bool, logLevel string) (string, error) {
	writeInplace = inplace
	SetLogLevel(logLevel)
	var args []string = []string{filepath, elepath, value}
	var updateCommands, updateCommandsError = readUpdateCommands(args, 3, "Must provide <filename> <path_to_update> <value>")
	if updateCommandsError != nil {
		return "", updateCommandsError
	}

	var buf YmlResultWriter
	err := updateDoc(args[0], updateCommands, &buf)
	return buf.Result, err
}

func UpWriteNodeFromStrForSimpleValue(ymlstr, elepath, value string, logLevel string) (string, error) {
	SetLogLevel(logLevel)
	var args []string = []string{ymlstr, elepath, value}
	var updateCommands, updateCommandsError = upReadUpdateCommandsSimpleValue(elepath, value)
	if updateCommandsError != nil {
		return "", updateCommandsError
	}

	var buf YmlResultWriter
	err := upUpdateDoc(args[0], updateCommands, &buf)

	return buf.Result, err
}

func UpWriteNodeFromStrForComplexValueFromYmlStr(ymlstr, elepath, complexymlvalue string, logLevel string) (string, error) {
	SetLogLevel(logLevel)
	var updateCommands, updateCommandsError = upReadUpdateCommandsComplexValue(elepath, complexymlvalue)
	if updateCommandsError != nil {
		return "", updateCommandsError
	}

	var buf YmlResultWriter
	err := upUpdateDoc(ymlstr, updateCommands, &buf)

	return buf.Result, err
}

func UpWriteNodeFromStrForComplexValueFromYmlFile(ymlstr, elepath, complexymlvaluefilepath string, logLevel string) (string, error) {
	SetLogLevel(logLevel)
	var updateCommands, updateCommandsError = upReadUpdateCommandsComplexValueFromFile(elepath, complexymlvaluefilepath)
	if updateCommandsError != nil {
		return "", updateCommandsError
	}

	var buf YmlResultWriter
	err := upUpdateDoc(ymlstr, updateCommands, &buf)

	return buf.Result, err
}

func upReadUpdateCommandsComplexValue(path, complextymlvalue string) ([]yqlib.UpdateCommand, error) {
	var updateCommands []yqlib.UpdateCommand = make([]yqlib.UpdateCommand, 0)

	var value yaml.Node
	err := upReadData(complextymlvalue, 0, &value)
	if err != nil && err != io.EOF {
		return nil, err
	}
	updateCommands = make([]yqlib.UpdateCommand, 1)
	updateCommands[0] = yqlib.UpdateCommand{Command: "update", Path: path, Value: value.Content[0], Overwrite: true}

	return updateCommands, nil
}

func upReadUpdateCommandsComplexValueFromFile(path, complextymlvalueFile string) ([]yqlib.UpdateCommand, error) {
	var updateCommands []yqlib.UpdateCommand = make([]yqlib.UpdateCommand, 0)

	var value yaml.Node
	err := readData(complextymlvalueFile, 0, &value)
	if err != nil && err != io.EOF {
		return nil, err
	}
	updateCommands = make([]yqlib.UpdateCommand, 1)
	updateCommands[0] = yqlib.UpdateCommand{Command: "update", Path: path, Value: value.Content[0], Overwrite: true}

	return updateCommands, nil
}

func upReadUpdateCommandsSimpleValue(path, value string) ([]yqlib.UpdateCommand, error) {
	var updateCommands []yqlib.UpdateCommand = make([]yqlib.UpdateCommand, 0)
	updateCommands = make([]yqlib.UpdateCommand, 1)
	log.Debug("path %v", path)
	log.Debug("Value %v", value)
	updateCommands[0] = yqlib.UpdateCommand{Command: "update", Path: path, Value: valueParser.Parse(value, customTag), Overwrite: true}
	return updateCommands, nil
}

func upReadData(ymlstr string, indexToRead int, parsedData interface{}) error {
	return upReadStream(ymlstr, func(decoder *yaml.Decoder) error {
		for currentIndex := 0; currentIndex < indexToRead; currentIndex++ {
			errorSkipping := decoder.Decode(parsedData)
			if errorSkipping != nil {
				return errors.Wrapf(errorSkipping, "Error processing document at index %v, %v", currentIndex, errorSkipping)
			}
		}
		return decoder.Decode(parsedData)
	})
}

func upUpdateDoc(ymlstr string, updateCommands []yqlib.UpdateCommand, writer io.Writer) error {
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	var updateData = func(dataBucket *yaml.Node, currentIndex int) error {
		if updateAll || currentIndex == docIndexInt {
			log.Debugf("Updating doc %v", currentIndex)
			for _, updateCommand := range updateCommands {
				log.Debugf("Processing update to Path %v", updateCommand.Path)
				errorUpdating := lib.Update(dataBucket, updateCommand, autoCreateFlag)
				if errorUpdating != nil {
					return errorUpdating
				}
			}
		}
		return nil
	}
	return upReadAndUpdate(writer, ymlstr, updateData)
}

func upReadAndUpdate(stdOut io.Writer, ymlstr string, updateData updateDataFn) error {
	var destination io.Writer
	var destinationName string
	destination = stdOut
	destinationName = "Stdout"

	log.Debugf("Writing to %v from %v", destinationName, "ymlstr")

	bufferedWriter := bufio.NewWriter(destination)
	defer safelyFlush(bufferedWriter)

	var encoder yqlib.Encoder
	if outputToJSON {
		encoder = yqlib.NewJsonEncoder(bufferedWriter, prettyPrint, indent)
	} else {
		encoder = yqlib.NewYamlEncoder(bufferedWriter, indent, colorsEnabled)
	}

	return upReadStream(ymlstr, mapYamlDecoder(updateData, encoder))
}
