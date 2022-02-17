package main

import "fmt"

func main() {
  // var mass string
  // mass = "I'm watching you!"
  // massage := "Hi my World"

  var name string
  var age int
  var weight float32
  // var isAdult bool

  name = "Максим"
  weight = 70.2
  age = 20

  printPersonInfo(name, age, weight)
  printPersonInfo("Вася", 46, 92)

  // isAdult := isAdult(15)
  // fmt.Println(isAdult)

  fmt.Println(isAdult(15))


  // var massage = "Hi my World"
  // fmt.Println(massage, mass)
  // fmt.Println(name)
  // fmt.Println(age)
  // fmt.Println(weight)
  // fmt.Println(isAdult)
}

func printPersonInfo(name string, age int, weight float32) {
  fmt.Printf("Имя: %s\nВозраст: %d\nВес: %.2f\n", name, age, weight)
}

// func isAdult(age int) bool {
//   if age >= 18 {
//       return true
//   }
//   return false
// }

func isAdult(age int) bool {
  return age >= 18
}
