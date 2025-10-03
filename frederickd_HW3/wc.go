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

// wc reads from r and returns line, word, and character counts
func wc(r io.Reader) (counts, error) {
	var c counts
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		c.lines++
		c.words += len(strings.Fields(line))
		c.chars += utf8.RuneCountInString(line) + 1
	}

	if err := scanner.Err(); err != nil {
		return c, err
	}

	return c, nil
}

// printCounts prints the counts in wc format with proper spacing
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

// parseArgs parses command-line arguments and returns files and option flags
func parseArgs(args []string) ([]string, bool, bool, bool, error) {
	showL, showW, showM := false, false, false
	var files []string

	i := 0
	for i < len(args) && strings.HasPrefix(args[i], "-") && args[i] != "-" {
		opt := args[i][1:]
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
				fmt.Fprintf(os.Stderr, "Try 'wc --help' for more information.\n")
				return nil, false, false, false, fmt.Errorf("invalid option")
			}
		}
		i++
	}

	// Remaining arguments are files
	files = args[i:]

	// If no options specified, enable all
	if !showL && !showW && !showM {
		showL, showW, showM = true, true, true
	}

	return files, showL, showW, showM, nil
}

// printError prints error messages in wc format
func printError(fname string, err error) {
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "wc: %s: No such file or directory\n", fname)
	} else if os.IsPermission(err) {
		fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", fname)
	} else {
		errMsg := err.Error()
		if idx := strings.LastIndex(errMsg, ": "); idx != -1 {
			errMsg = errMsg[idx+2:]
		}
		fmt.Fprintf(os.Stderr, "wc: %s: %s\n", fname, errMsg)
	}
}

// handleStdin processes stdin input and prints counts
func handleStdin(showL, showW, showM bool, label string) error {
	c, err := wc(os.Stdin)
	if err != nil {
		return err
	}
	printCounts(c, showL, showW, showM, label)
	return nil
}

// processFile processes a single file (or stdin if fname == "-")
func processFile(fname string, showL, showW, showM bool) (counts, error) {
	if fname == "-" {
		c, err := wc(os.Stdin)
		if err != nil {
			return c, err
		}
		printCounts(c, showL, showW, showM, "-")
		return c, nil
	}

	f, err := os.Open(fname)
	if err != nil {
		return counts{}, err
	}
	defer f.Close()

	c, err := wc(f)
	if err != nil {
		return c, err
	}
	printCounts(c, showL, showW, showM, fname)
	return c, nil
}

// processFiles processes multiple files and returns totals, errors, and success count
func processFiles(files []string, showL, showW, showM bool) (counts, bool, int) {
	var total counts
	hadError := false
	successCount := 0

	for _, fname := range files {
		c, err := processFile(fname, showL, showW, showM)
		if err != nil {
			printError(fname, err)
			hadError = true
			continue
		}
		total.lines += c.lines
		total.words += c.words
		total.chars += c.chars
		successCount++
	}

	return total, hadError, successCount
}

func main() {
	args := os.Args[1:]

	files, showL, showW, showM, err := parseArgs(args)
	if err != nil {
		os.Exit(1)
	}

	// If no files given → stdin
	if len(files) == 0 {
		if err := handleStdin(showL, showW, showM, ""); err != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// If only "-" → stdin
	if len(files) == 1 && files[0] == "-" {
		if err := handleStdin(showL, showW, showM, ""); err != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Process file list
	total, hadError, successCount := processFiles(files, showL, showW, showM)

	// Print total if multiple files
	if successCount > 1 {
		printCounts(total, showL, showW, showM, "total")
	}

	if hadError {
		os.Exit(1)
	}
}

