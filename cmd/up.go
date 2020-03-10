package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stephencheng/yq/v3/pkg/yqlib"
	yaml "gopkg.in/yaml.v3"
	"io"
	"strings"
)


var(
	p=fmt.Println
)

func CreateUpCmd() *cobra.Command {
	var cmdRead = &cobra.Command{
		Use:     "read [yaml_file] [path_expression]",
		Aliases: []string{"u"},
		Short:   "yq u [--printMode/-p pv] sample.yaml 'b.e(name==fr*).value'",
		Example: `
yq read things.yaml 'a.b.c'
yq u - 'a.b.c' # reads from stdin
yq u things.yaml 'a.*.c'
yq u things.yaml 'a.**.c' # deep splat
yq u things.yaml 'a.(child.subchild==co*).c'
yq u -d1 things.yaml 'a.array[0].blah'
yq u things.yaml 'a.array[*].blah'
yq u -- things.yaml '--key-starting-with-dashes.blah'
      `,
		Long: "Outputs the value of the given path in the yaml file to STDOUT",
		RunE: CmReadProperty,
	}
	cmdRead.PersistentFlags().StringVarP(&docIndex, "doc", "d", "0", "process document index number (0 based, * for all documents)")
	cmdRead.PersistentFlags().StringVarP(&printMode, "printMode", "p", "v", "print mode (v (values, default), p (paths), pv (path and value pairs)")
	cmdRead.PersistentFlags().StringVarP(&defaultValue, "defaultValue", "D", "", "default value printed when there are no results")
	cmdRead.PersistentFlags().BoolVarP(&printLength, "length", "l", false, "print length of results")
	cmdRead.PersistentFlags().BoolVarP(&collectIntoArray, "collect", "c", false, "collect results into array")
	cmdRead.PersistentFlags().BoolVarP(&explodeAnchors, "explodeAnchors", "X", false, "explode anchors")
	return cmdRead
}


func CmReadProperty(cmd *cobra.Command, args []string) error {
	var path = "tasks.**.task"

	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	var ymlstr=`
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
	out := cmd.OutOrStdout()
	return printResults(matchingNodes, out)
}

func cmReadYml(ymlstr string, path string, updateAll bool, docIndexInt int) ([]*yqlib.NodeContext, error) {
	return cmDoReadYamlStr(ymlstr, createReadFunction(path), updateAll, docIndexInt)
}

func cmDoReadYamlStr(ymlstr string, readFn readDataFn, updateAll bool, docIndexInt int) ([]*yqlib.NodeContext, error) {
	var matchingNodes []*yqlib.NodeContext

	var currentIndex = 0
	var errorReadingStream = cmReadStream(ymlstr, func(decoder *yaml.Decoder) error {
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

func cmReadStream(ymlstr string, yamlDecoder yamlDecoderFn) error {
	var stream io.Reader
	stream= strings.NewReader(ymlstr)

	return yamlDecoder(yaml.NewDecoder(stream))
}
