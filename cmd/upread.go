package cmd

import (
	"github.com/stephencheng/yq/v3/pkg/yqlib"
	yaml "gopkg.in/yaml.v3"
	"io"
	"strings"
)

func upReadYml(ymlstr string, path string, updateAll bool, docIndexInt int) ([]*yqlib.NodeContext, error) {
	return upDoReadYamlStr(ymlstr, createReadFunction(path), updateAll, docIndexInt)
}

func upDoReadYamlStr(ymlstr string, readFn readDataFn, updateAll bool, docIndexInt int) ([]*yqlib.NodeContext, error) {
	var matchingNodes []*yqlib.NodeContext

	var currentIndex = 0
	var errorReadingStream = upReadStream(ymlstr, func(decoder *yaml.Decoder) error {
		for {
			var dataBucket yaml.Node
			errorReading := decoder.Decode(&dataBucket)

			if errorReading == io.EOF {
				return handleEOF(updateAll, docIndexInt, currentIndex)
			} else if errorReading != nil {
				return errorReading
			}

			var errorParsing error
			matchingNodes, errorParsing = appendDocument(matchingNodes, dataBucket, readFn, updateAll, docIndexInt, currentIndex)
			if errorParsing != nil {
				return errorParsing
			}
			if !updateAll && currentIndex == docIndexInt {
				log.Debug("all done")
				return nil
			}
			currentIndex = currentIndex + 1
		}
	})
	return matchingNodes, errorReadingStream
}

func upReadStream(ymlstr string, yamlDecoder yamlDecoderFn) error {
	var stream io.Reader
	stream = strings.NewReader(ymlstr)

	return yamlDecoder(yaml.NewDecoder(stream))
}
