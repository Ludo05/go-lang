package Maths

import "fmt"

func AddTwo(num1 int8, num2 int8)  int8 {
	return num1 * num2
}

//Functions and Constants need to be declared with a capital letter in order to be used in other files
func Swap(nameOne, nameTwo string) (string, string) {
	return nameTwo, nameOne
}

func OneForTwoExport(sum int) (x,y int) {
	privateFun()
	x = sum * 2
	y = sum * 5
	return
}

func Div(x, y float64) (float64, error) {
	if x == 0.0 {
		return 0.0, fmt.Errorf("cant divide by zero")
	}
	return x/y, nil
};

func privateFun() () {
	fmt.Println("Private method")
}