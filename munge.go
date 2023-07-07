package main

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var tmpWordList []string
var inputFileLines []string

func mungeInit(arg UserInput) error {
	level := arg.Level
	wordListPath := arg.Word
	input := arg.Input
	output := arg.Output
	rDeplicate := arg.D

	if level < 0 {
		level = 0
	}

	if level > 9 {
		level = 9
	}

	if wordListPath != "" && input == "" {
		err := readWordList(wordListPath)
		if err != nil {
			return err
		}

		if rDeplicate {
			inputFileLines = removeDuplication(inputFileLines)
		}

		for _, line := range inputFileLines {
			mungeword(line, level)
		}
	} else if wordListPath == "" && input != "" {
		mungeword(input, level)
	} else {
		return errors.New("input or wordlist required")
	}

	err := writeMunge(output)
	if err != nil {
		return err
	}

	return nil
}

func readWordList(path string) error {
	readFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		inputFileLines = append(inputFileLines, fileScanner.Text())
	}

	return nil
}

func writeMunge(fileName string) error {
	if fileName == "" {
		fileName = "munge-passwords.txt"
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, data := range tmpWordList {
		_, err := datawriter.WriteString(data + "\n")
		if err != nil {
			return err
		}
	}

	color.Greenp("Saved to ", fileName, " With ", len(tmpWordList), " of Lines", "\n")
	datawriter.Flush()

	return nil
}

func removeDuplication(arr []string) []string {
	map_var := map[string]bool{}
	result := []string{}
	for e := range arr {
		if !map_var[arr[e]] {
			map_var[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func munge(wrd string, level int) {
	caser := cases.Title(language.English)
	if level > 0 {
		tmpWordList = append(tmpWordList, wrd)
		tmpWordList = append(tmpWordList, strings.ToUpper(wrd))
		tmpWordList = append(tmpWordList, caser.String(wrd))
	}
	if level > 2 {
		temp := caser.String(wrd)
		tmpWordList = append(tmpWordList, strings.ToLower(temp))
	}
	if level > 4 {
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "e", "3"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "a", "4"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "o", "0"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "i", "!"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "i", "1"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "l", "1"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "a", "@"))
		tmpWordList = append(tmpWordList, strings.ReplaceAll(wrd, "s", "$"))
	}
	if level > 4 {
		temp := wrd
		temp = strings.ReplaceAll(temp, "e", "3")
		temp = strings.ReplaceAll(temp, "a", "4")
		temp = strings.ReplaceAll(temp, "o", "0")
		temp = strings.ReplaceAll(temp, "i", "1")
		temp = strings.ReplaceAll(temp, "s", "$")
		tmpWordList = append(tmpWordList, temp)
	}
	if level > 4 {
		temp := wrd
		temp = strings.ReplaceAll(temp, "e", "3")
		temp = strings.ReplaceAll(temp, "a", "@")
		temp = strings.ReplaceAll(temp, "o", "0")
		temp = strings.ReplaceAll(temp, "i", "1")
		temp = strings.ReplaceAll(temp, "s", "$")
		tmpWordList = append(tmpWordList, temp)
	}
	if level > 4 {
		temp := wrd
		temp = strings.ReplaceAll(temp, "e", "3")
		temp = strings.ReplaceAll(temp, "a", "4")
		temp = strings.ReplaceAll(temp, "o", "0")
		temp = strings.ReplaceAll(temp, "i", "!")
		temp = strings.ReplaceAll(temp, "s", "$")
		tmpWordList = append(tmpWordList, temp)
	}
	if level > 4 {
		temp := wrd
		temp = strings.ReplaceAll(temp, "e", "3")
		temp = strings.ReplaceAll(temp, "a", "4")
		temp = strings.ReplaceAll(temp, "o", "0")
		temp = strings.ReplaceAll(temp, "l", "1")
		temp = strings.ReplaceAll(temp, "s", "$")
		tmpWordList = append(tmpWordList, temp)
	}
}

func mungeword(wrd string, level int) {
	munge(wrd, level)
	if level > 4 {
		munge(wrd+"1", level)
		munge(wrd+"123456", level)
		munge(wrd+"12", level)
		munge(wrd+"2", level)
		munge(wrd+"123", level)
		munge(wrd+"!", level)
		munge(wrd+".", level)
	}
	if level > 6 {
		munge(wrd+"?", level)
		munge(wrd+"_", level)
		munge(wrd+"0", level)
		munge(wrd+"01", level)
		munge(wrd+"69", level)
		munge(wrd+"21", level)
		munge(wrd+"22", level)
		munge(wrd+"23", level)
		munge(wrd+"1234", level)
		munge(wrd+"8", level)
		munge(wrd+"9", level)
		munge(wrd+"10", level)
		munge(wrd+"11", level)
		munge(wrd+"13", level)
		munge(wrd+"3", level)
		munge(wrd+"4", level)
		munge(wrd+"5", level)
		munge(wrd+"6", level)
		munge(wrd+"7", level)
	}
	if level > 7 {
		munge(wrd+"00", level)
		munge(wrd+"07", level)
		munge(wrd+"08", level)
		munge(wrd+"09", level)
		munge(wrd+"14", level)
		munge(wrd+"15", level)
		munge(wrd+"16", level)
		munge(wrd+"17", level)
		munge(wrd+"18", level)
		munge(wrd+"19", level)
		munge(wrd+"24", level)
		munge(wrd+"77", level)
		munge(wrd+"88", level)
		munge(wrd+"99", level)
		munge(wrd+"12345", level)
		munge(wrd+"123456789", level)
	}
}
