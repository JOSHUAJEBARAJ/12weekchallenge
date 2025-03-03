package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type InputFlags struct {
	LineCountFlag      bool
	CharacterCountFlag bool
	WordCountFlag      bool
}

type Output struct {
	LineCount      int
	WordCount      int
	CharacterCount int
}

func OpenFile(fileName string) (*os.File, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return f, nil

}

func countAll(r io.Reader) Output {

	var out Output

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		out.LineCount = out.LineCount + 1
		out.WordCount = len(strings.Fields(scanner.Text())) + out.WordCount
		out.CharacterCount = len([]rune(scanner.Text())) + out.CharacterCount

	}

	return out

}

func FormatOutput(i InputFlags, o Output) string {

	var parts []string
	if i.LineCountFlag {
		parts = append(parts, fmt.Sprintf("%d", o.LineCount))
	}
	if i.WordCountFlag {
		parts = append(parts, fmt.Sprintf("%d", o.WordCount))
	}
	if i.CharacterCountFlag {
		parts = append(parts, fmt.Sprintf("%d", o.CharacterCount))
	}
	return strings.Join(parts, " ")

}
func main() {

	var lcountFlag, wcountFlag, ccountFlag bool

	flag.BoolVar(&lcountFlag, "l", false, "line count of the file")
	flag.BoolVar(&wcountFlag, "w", false, "word count of the file")
	flag.BoolVar(&ccountFlag, "c", false, "character count of the file")

	flag.Parse()

	var input InputFlags
	input.LineCountFlag = lcountFlag
	input.CharacterCountFlag = ccountFlag
	input.WordCountFlag = wcountFlag

	fileNames := flag.Args()

	if !lcountFlag && !wcountFlag && !ccountFlag {
		input.LineCountFlag = true
		input.CharacterCountFlag = true
		input.WordCountFlag = true
	}

	for _, file := range fileNames {
		// open the file
		f, err := OpenFile(file)
		defer f.Close()
		if err != nil {
			PrintErr(err)
		}

		out := countAll(f)
		finalOutput := FormatOutput(input, out)
		finalOutput = fmt.Sprintf("   %s %s", finalOutput, file)
		fmt.Println(finalOutput)

	}

}

func PrintErr(err error) {

	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
