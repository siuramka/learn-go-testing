package dictionary

import "testing"

func TestSearch_FoundWord(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")

	want := "this is just a test"
	assertEquals(t, got, want)
}

func TestSearch_NotFoundWord(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	_, err := dictionary.Search("nonExistingWord")

	assertError(t, err, ErrMsgWordNotFound)
}

func TestAdd_AddEntry(t *testing.T) {
	dictionary := make(Dictionary)
	word := "joe"
	desc := "momma"

	err := dictionary.Add(word, desc)

	got, _ := dictionary.Search(word)
	assertNoError(t, err)
	assertEquals(t, got, desc)
}

func TestAdd_DontAddExisingEntry(t *testing.T) {
	word := "joe"
	desc := "momma"
	dictionary := Dictionary{word: desc}

	err := dictionary.Add(word, desc)

	assertError(t, err, ErrMsgWordExists)
}

func TestUpdate_UpdatesExistingWord(t *testing.T) {
	word := "joe"
	desc := "momma"
	dictionary := Dictionary{word: desc}
	newDesc := "cool"

	err := dictionary.Update(word, newDesc)

	got, _ := dictionary.Search(word)
	assertNoError(t, err)
	assertEquals(t, got, newDesc)
}

func TestUpdate_ErrWhenUpdateNotExistingWord(t *testing.T) {
	dictionary := Dictionary{"joe": "momma"}
	newDesc := "cool"
	notExistingWord := "newJoe"

	err := dictionary.Update(notExistingWord, newDesc)

	assertError(t, err, ErrMsgWordNotFound)
}

func TestDelete_ShouldDelete(t *testing.T) {
	word := "joe"
	desc := "momma"
	dictionary := Dictionary{word: desc}

	errDel := dictionary.Delete(word)

	assertNoError(t, errDel)
	_, errSearch := dictionary.Search(word)
	assertError(t, errSearch, ErrMsgWordNotFound)
}

func TestDelete_ErrWhenDeleteNonExisting(t *testing.T) {
	dictionary := Dictionary{"joe": "momma"}
	notExistingWord := "mio"

	err := dictionary.Delete(notExistingWord)

	assertError(t, err, ErrMsgWordNotFound)
}

func assertEquals(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t testing.TB, err error, expectedMsg string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedMsg {
		t.Errorf("expected error %s, got %s", expectedMsg, err.Error())
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
