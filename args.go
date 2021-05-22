package main

import (
	"flag"
	"fmt"
	"strings"
)

type Arguments struct {
	s              bool   // Singleton Mode
	b              bool   // Batch Mode
	base           string // Base URL
	ep             string // [Method Endpoind]
	method         string // comma separated list of methods
	body           string // body to pass
	bodyFile       string // path to json file containing body
	hits           int64  // number of concurrent hits to perform
	configFilePath string // path to config file
}

func ParseArgs(values Arguments) Arguments {
	flag.BoolVar(&values.s, "s", false, "Runs fatigu in singleton mode.")
	flag.BoolVar(&values.b, "b", false, "Runs fatigu in Batch mode. Path to config file must be given.")
	flag.StringVar(&values.base, "base", "", "Base URL of API.")
	flag.StringVar(&values.ep, "ep", "", "Endpoint to hit. [Base + Endpoint]")
	flag.StringVar(&values.method, "method", "GET", "Comma separated list of methods to use.")
	flag.StringVar(&values.body, "body", "", "Body to send in API call.")
	flag.StringVar(&values.bodyFile, "bodyfile", "", "Path to json file to use as body content.")
	flag.StringVar(&values.configFilePath, "file", "", "Path to config file. Required only in case of batch mode.")
	flag.Int64Var(&values.hits, "hits", 1000, "Number of concurrent hits to perform.")
	flag.Parse()
	return values
}

func (a Arguments) String() string {
	var repr = ""
	repr = repr + strings.Repeat("-", 56) + "\n"
	repr = repr + fmt.Sprintf("| %-20s | %-30s|\n", "Flag", "Value")
	repr = repr + strings.Repeat("-", 56) + "\n"
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Singleton Mode", a.s)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Batch Mode", a.b)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Base URL", a.base)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Endpoint", a.ep)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Methods", a.method)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Body", a.body)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Body File", a.bodyFile)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Config File", a.configFilePath)
	repr = repr + fmt.Sprintf("| %-20s | %-30v|\n", "Concurrent Hits", a.hits)
	repr = repr + strings.Repeat("-", 56) + "\n"

	return repr
}

func ValidateArgs(values Arguments) {
	if !values.s && !values.b {
		panic("Please provide a mode for execution.")
	} else if values.s && values.b {
		panic("Please provide a single mode of execution.")
	} else if len(values.body) > 0 && len(values.bodyFile) > 0 {
		panic("Please provide only one arg for body.")
	} else if values.b && len(values.configFilePath) == 0 {
		panic("Please provide a config for batch mode execution.")
	} else if values.method == "" {
		panic("Please provide proper methods.")
	} else {
	}
}
