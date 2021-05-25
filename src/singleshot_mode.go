package main

import (
	"sync"
)

func singleshot(v RequestConfiguration) []Metric {

	var holder *sync.WaitGroup = new(sync.WaitGroup)
	var queue = make(chan Metric, 100)
	var metrics []Metric = make([]Metric, 0, v.Hits)
	// fmt.Println(v)
	for i := int64(0); i < int64(v.Hits); i++ {

		go func(holder *sync.WaitGroup, v RequestConfiguration, id int64, queue chan<- Metric) {

			metric, _ := Hit(v.Method, v.BaseURL+v.Ep, string(v.Body), v.Headers)
			metric.RoutineId = id
			// metrics = append(metrics, metric)
			queue <- metric

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
