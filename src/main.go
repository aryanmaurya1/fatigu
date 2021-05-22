package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var holder *sync.WaitGroup

func main() {

	rand.Seed(time.Now().UnixNano())

	// Arguments parsing and validation
	var values Arguments
	values = ParseArgs(values)
	fmt.Println(values)
	values = ValidateArgs(values)

	holder = new(sync.WaitGroup)

	for i := 0; i < int(values.hits); i++ {
		// methods := []string{"GET", "POST"}
		// ep := []string{"/", "/user", "/user/" + fmt.Sprintf("%d", time.Now().UnixNano())}

		go func(holder *sync.WaitGroup) {
			// Hit(methods[rand.Int()%2], values.base, ep[rand.Int()%3], values.body)
			Hit(values.method, values.base+values.ep, values.body, nil)
			holder.Done()
		}(holder)
		holder.Add(1)

	}
	holder.Wait()
}
