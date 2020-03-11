package cmd

import (
	"github.com/spf13/cobra"
)

func CreateUpReadCmd() *cobra.Command {
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
		RunE: upReadProperty,
	}
	cmdRead.PersistentFlags().StringVarP(&docIndex, "doc", "d", "0", "process document index number (0 based, * for all documents)")
	cmdRead.PersistentFlags().StringVarP(&printMode, "printMode", "p", "v", "print mode (v (values, default), p (paths), pv (path and value pairs)")
	cmdRead.PersistentFlags().StringVarP(&defaultValue, "defaultValue", "D", "", "default value printed when there are no results")
	cmdRead.PersistentFlags().BoolVarP(&printLength, "length", "l", false, "print length of results")
	cmdRead.PersistentFlags().BoolVarP(&collectIntoArray, "collect", "c", false, "collect results into array")
	cmdRead.PersistentFlags().BoolVarP(&explodeAnchors, "explodeAnchors", "X", false, "explode anchors")
	return cmdRead
}

func upReadProperty(cmd *cobra.Command, args []string) error {
	var path = "tasks.**.task"

	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := upReadYml(TestYmlStr, path, updateAll, docIndexInt)
	if errorReadingStream != nil {
		return errorReadingStream
	}
	out := cmd.OutOrStdout()
	return printResults(matchingNodes, out)
}
