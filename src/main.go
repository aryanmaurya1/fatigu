package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var holder *sync.WaitGroup

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

	// Processing based on the mode.
	if values.s {
		singleshot(values)
	}

	if len(values.logFile) > 0 {
		// write logic for log writing
	}

}
