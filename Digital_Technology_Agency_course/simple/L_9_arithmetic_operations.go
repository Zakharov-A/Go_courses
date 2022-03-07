package main

import "fmt"

// ---- Arithmetic operations ----

func main() {
	var one = 1
	var two = 2
	var six = 6
	fmt.Printf("[+] 1 + 2 =[%d] \n", one+two) // -- Addition --
	fmt.Printf("[-] 2 - 1 =[%d] \n", two-one) // -- Subtraction --
	fmt.Printf("[*] 2 - 1 =[%d] \n", two*two) // -- Multiplication --
	fmt.Printf("[/] 6 / 2 =[%d] \n", six/two) // -- Division --
	fmt.Printf("[%%] 35 %% 3 =[%d] \n", 35%3) // -- Remainder of the division --

	// ---- Increment and decriment ----
	two++
	fmt.Printf("[++] 2 = [%d] \n", two) // -- Increment --
	two--
	fmt.Printf("[--] 2 = [%d] \n", two) // -- Decriment --

}

// ----
