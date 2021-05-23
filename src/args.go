package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
)

type Arguments struct {
	s bool // Singleshot Mode
	b bool // Batch Mode

	base           string            // Base URL
	ep             string            // [Method Endpoind]
	method         string            // comma separated list of methods
	header         string            // json parsable string in form '{"key":"value"}'
	parsedHeader   map[string]string // Actual parsed headers.
	body           string            // body to pass
	bodyFile       string            // path to json file containing body
	configFilePath string            // path to config file
	logFile        string            // file to write logs

	hits int64 // number of concurrent hits to perform
}

func ParseArgs(values Arguments) Arguments {
	flag.BoolVar(&values.s, "s", false, "Runs in singleshot mode.")
	flag.BoolVar(&values.b, "b", false, "Runs in Batch mode. Path to config file must be given.")
	flag.StringVar(&values.base, "base", "", "Base URL of API.")
	flag.StringVar(&values.ep, "ep", "", "Endpoint to hit. [Base + Endpoint]")
	flag.StringVar(&values.method, "method", "GET", "Comma separated list of methods to use.")
	flag.StringVar(&values.body, "body", "", "Body to send in API call.")
	flag.StringVar(&values.bodyFile, "body-file", "", "Path to json file to use as body content.")
	flag.StringVar(&values.configFilePath, "config-file", "", "Path to config file. Required only in case of batch mode.")
	flag.StringVar(&values.logFile, "log-file", "", "File to write logs.")
	flag.StringVar(&values.header, "headers", "{}", "Headers for request in form of key-value pair. (Valid JSON)")

	flag.Int64Var(&values.hits, "hits", 1000, "Number of concurrent hits to perform.")

	flag.Parse()

	return values
}

func (a Arguments) String() string {
	var repr = strings.Builder{}
	executionMode := "singleshot"
	if a.b {
		executionMode = "Batch"
	}
	repr.Grow(1200)
	repr.WriteString(strings.Repeat("-", 76) + "\n")
	repr.WriteString(fmt.Sprintf("| %-20s | %-50s|\n", "Flag", "Value"))
	repr.WriteString(strings.Repeat("-", 76) + "\n")
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Execution Mode", executionMode))
	// repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Batch Mode", a.b))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Base URL", a.base))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Endpoint", a.ep))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Methods", a.method))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Headers", a.header))
	// repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Parsed Headers", a.parsedHeader))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Body", a.body))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Body File", a.bodyFile))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Config File", a.configFilePath))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Concurrent Hits", a.hits))
	repr.WriteString(fmt.Sprintf("| %-20s | %-50v|\n", "Log File", a.logFile))
	repr.WriteString(strings.Repeat("-", 76) + "\n")

	return repr.String()
}

func ValidateArgs(values Arguments) Arguments {

	// validating execution mode and path to config file.
	if !values.s && !values.b {
		log.Fatal("Please provide a mode for execution.")
	} else if values.s && values.b {
		log.Fatal("Please provide a single mode of execution.")
	} else if values.b && len(values.configFilePath) == 0 {
		log.Fatal("Please provide a config for batch mode execution.")
	}

	// checking for correct files and arguments
	if len(values.body) > 0 && len(values.bodyFile) > 0 {
		log.Fatal("Please provide only one source for body.")
	}

	if len(values.bodyFile) > 0 && len(values.configFilePath) > 0 {
		log.Fatal("Please provide either body-file or config-file")
	}

	// validating presence of args in Singleshot mode
	if values.configFilePath == "" {
		if values.method == "" {
			log.Fatal("Please provide proper methods.")
		} else if values.base == "" {
			log.Fatal("Please provide a base URL.")
		}
	}

	// Fixing some args value
	if values.base[len(values.base)-1] == '/' && values.ep == "" {
		values.base = values.base[0 : len(values.base)-1]
		values.ep = "/"
	}

	// Parsing header flag value (JSON) as map.
	err := json.Unmarshal([]byte(values.header), &values.parsedHeader)
	if err != nil {
		fmt.Println(err.Error())
		values.parsedHeader = map[string]string{}
	}
	return values
}
