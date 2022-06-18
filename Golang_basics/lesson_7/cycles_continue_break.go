package main

import "fmt"

func main() {
	fmt.Println("Hi")

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		// fmt.Println(i)
	}

label1:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			// fmt.Println("I:", i, "j", j)
			if i >= 10 {
				continue label1
			}
		}
	}

	for i := 0; i <= 20; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

label2:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I:", i, "j", j)
			if i >= 10 {
				break label2
			}
		}
	}

}
