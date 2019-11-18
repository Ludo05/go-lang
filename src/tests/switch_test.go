package tests

import (
	"awesomeProject/src/SwitchStatement"
	"testing"
)

func TestSwitch(t *testing.T) {
	t.Run("Switch statement tests", func(t *testing.T) {
		var got string = SwitchStatement.SayHello("Joe","french")
		var want string = "Bonjour Joe"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Switch statement tests", func(t *testing.T) {
		var got string = SwitchStatement.SayHello("Joe","spanish")
		var want string = "Hola Joe"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Switch statement tests", func(t *testing.T) {
		var got string = SwitchStatement.SayHello("","")
		var want string = "Hello World"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Switch statement tests", func(t *testing.T) {
		var got string = SwitchStatement.SayHello("","french")
		var want string = "Bonjour World"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}