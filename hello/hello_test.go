package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 告訴測試套件這個函式為輔助函式，當測試失敗時所報告的行數會在呼叫的函式當中，而不是輔助函式內部
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	// 子測試
	t.Run("say Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("William", "French")
		want := "Bonjour, William"
		assertCorrectMessage(t, got, want)
	})
}
