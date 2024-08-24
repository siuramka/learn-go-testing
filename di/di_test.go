package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCalculate_SaveToRepository(t *testing.T) {
	repo := NewSumRepo()
	repo.Add(10)
	repo.Add(10)
	repo.Add(10)

	got := Calculate(repo)
	want := 30

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
