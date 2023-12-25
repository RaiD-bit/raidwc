package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	stat, _ := os.Stdin.Stat()
	var data []byte
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Standard input is available (piped input)
		//fmt.Println("Reading from standard input...")
		var e error
		data, e = io.ReadAll(os.Stdin)
		if e != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input:", e)
		}
	}

	isNumBytes := flag.Bool("c", false, "get byte count")
	isLineCount := flag.Bool("l", false, "get line count")
	isWordCount := flag.Bool("w", false, "get word count")
	isCharCount := flag.Bool("m", false, "get char count")
	args := os.Args[1:]
	flag.Parse()

	var filenames []string
	var flags []string
	defaultFlags := []string{"-l", "-w", "-c"}

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			filenames = append(filenames, arg)
		} else if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}
	if flags == nil || len(flags) == 0 {
		flags = defaultFlags
		*isNumBytes = true
		*isLineCount = true
		*isWordCount = true
	}
	// lets extract filename args
	if len(filenames) == 0 || filenames == nil && data != nil {
		var result []string
		result = processFlags(flags, isNumBytes, result, data, isLineCount, isWordCount, isCharCount)
		//result = append(result, filename)
		final := strings.Join(result, "   ")
		fmt.Println("    " + final)
	} else if filenames != nil || len(filenames) == 0 {
		for _, filename := range filenames {
			var result []string

			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fileContent, err := io.ReadAll(file)

			result = processFlags(flags, isNumBytes, result, fileContent, isLineCount, isWordCount, isCharCount)
			final := strings.Join(result, "   ")
			fmt.Println("    " + final + " " + filename)
		}
	}

}

func processFlags(flags []string, isNumBytes *bool, result []string, fileContent []byte, isLineCount *bool, isWordCount *bool, isCharCount *bool) []string {
	for _, flagName := range flags {
		if flagName == "-c" && *isNumBytes {
			result = append(
				result,
				strconv.Itoa(len(fileContent)),
			)
		}
		if flagName == "-l" && *isLineCount {
			result = append(
				result,
				strconv.Itoa(getLineCount(fileContent)),
			)
		}
		if flagName == "-w" && *isWordCount {
			result = append(
				result,
				strconv.Itoa(getWordCount(fileContent)),
			)
		}
		if flagName == "-m" && *isCharCount {
			result = append(
				result,
				strconv.Itoa(getCharCount(fileContent)),
			)
		}

	}
	return result
}

func getCharCount(content []byte) int {
	return utf8.RuneCount(content)
}

func getWordCount(content []byte) int {
	words := bytes.Fields(content)
	return len(words)
}

func getLineCount(content []byte) int {
	reader := bytes.NewReader(content)
	scanner := bufio.NewScanner(reader)
	lineCnt := 0
	for scanner.Scan() {
		lineCnt++
	}
	return lineCnt
}
