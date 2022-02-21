package main

import "fmt"

// var arr = [10]int{150, 155, 9, 6, 5}

// func main() {
// 	// arr = append(arr, 200)

// 	fmt.Printf(`Array:%v`, arr[2])

// }

func main() {
	//Объявление массива
	// var x_exam [5]float64
	// x_exam[0] = 98
	// x_exam[1] = 93
	// x_exam[2] = 77
	// x_exam[3] = 82
	// x_exam[4] = 83
	//Короткое объявление массива
	// x_exam := [6]float64{98, 93, 77, 82, 83, 99}
	// x_exam := [6]float64{
	// 	98,
	// 	93,
	// 	77,
	// 	82,
	// 	83,
	// 	99,
	// }

	// var tot float64 = 0
	// for _, value := range x_exam {
	// 	tot += value
	// }
	// fmt.Println(tot / float64(len(x_exam)))

	// Срезы
	// var xlist []float64
	// x := make([]float64, 5, 10)
	// fmt.Println(xlist, x)

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	fmt.Println(x)

}
