package _map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, "test", definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Excepted '%s' to be delete", word)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s given, %s", got, want, "test")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %s want %s", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find add word:", err)
	}

	if got != definition {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}
