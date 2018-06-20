package helloworld

import (
	"bufio"
	"fmt"
	"os"
)

var Test = 0

// PrintDuplineInfo print duplicate line info
func PrintDuplineInfo() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Printf("runing:%s\n", line)
		counts[line]++
	}
	if err := input.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
	fmt.Println("Over")
}

// PrintDuplineInfoOfFile file
func PrintDuplineInfoOfFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
