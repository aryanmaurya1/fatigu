package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type sample struct {
// 	a int
// }
// func call (a sample) {
// 	a.a = 100000
// 	fmt.Println(a)
// }
func main() {
	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	ValidateArgs(values)

	fmt.Println(values)

	rand.Seed(time.Now().UnixNano())
	// var randomNuber int = rand.Int()

	// var a sample
	// fmt.Println(a)
	// call(a)
	// fmt.Println(a)

}
