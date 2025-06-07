package gomaps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a string"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a string"
		assertString(t, got, want)
	})
	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("adding new word", func(t *testing.T) {
		word := "test"
		definition := "this is a  test"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("word already exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a  test"}
		word := "test"
		definition := "this is a  test"
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update defination", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a  test"}
		word := "test"
		newDefinition := "this is a  test"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a  test"}
		word := "test1"
		newDefinition := "this is a  test"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete a known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a  test"}
		word := "test"
		err := dictionary.Delete(word)
		if err != nil {
			t.Fatal("error is not expected here", err)
		}
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)

	})

	t.Run("delete a unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a  test"}
		word := "test1"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesntExist)

	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't go one")
	}
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, definition)
}
