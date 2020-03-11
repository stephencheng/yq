package cmd

func UpValidateYmlFile(filename string, logLevel string) error {
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	_, errorReadingStream := readYamlFile(filename, "", updateAll, docIndexInt)

	return errorReadingStream
}

func UpValidateYmlStr(ymlstr string, logLevel string) error {
	SetLogLevel(logLevel)
	var updateAll, docIndexInt, errorParsingDocIndex = parseDocumentIndex()
	if errorParsingDocIndex != nil {
		return errorParsingDocIndex
	}

	_, errorReadingStream := upReadYml(ymlstr, "", updateAll, docIndexInt)
	if errorReadingStream != nil {
		return errorReadingStream
	}

	return errorReadingStream
}
