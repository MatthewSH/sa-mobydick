package main

import (
	"bufio"
	"container/list"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	stopWordsFile, err := os.Open("./data/stop-words.txt")
	stopWordsList := list.New()

	if err != nil {
		log.Fatal(err)
	}

	defer stopWordsFile.Close()

	stopWordScanner := bufio.NewScanner(stopWordsFile)

	for stopWordScanner.Scan() {
		if len(stopWordScanner.Text()) > 0 && !strings.Contains(stopWordScanner.Text(), "#") {
			stopWordsList.PushBack(strings.TrimSpace(stopWordScanner.Text()))
		}
	}

	if err := stopWordScanner.Err(); err != nil {
		log.Fatal(err)
	}

	mobyFile, err := os.Open("./data/mobydick.txt")
	mobyWordList := list.New()

	if err != nil {
		log.Fatal(err)
	}

	defer mobyFile.Close()

	mobyScanner := bufio.NewScanner(mobyFile)

	puncRegex := regexp.MustCompile(`[.,\/#!$%\^&*;:{}=_~()\"\[\]\?]`)
	splitRegex := regexp.MustCompile(`[ \n\r]+`)

	for mobyScanner.Scan() {
		mobyLine := strings.ReplaceAll(mobyScanner.Text(), "’s", "")
		mobyLine = puncRegex.ReplaceAllString(mobyLine, "")
		lineArray := splitRegex.Split(mobyLine, -1)

		for _, word := range lineArray {
			mobyWordList.PushBack(word)
		}
	}
}