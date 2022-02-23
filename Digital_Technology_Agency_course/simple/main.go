package main

import "fmt"

// ----variable declaration ----
var arr = [10]int{150, 155, 9, 6, 5, 7, 89, 99} //Объявление массива

func main() {
	arr[4] = 102
	fmt.Println(arr)
}

//----

//---- Array declaration option 1  ----
// func main() {
// 	var x_exam [6]float64
// 	x_exam[0] = 98
// 	x_exam[1] = 93
// 	x_exam[2] = 77
// 	x_exam[3] = 82
// 	x_exam[4] = 83
// 	x_exam[5] = 86

// 	var total float64 = 0
// 	for i := 0; i < 6; i++ {
// 		total += x_exam[i]
// 	}

// 	fmt.Println(total / 6)
// }
// ----

//---- Array declaration option 2 conversion  ----

// func main() {
// 	var x_exam [6]float64
// 	x_exam[0] = 98
// 	x_exam[1] = 93
// 	x_exam[2] = 77
// 	x_exam[3] = 82
// 	x_exam[4] = 83
// 	x_exam[5] = 86

// 	var total float64 = 0
// 	for i := 0; i < len(x_exam); i++ {
// 		total += x_exam[i]
// 	}

// 	fmt.Println(total / float64(len(x_exam)))
// }
//----

//---- Array declaration option 2 conversion ----
// func main() {
// 	var x_el [6]float64
// 	x_el[0] = 98
// 	x_el[1] = 93
// 	x_el[2] = 77
// 	x_el[3] = 82
// 	x_el[4] = 83
// 	x_el[5] = 86

// 	var total float64 = 0
// 	for _, value := range x_el {
// 		total += value
// 	}

// 	fmt.Println(total / float64(len(x_el)))
// }
// ----

//---- Short array declaration ----
// func main() {
// x_el := [6]float64{98, 93, 77, 82, 83, 99} //option 1
// 	x_el := [6]float64{ //option 1
// 		98,
// 		93,
// 		77,
// 		82,
// 		83,
// 		99,
// 	}
// 	fmt.Println(x_el)
// }
// ----

// ---- Slices [low : high] ----
// func main() {
// 	x_el := []float64{98, 93, 77, 82, 83, 99}
// 	x_l := x_el[:]
// 	fmt.Println(x_l)
// }
// ----

// ---- Slice functions append ----
// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := append(slice1, 4, 5)
// 	fmt.Println(slice1, slice2)
// }
// ----

// ---- Slice functions copy ----
// func main() {
// 	slice1 := []int{5, 4, 3}
// 	slice2 := make([]int, 2)
// 	copy(slice2, slice1)
// 	fmt.Println(slice1, slice2)
// }
// ----

// ---- Map [string]int ----
// func main() {
// 	x := make(map[string]int)
// 	x["key"] = 10
// 	fmt.Println(x["key"])

// }

//----

// ---- Map [int]int ----
// func main() {
// 	elements := make(map[string]string)
// 	elements["H"] = "Hydrogen"
// 	elements["He"] = "Helium"
// 	elements["Li"] = "Lithium"
// 	elements["Be"] = "Beryllium"
// 	elements["B"] = "Boron"
// 	elements["C"] = "Carbon"
// 	elements["N"] = "Nitrogen"
// 	elements["O"] = "Oxygen"
// 	elements["F"] = "Fluorine"
// 	elements["Ne"] = "Neon"

// 	fmt.Println(elements["Li"])
// }

// ----

// ---- variable declaration and adding an element ----

// func main() {

// 	var arr []float64
// 	arr = append(arr, 200) //Добавление элемента массива

// 	fmt.Printf(`Array:%v`, arr[2]) //Вывод массива в консоль
// }
// ----
