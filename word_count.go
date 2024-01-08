package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wordcount := 0

	for scanner.Scan() {
		wordcount++
	}
	return wordcount
}

func main() {
	lines := flag.Bool("l", false, "count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}
