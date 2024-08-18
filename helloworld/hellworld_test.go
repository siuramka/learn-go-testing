package helloworld

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world", "english")
	want := "Hello world"

	assertCorrectMessage(t, got, want)
}

func TestHelloSubtest(t *testing.T) {
	t.Run("returns 'Hello World'", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("returns 'Hello World' when have empty name", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("returns 'Hola World' when have empty name and spanish language", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns 'Labas World' when have empty name and lithuanian language", func(t *testing.T) {
		got := Hello("", "lithuanian")
		want := "Labas World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
