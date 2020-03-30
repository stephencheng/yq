package cmd

func UpReadYmlFile(filepath, epath, logLevel string, toArray bool) (string, error) {
	collectIntoArray = toArray
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return "", errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := readYamlFile(filepath, epath, updateAll, docIndexInt)

	if errorReadingStream != nil {
		return "", errorReadingStream
	}
	var buf YmlResultWriter
	err := printResults(matchingNodes, &buf)

	log.Infof("read result: \n%s", buf.Result)
	return buf.Result, err
}

func UpReadYmlStr(ymlstr, epath, logLevel string, toArray bool) (string, error) {
	collectIntoArray = toArray
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return "", errorParsingDocIndex
	}

	matchingNodes, errorReadingStream := upReadYml(ymlstr, epath, updateAll, docIndexInt)
	if errorReadingStream != nil {
		return "", errorReadingStream
	}
	var buf YmlResultWriter
	err := printResults(matchingNodes, &buf)

	log.Infof("read result: \n%s", buf.Result)
	return buf.Result, err
}
