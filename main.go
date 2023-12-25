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
	for _, filename := range filenames {
		// read files
		//fileInfo, err := os.Stat(filename)

		var result []string

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fileContent, err := io.ReadAll(file)

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
		//result = append(result, filename)
		final := strings.Join(result, "   ")
		fmt.Println("    " + final + " " + filename)
	}
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
