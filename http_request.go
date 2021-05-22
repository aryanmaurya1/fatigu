package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func T() {
	b := "{'a' : 'b'}"
	stringReader := strings.NewReader(b)
	// fmt.Println(n, err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.writups.tech/", stringReader)
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

	fmt.Println(string(bodyBytes))
}
