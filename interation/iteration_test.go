package interation

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("returns aaaaa", func(t *testing.T) {
		count := 10
		repeated := Repeat("a", count)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("returns panic when char is not char", func(t *testing.T) {
		defer func() {
			const expectedPanic = "char is not char"
			r := recover()
			if r != "char is not char" {
				t.Errorf("expected %q but got %q", expectedPanic, r)
			}
		}()
		const count = 10
		Repeat("abc", 10)
		t.Errorf("should have paniced ")
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
