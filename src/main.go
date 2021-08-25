package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var outputBuffer *os.File = os.Stdout

	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	values = ValidateArgs(values)

	// Validating configs from user if 'y' flag is not set.
	if !values.y {
		fmt.Println(values)
		var userResp string
		fmt.Printf("Is above mentioned config is valid ? [Y|N] : ")
		fmt.Scanf("%s", &userResp)
		if (userResp != "y") && (userResp != "Y") {
			os.Exit(0)
		}
	}

	// Setting up  log file if a valid path is provided
	logFile := values.logFile
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		HandlerError(err)
		outputBuffer = file

	}

	fmt.Fprintln(outputBuffer, time.Now())
	_, err := fmt.Fprintln(outputBuffer, values)
	HandlerError(err)

	requstConfigData := GetRequestConfigurationFromArgs(values)

	go showSpinner(100 * time.Millisecond) // Show a loading spinner on the terminal

	// Processing based on the mode.
	if values.s {
		for i := values.hitStart; i <= values.hitStop; i = i + values.hitStep {
			requstConfigData.Hits = i
			metrics := singleshot(requstConfigData, values.l, outputBuffer)
			analysis := Analyze(metrics)
			_, err = fmt.Fprintln(outputBuffer, analysis)
			HandlerError(err)
		}
	}
	outputBuffer.Close()
}
