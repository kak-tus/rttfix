package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type service struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}

func main() {
	arg := []byte(os.Args[1])

	parsed := make([]service, 1000)

	err := json.Unmarshal(arg, &parsed)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("err: %s", err))
		os.Exit(1)
	}

	if len(parsed) == 0 {
		os.Exit(0)
	}

	filename := "/tmp/rttfix_" + parsed[0].Name

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		content = []byte("")
	}

	found := 0

	for i := 0; i < len(parsed); i++ {
		port := strconv.Itoa(parsed[i].Port)

		if parsed[i].Address+":"+port == string(content) {
			found = i
		}
	}

	if found > 0 {
		parsed = append([]service{parsed[found]}, parsed...)
	} else {
		port := strconv.Itoa(parsed[0].Port)
		content = []byte(parsed[0].Address + ":" + port)

		err := ioutil.WriteFile(filename, content, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("err: %s", err))
		}
	}

	result, err := json.Marshal(parsed)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("err: %s", err))
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, fmt.Sprintf("%s", result))
	os.Exit(0)
}
