package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("non-existing")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		someValue := "some-value"
		dictionary.Add("test", someValue)
		got, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertStrings(t, got, someValue)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "teszt"}
		err := dictionary.Add("test", "test")

		assertError(t, err, ErrKeyAlreadyAdded)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "proba"}
		dictionary.Update("test", "teszt")
		got, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertStrings(t, got, "teszt")
	})

	t.Run("non existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "teszt")
		assertError(t, err, ErrorKeyDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "teszt"}
	dictionary.Delete("test")
	_, err := dictionary.Search("test")

	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
}
