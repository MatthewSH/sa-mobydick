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
			stopWordsList.PushBack(strings.ToLower(strings.TrimSpace(stopWordScanner.Text())))
		}
	}

	if err := stopWordScanner.Err(); err != nil {
		log.Fatal(err)
	}

	mobyFile, err := os.Open("./data/mobydick.txt")
	mobyWordList := make(map[string]int)

	if err != nil {
		log.Fatal(err)
	}

	defer mobyFile.Close()

	mobyScanner := bufio.NewScanner(mobyFile)

	puncRegex := regexp.MustCompile(`[.,\/#!$%\^&*;:{}=_~()\"\[\]\?]`)

	for mobyScanner.Scan() {
		mobyLine := mobyScanner.Text()
		mobyLine = strings.TrimSpace(mobyLine)

		if len(mobyScanner.Text()) > 0 {
			mobyLine = strings.ReplaceAll(mobyLine, "’s", "")
			mobyLine = strings.ReplaceAll(mobyLine, "'s", "")
			mobyLine = puncRegex.ReplaceAllString(mobyLine, "")
			lineArray := strings.Fields(mobyLine)

			for _, word := range lineArray {
				word = strings.ToLower(word)

				mobyWordList[word]++
			}
		}
	}
}

func ListContains(haystack *list.List, needle string) bool {
	for item := haystack.Front(); item != nil; item = item.Next() {
		if item.Value == needle {
			return true
		}
	}

	return false
}