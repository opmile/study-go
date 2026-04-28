package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("case_A: saying hello to people", func (t *testing.T) {
		got := Hello("milena", "")
		want := "hello, milena"
		assertCorrectMessage(t, got, want)
	})
	t.Run("case_B: empty string name", func (t *testing.T) {
		got := Hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("case_C: greeting in Spanish", func (t *testing.T) {
		got := Hello("elodi", "spanish")
		want := "hola, elodi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("case_D: greeting in French", func (t *testing.T) {
		got := Hello("piere", "french")
		want := "bonjour, piere"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want { t.Errorf("got %q want %q", got, want) }
}