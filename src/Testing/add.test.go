package Testing

import (
	"awesomeProject/src/Maths"
	"testing"
	)


func TestAdd(t *testing.T){
	number := Maths.AddTwo(3,5);
	if number != 15 {
		t.Error("Expected: %n Actual: %n",number,15)
	}
}