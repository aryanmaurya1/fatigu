package main

import (
	"fmt"
	"os"
	"sync"
)

func singleshot(v RequestConfiguration, loggingMode bool, outputBuffer *os.File) []Metric {

	var holder *sync.WaitGroup = new(sync.WaitGroup)
	var queue = make(chan Metric, 100)
	var metrics []Metric = make([]Metric, 0, v.Hits)
	for i := int64(0); i < int64(v.Hits); i++ {

		go func(holder *sync.WaitGroup, v RequestConfiguration, id int64, queue chan<- Metric) {

			metric, body := Hit(v.Method, v.BaseURL+v.Ep, string(v.Body), v.Headers)
			metric.RoutineId = id
			queue <- metric

			if loggingMode {
				_, err := fmt.Fprintln(outputBuffer, "Time : ", metric.ElaspedInms, "ms")
				HandlerError(err)

				_, err = fmt.Fprintln(outputBuffer, string(body))
				HandlerError(err)

				_, err = fmt.Fprintf(outputBuffer, "\n\n")
				HandlerError(err)

			}
			holder.Done()

		}(holder, v, i, queue)
		holder.Add(1)

	}
	for i := int64(0); i < int64(v.Hits); i++ {
		metrics = append(metrics, <-queue)
	}
	holder.Wait()
	return metrics
}
