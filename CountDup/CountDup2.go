package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countMap := make(map[string]int)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		countLines(os.Stdin, countMap)
	} else {
		for _, fileName := range fileNames {
			file, err := os.Open(fileName)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, countMap)
			file.Close()
		}

	}

	for line, n := range countMap {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(file *os.File, countMap map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		countMap[input.Text()]++
	}
}
