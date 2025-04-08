package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func wc(input *bufio.Scanner) (int, int, int) {
	lines, words, chars := 0, 0, 0
	for input.Scan() {
		line := input.Text()
		lines++
		words += len(strings.Fields(line))
		chars += len(line) + 1
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
	return lines, words, chars
}

func main() {
	// Define flags for command-line arguments
	linesFlag := flag.Bool("l", false, "Count number of lines")
	wordsFlag := flag.Bool("w", false, "Count number of words")
	charsFlag := flag.Bool("c", false, "Count number of characters")
	allFlag := flag.Bool("a", false, "Count lines, words, and characters")
	flag.Parse()

	// Ensure at least one flag is provided
	if !*linesFlag && !*wordsFlag && !*charsFlag && !*allFlag {
		fmt.Println("Usage: ccwc [-l] [-w] [-c] [-a] <filename>")
		os.Exit(1)
	}

	// Ensure a file is provided
	if len(flag.Args()) < 1 {
		fmt.Println("Error: No file provided")
		os.Exit(1)
	}

	// Process each file
	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s not found\n", filename)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines, words, chars := wc(scanner)

		// Print results based on flags
		if *linesFlag {
			fmt.Printf("Lines: %d\n", lines)
		}
		if *wordsFlag {
			fmt.Printf("Words: %d\n", words)
		}
		if *charsFlag {
			fmt.Printf("Characters: %d\n", chars)
		}
		if *allFlag {
			fmt.Printf("Lines: %d, Words: %d, Characters: %d\n", lines, words, chars)
		}
	}
}
