package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	values = ValidateArgs(values)
	fmt.Println(values)

	// Validating configs from user
	var userResp string
	fmt.Printf("Is above mentioned config is valid ? [Y|N] : ")
	fmt.Scanf("%s", &userResp)
	if (userResp != "y") && (userResp != "Y") {
		os.Exit(0)
	}
	requstConfigData := GetRequestConfigurationFromArgs(values)
	// Processing based on the mode.
	if values.s {
		singleshot(requstConfigData)
	}

	// Handle writing to log file
}
