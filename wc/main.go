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
	totalCountFlag     bool
}

type Output struct {
	LineCount      int
	WordCount      int
	CharacterCount int
}

var hasErrorsBool bool

type TotalCountOutput struct {
	TotalLineCount      int
	TotalWordCount      int
	TotalCharacterCount int
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
		out.CharacterCount = len([]rune(scanner.Text())) + out.CharacterCount + 1

	}
	fmt.Println(out.LineCount)
	if err := scanner.Err(); err != nil {
		PrintErr(err)
	}

	return out

}

func FormatOutput(i InputFlags, o Output, name string) string {

	var parts []string
	if i.LineCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.LineCount))
	}
	if i.WordCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.WordCount))
	}
	if i.CharacterCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.CharacterCount))
	}
	return fmt.Sprintf("%s %s", strings.Join(parts, " "), name)

}

func FormatFinalOutput(i InputFlags, o TotalCountOutput) string {

	var parts []string
	if i.LineCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.TotalLineCount))
	}
	if i.WordCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.TotalWordCount))
	}
	if i.CharacterCountFlag {
		parts = append(parts, fmt.Sprintf("%8d", o.TotalCharacterCount))
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

	var to TotalCountOutput
	if len(flag.Args()) >= 2 {
		input.totalCountFlag = true

	} else {
		input.totalCountFlag = false
	}
	if len(fileNames) == 0 {

		out := countAll(os.Stdin)
		output := FormatOutput(input, out, "")
		fmt.Println(output)

	}
	for _, file := range fileNames {
		// open the file
		f, err := OpenFile(file)

		if err != nil {
			PrintErr(err)
			continue
		}
		defer f.Close()

		out := countAll(f)
		if input.totalCountFlag {
			to.TotalCharacterCount += out.CharacterCount
			to.TotalLineCount += out.LineCount
			to.TotalWordCount += out.WordCount

		}
		finalOutput := FormatOutput(input, out, file)
		fmt.Println(finalOutput)

	}
	if input.totalCountFlag {
		finalTotalOutput := FormatFinalOutput(input, to)
		finalOutput := fmt.Sprintf("%s total", finalTotalOutput)
		fmt.Println(finalOutput)
	}

	defer func() {
		if hasErrorsBool {
			os.Exit(1)
		}

	}()

}

func PrintErr(err error) {

	fmt.Fprintln(os.Stderr, err)
	hasErrorsBool = true
}
