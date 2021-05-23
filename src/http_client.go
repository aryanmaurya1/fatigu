package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func createRequest(method, url string, headers map[string]string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	return req
}

func Hit(method, url, body string, headers map[string]string) (Metric, []byte) {

	metric := Metric{}
	bodyReader := strings.NewReader(body)

	client := &http.Client{}

	// creating request object
	req := createRequest(method, url, headers, bodyReader)

	metric.Start = time.Now().UnixNano()
	resp, err := client.Do(req)
	metric.End = time.Now().UnixNano()

	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	metric.Elasped = metric.End - metric.Start
	metric.ElaspedInms = (metric.Elasped) / int64(time.Millisecond)

	fmt.Println("Time : ", metric.ElaspedInms, "ms")
	// fmt.Println(string(bodyBytes))
	return metric, bodyBytes
}
