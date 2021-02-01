package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	countMap := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			//handle error
		}

		for _, line := range strings.Split(string(data), "\n") {
			countMap[line]++

		}

	}

	for line, count := range countMap {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}

}
