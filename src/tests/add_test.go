package tests

import (
	"awesomeProject/src/Maths"
	"testing"
)

func TestAdd(t *testing.T){
	t.Run("Should * two numbers together", func(t *testing.T) {
		got := int(Maths.AddTwo(4,3))
		want := 12

		if got != want {
			t.Errorf("Got %v want %v", got,want)
		}
	})

	t.Run("Should swap letters around", func(t *testing.T) {
		one, two := Maths.Swap("Hello","World")
		got := [2]string{one,two}
		want := [2]string{"World","Hello"}

		if got != want {
			t.Errorf("Got %v want %v", got,want)
		}
	})
}
