package dictionary

import "errors"

type Dictionary map[string]string

type DictError string

func (e DictError) Error() string {
	return string(e)
}

const (
	ErrMsgWordNotFound = "word not found"
	ErrMsgWordExists   = "word already exists"
)

var (
	ErrWordNotFound = DictError(ErrMsgWordNotFound)
	ErrWordExists   = DictError(ErrMsgWordExists)
)

func (d Dictionary) Search(word string) (string, error) {
	desc, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return desc, nil
}

func (d Dictionary) Add(word string, desc string) error {
	_, searchErr := d.Search(word)

	switch {
	case errors.Is(searchErr, ErrWordNotFound):
		d[word] = desc
	case searchErr == nil:
		return ErrWordExists
	}

	return nil
}

func (d Dictionary) Update(word string, desc string) error {
	_, searchErr := d.Search(word)

	switch {
	case errors.Is(searchErr, ErrWordNotFound):
		return ErrWordNotFound
	case searchErr == nil:
		d[word] = desc
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, searchErr := d.Search(word)

	switch {
	case errors.Is(searchErr, ErrWordNotFound):
		return ErrWordNotFound
	case searchErr == nil:
		delete(d, word)
	}

	return nil
}
