package main

import (
	"fmt"
	"strings"
	"time"
)

func MinElaspedTime(metrics []Metric) Metric {
	var min Metric
	for i := 0; i < len(metrics); i++ {
		if i == 0 {
			min = metrics[0]
		}
		if metrics[i].Elasped < min.Elasped {
			min = metrics[i]
		}
	}
	return min
}

func MaxElaspedTime(metrics []Metric) Metric {
	var max Metric
	for i := 0; i < len(metrics); i++ {
		if i == 0 {
			max = metrics[0]
		}
		if metrics[i].Elasped > max.Elasped {
			max = metrics[i]
		}
	}
	return max
}

func AverageElaspedTime(metrics []Metric) float32 {
	var avg float32
	var metricsCount = len(metrics)

	for i := 0; i < metricsCount; i++ {
		avg = avg + float32(metrics[i].Elasped)
	}
	avg = avg / float32(metricsCount)
	return avg
}

func Analyze(metrics []Metric) string {
	var result strings.Builder = strings.Builder{}
	timeInms := AverageElaspedTime(metrics) / float32(time.Millisecond)
	result.Grow(1000)

	result.WriteString(fmt.Sprintf("Concurrency : %d \n", len(metrics)))
	result.WriteString(fmt.Sprintln(MaxElaspedTime(metrics)))
	result.WriteString(fmt.Sprintln(MinElaspedTime(metrics)))
	result.WriteString(fmt.Sprintf("Average Time : %.2f ms \n", timeInms))

	return result.String()
}
