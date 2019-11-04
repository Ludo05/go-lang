package Loops

import "fmt"

func PickLoop(i int) []int {
	if i < 6 {
		var arr []int
		for a := 0; a <= i; a++ {
			arr = append(arr, a )
		}
	return arr
	}
	return []int{0}
}

func ArraySum(arr [6]int) (sum int) {
	for i := 0; i < len(arr); i++ {

		sum = sum + arr[i]
	}
	return sum
}
//It returns a pointer to a integer and the integer that we are returning is the result (ref to the result)
func RangeLoop(values ...int) *int {
	fmt.Print(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}