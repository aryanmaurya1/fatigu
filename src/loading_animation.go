package main

import (
	"fmt"
	"time"
)

func showSpinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("Hitting... \r%c", r)
			time.Sleep(delay)
		}
	}
}
