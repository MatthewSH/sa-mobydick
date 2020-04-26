package main

import (
	"bufio"
	"container/list"
	"log"
	"os"
	"regexp"
	"sort"
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

				if !ListContains(stopWordsList, word) {
					mobyWordList[word]++
				}
			}
		}
	}

	// Since maps don't work the same in Go...we need to do some custom stuff.
	// Did not realize when I started that Go handles map WAY differently than any other language
	// I've ever worked with.

	type kv struct {
		Key string
		Value int
	}

	var sortedWordMap []kv

	for key, value := range mobyWordList {
		sortedWordMap = append(sortedWordMap, kv{ key, value})
	}

	sort.Slice(sortedWordMap, func(first int, second int) bool {
		return sortedWordMap[first].Value > sortedWordMap[second].Value
	})
}

func ListContains(haystack *list.List, needle string) bool {
	for item := haystack.Front(); item != nil; item = item.Next() {
		if item.Value == needle {
			return true
		}
	}

	return false
}