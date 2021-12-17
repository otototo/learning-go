package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMsg := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("Saying Hello to People", func(t *testing.T) {
		got := Hello("Paulina", "")
		want := "Hello, Paulina!"

		assertCorrectMsg(t, got, want)
	})
	t.Run("Saying 'Hello, World!' when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMsg(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMsg(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Jaque", "French")
		want := "Bonjour, Jaque!"

        assertCorrectMsg(t, got, want)
	})
}
