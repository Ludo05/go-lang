package Loops

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