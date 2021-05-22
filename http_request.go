package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Hit(method, base, ep, body string) (int64, string) {

	bodyReader := strings.NewReader(body)
	client := &http.Client{}

	// Add url params in ep
	fmt.Println(method, base+ep, body)
	req, err := http.NewRequest(method, base+ep, bodyReader)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0, string(bodyBytes)
}
