package dictionary

import "errors"

type Word string
type Description string

type Dictionary map[Word]Description

const wordNotFound = "word not found"

func (d Dictionary) Search(word Word) (Description, error) {
	desc, ok := d[word]
	if !ok {
		return "", errors.New(wordNotFound)
	}
	return desc, nil
}

func (d Dictionary) Add(word Word, desc Description) {
	d[word] = desc
}
