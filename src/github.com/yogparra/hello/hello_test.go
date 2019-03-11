package main

import "testing"

/*
func TestHolaMundo(t *testing.T) {
	expected := "Hola mundo desde go."
	got := HolaMundo()

	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
*/

/* Se incorpora un nombre al saludo*/
/*
func TestHolaMundo(t *testing.T) {
	expected := "Hola, Guillermo"
	got := HolaMundo("Guillermo")

	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
*/

/* Se incorpora casos dintintos en un mismo test*/
/*
func TestHolaMundo(t *testing.T) {

	t.Run("01.- Sin nombre, respuesta Hola Mundo", func(t *testing.T) {
		expected := "Hola, Mundo"
		got := HolaMundo("")

		if expected != got {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})

	t.Run("02.- Con nombre, respuesta Hola nombre", func(t *testing.T) {
		expected := "Hola, Andres"
		got := HolaMundo("Andres")

		if expected != got {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})
}
*/

/* Se refactoriza el codigo */

/*
func TestHolaMundo(t *testing.T) {
	t.Run("01.- Sin nombre, respuesta Hola Mundo", func(t *testing.T) {
		expected := "Hola, Mundo"
		got := HolaMundo("")

		controError(expected, got, t)
	})

	t.Run("02.- Con nombre, respuesta Hola nombre", func(t *testing.T) {
		expected := "Hola, Andres"
		got := HolaMundo("Andres")

		controError(expected, got, t)
	})
}

func controError(expected string, got string, t *testing.T) {
	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
*/

/* refactorizando 02 */
func TestHolaMundo(t *testing.T) {

	controlError := func(expected string, got string, t *testing.T) {
		t.Helper()
		if expected != got {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	}

	t.Run("01.- Sin nombre, respuesta Hola Mundo", func(t *testing.T) {
		expected := "Hola, Mundo"
		got := HolaMundo("", "español")

		controlError(expected, got, t)
	})

	t.Run("02.- Con nombre, respuesta Hola nombre", func(t *testing.T) {
		expected := "Salut, Andres"
		got := HolaMundo("Andres", "franses")

		controlError(expected, got, t)
	})

}
