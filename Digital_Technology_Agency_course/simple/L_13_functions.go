package main

import "fmt"

// ---- function option ----

// func average(xs []float64) float64 {
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(xs))
// }

// ----

// ---- function option 2 ----

// func average(xs []float64) float64 {
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }

// func main() {
// 	someOtherName := []float64{198, 93, 77, 82, 83}
// 	fmt.Println(average(someOtherName))
// }

// ----

// ---- function option 3 ----

// func f(x int) {
// 	fmt.Println(x)
// }
// func main() {
// 	x := 5
// 	f(x)
// }

// ----

// ---- call stack ----

// func main() {
// 	fmt.Println(f1())
// }
// func f1() int {
// 	return f2()
// }
// func f2() int {
// 	return 1
// }

// ----

// ---- call stack option 2 ----

// func main() {
// 	fmt.Println(f1())
// }
// func f1() int {
// 	return f2()
// }
// func f2() (r int) {
// 	r = 34
// 	return
// }

// ----

// ---- Возврат нескольких значений ----

// func f() (int, int) {
// 	return 5, 6
// }

// func main() {
// 	x, y := f()
// 	fmt.Println(x, y)
// }

// ----

// ---- Variable number of function arguments ----

// func add(args ...int) int {
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }
// func main() {
// 	fmt.Println(add(1, 2, 3))
// }

// ---- Closures ----

// func main() {
// 	add := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(add(1, 1))
// }

// ----

// ---- Closures Option 2 ----

// func main() {
// 	x := 0
// 	increment := func() int {
// 		x++
// 		return x
// 	}
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

// ----

// ---- Closures Option 3 ----

// func makeEvenGenerator() func() uint {
// 	i := uint(0)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }
// func main() {
// 	nextEven := makeEvenGenerator()
// 	fmt.Println(nextEven()) // 0
// 	fmt.Println(nextEven()) // 2
// 	fmt.Println(nextEven()) // 4
// }

// ----

// ---- deferred call ----

// func first() {
// 	fmt.Println("1st")
// }
// func second() {
// 	fmt.Println("2nd")
// }
// func main() {
// 	defer second()
// 	first()
// }

// ---- panic and recovery ----

// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("This is PANIC")
// }

// ---- examples of functions from the course DTA ----

// ---- simple example ----
// func main() {
// 	Print()

// }

// func Print() {
// 	fmt.Printf("Print something")
// }

// ----

// ---- simple example 2 ----

// func main() {
// 	Print("Print something\n")
// 	Print("I'm very smart")

// }

// func Print(input string) {
// 	fmt.Printf(input)
// }

// ----

// ---- function with indefinite number of arguments ----

func main() {
	summa(1, 2, 3)
}

func Print(input string) {
	fmt.Println(input)
}

func sum(a, b int) {
	fmt.Println(a + b)
}

func summa(arg ...int) {
	var total = 0
	for _, value := range arg {
		total += value
	}
	fmt.Println(total)
}

//  ----
