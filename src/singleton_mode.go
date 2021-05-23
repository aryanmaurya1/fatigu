package main

import "sync"

func singleshot(values Arguments) {
	holder = new(sync.WaitGroup)

	for i := 0; i < int(values.hits); i++ {
		// methods := []string{"GET", "POST"}
		// ep := []string{"/", "/user", "/user/" + fmt.Sprintf("%d", time.Now().UnixNano())}

		go func(holder *sync.WaitGroup, values Arguments) {
			// Hit(methods[rand.Int()%2], values.base, ep[rand.Int()%3], values.body)
			Hit(values.method, values.base+values.ep, values.body, values.parsedHeader)
			holder.Done()
		}(holder, values)
		holder.Add(1)

	}
	holder.Wait()
}
