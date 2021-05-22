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
	if values.base[len(values.base)-1] == '/' && values.ep == "" {
		values.base = values.base[0 : len(values.base)-1]
		values.ep = "/"
	}
	Hit("GET", values.base, values.ep, values.body)
}
