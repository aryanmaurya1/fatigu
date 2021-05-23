package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) []byte {
	var f, err = os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	scanner := bufio.NewScanner(f)

	var jsonData []byte
	for scanner.Scan() {
		jsonData = append(jsonData, scanner.Bytes()...)
	}
	return jsonData
}

func GetParsedValues(data []byte) map[string]interface{} {

	var parsedValue map[string]interface{}
	err := json.Unmarshal(data, &parsedValue)
	if err != nil {
		fmt.Println(err)
	}
	return parsedValue
}
