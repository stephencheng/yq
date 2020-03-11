package cmd

func UpReadYmlFile(filepath, epath, logLevel string) (string, error) {
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return "", errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := readYamlFile(filepath, epath, updateAll, docIndexInt)

	if errorReadingStream != nil {
		return "", errorReadingStream
	}
	var buf *YmlResult
	err := printResults(matchingNodes, buf)

	log.Infof("read result: \n%s", readResult)
	return readResult, err
}

func UpReadYmlStr(ymlstr, epath, logLevel string) (string, error) {
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return "", errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := upReadYml(ymlstr, epath, updateAll, docIndexInt)
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
	readResult = string(data)
	return len(data), nil
}
