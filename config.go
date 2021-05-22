package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadConfigFile(path string) []byte {
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
