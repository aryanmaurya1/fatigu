package main

import (
	"fmt"
	"sync"
)

func singleshot(v RequestConfiguration) []Metric {

	var holder *sync.WaitGroup = new(sync.WaitGroup)
	var metrics []Metric = make([]Metric, 0, v.Hits)
	fmt.Println(v)
	for i := int64(0); i < int64(v.Hits); i++ {

		go func(holder *sync.WaitGroup, v RequestConfiguration, id int64) {

			metric, _ := Hit(v.Method, v.BaseURL+v.Ep, string(v.Body), v.Headers)
			metric.RoutineId = id
			metrics = append(metrics, metric)

			holder.Done()

		}(holder, v, i)
		holder.Add(1)

	}
	holder.Wait()
	return metrics
}
