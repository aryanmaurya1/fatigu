package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	values = ValidateArgs(values)
	fmt.Println(values)

	if !values.y {
		// Validating configs from user
		var userResp string
		fmt.Printf("Is above mentioned config is valid ? [Y|N] : ")
		fmt.Scanf("%s", &userResp)
		if (userResp != "y") && (userResp != "Y") {
			os.Exit(0)
		}
	}

	requstConfigData := GetRequestConfigurationFromArgs(values)

	var result strings.Builder = strings.Builder{} // Final logging string
	result.Grow(1200)

	// Processing based on the mode.
	if values.s {
		for i := values.hitStart; i <= values.hitStop; i = i + values.hitStep {
			requstConfigData.Hits = i
			metrics := singleshot(requstConfigData)
			analysis := Analyze(metrics)
			result.WriteString(analysis)
			fmt.Println(analysis)
			result.WriteString("\n\n")
		}
	}

	// Handle writing to log file
	// fmt.Println(result.String())
}
