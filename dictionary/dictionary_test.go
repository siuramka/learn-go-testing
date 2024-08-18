package dictionary

import "testing"

func TestSearch_FoundWord(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")

	want := Description("this is just a test")
	AssertSearch(t, got, want)
}

func TestSearch_NotFoundWord(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	_, err := dictionary.Search("nonExistingWord")

	AssertError(t, err, wordNotFound)
}

func TestAdd_AddsEntry(t *testing.T) {
	dictionary := make(Dictionary)
	const word = Word("joe")
	const desc = Description("momma")

	dictionary.Add(word, desc)

	got, _ := dictionary.Search(word)
	AssertSearch(t, got, desc)
}

func AssertSearch(t testing.TB, got, want Description) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func AssertError(t testing.TB, err error, expectedMsg string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedMsg {
		t.Errorf("expected error %s, got %s", expectedMsg, err.Error())
	}
}
