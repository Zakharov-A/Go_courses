package main

import "fmt"

// var arr = [10]int{150, 155, 9, 6, 5}

// func main() {
// 	// arr = append(arr, 200)

// 	fmt.Printf(`Array:%v`, arr[2])

// }

func main() {
	var x_exam [5]float64
	x_exam[0] = 98
	x_exam[1] = 93
	x_exam[2] = 77
	x_exam[3] = 82
	x_exam[4] = 83

	var tot float64
	for i := 0; i < len(x_exam); i++ {
		tot += x_exam[i]
	}
	fmt.Println(tot / float64(len(x_exam)))

}
