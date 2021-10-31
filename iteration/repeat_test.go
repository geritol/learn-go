package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("*", 5)
	fmt.Println(repeated)
	// Output: *****
}

func TestRepeat(t *testing.T) {
	t.Run("repeats string 5 times", func(t *testing.T) {

		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeats string 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
