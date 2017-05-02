package main

import (
	"bufio"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

func ReadData(filename string) ([]*Sentence, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}

	sentences := make([]*Sentence, 0)
	for _, sent := range splitBySentence(string(data)) {
		s, err := makeSentence(sent)
		if err != nil {
			break
		}
		sentences = append(sentences, s)
	}
	return sentences, nil
}

func SaveModel(weight *[]float64, filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(file)
	enc.Encode(&weight)
	return nil
}

func LoadModel(filename string) (*[]float64, error) {
	var w []float64
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	decoder := gob.NewDecoder(file)
	decoder.Decode(&w)
	return &w, nil
}
