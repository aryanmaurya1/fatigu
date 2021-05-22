package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	fmt.Println(values)
	ValidateArgs(values)

	ReadConfigFile(values.configFilePath)
}
