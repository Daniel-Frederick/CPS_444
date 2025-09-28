package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type counts struct {
	lines int
	words int
	chars int
}

func wc(r io.Reader) counts {
	var c counts
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		c.lines++
		c.words += len(strings.Fields(line))
		c.chars += utf8.RuneCountInString(line) + 1 // +1 for newline
	}
	return c
}

func printCounts(c counts, showL, showW, showM bool, label string) {
	if showL {
		fmt.Printf("%7d", c.lines)
	}
	if showW {
		fmt.Printf("%7d", c.words)
	}
	if showM {
		fmt.Printf("%7d", c.chars)
	}
	if label != "" {
		fmt.Printf(" %s", label)
	}
	fmt.Println()
}

// parsing takes command-line args and returns (files, showL, showW, showM)
func parsing(args []string) ([]string, bool, bool, bool) {
	showL, showW, showM := false, false, false

	// consume options first
	for len(args) > 0 && strings.HasPrefix(args[0], "-") {
		opt := args[0][1:]
		if opt == "" {
			break
		}
		for _, ch := range opt {
			switch ch {
			case 'l':
				showL = true
			case 'w':
				showW = true
			case 'm':
				showM = true
			default:
				fmt.Fprintf(os.Stderr, "wc: invalid option -- '%c'\n", ch)
				os.Exit(1)
			}
		}
		args = args[1:]
	}

	// whatever remains are files
	files := args

	// default: all enabled if none given
	if !showL && !showW && !showM {
		showL, showW, showM = true, true, true
	}

	return files, showL, showW, showM
}

func main() {
	args := os.Args[1:]
	files, showL, showW, showM := parsing(args)

	if len(files) == 0 {
		// stdin
		c := wc(os.Stdin)
		printCounts(c, showL, showW, showM, "")
		return
	}

	var total counts
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", fname, err)
			os.Exit(1)
		}
		c := wc(f)
		f.Close()
		printCounts(c, showL, showW, showM, fname)

		total.lines += c.lines
		total.words += c.words
		total.chars += c.chars
	}

	if len(files) > 1 {
		printCounts(total, showL, showW, showM, "total")
	}
}
