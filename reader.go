package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Sentence struct {
	words []*Word
}

func makeSentence(s string) (*Sentence, error) {
	lines := strings.Split(s, "\n")
	if len(lines) < 4 {
		return nil, errors.New("Invalid line")
	}
	words := strings.Split(strings.TrimSpace(lines[0]), "\t")
	posTags := strings.Split(strings.TrimSpace(lines[1]), "\t")
	heads := strings.Split(strings.TrimSpace(lines[3]), "\t")

	sent := make([]*Word, 0)
	sent = append(sent, makeRootWord())
	for i := 0; i < len(words); i++ {
		head, err := strconv.ParseInt(heads[i], 10, 0)
		if err != nil {
			return nil, err
		}
		sent = append(sent, makeWord(words[i], posTags[i], i+1, int(head)))
	}
	return &Sentence{sent}, nil
}

func splitBySentence(s string) []string {
	return strings.Split(s, "\n\n")
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, sent := range splitBySentence(string(data)) {
		s, _ := makeSentence(sent)
		for _, sss := range s.words {
			fmt.Println(sss)
		}
	}
}
