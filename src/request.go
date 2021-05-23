package main

type RequestConfiguration struct {
	Method  string
	BaseURL string
	Ep      string
	Body    []byte
	Headers map[string]string

	Hits int64
}

func GetRequestConfigurationFromArgs(arg Arguments) RequestConfiguration {
	var requestConfig RequestConfiguration

	requestConfig.Method = arg.method
	requestConfig.BaseURL = arg.base
	requestConfig.Ep = arg.ep
	requestConfig.Headers = arg.parsedHeader
	requestConfig.Hits = arg.hits

	if len(arg.bodyFile) > 0 && arg.body == "" {
		requestConfig.Body = ReadFile(arg.bodyFile)
	} else {
		requestConfig.Body = []byte(arg.body)
	}

	return requestConfig
}
