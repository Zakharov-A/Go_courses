package main

import (
  "fmt"
  "math"
)

func main() {
  var a float64
  var b float64
  var c float64

  fmt.Println("Quadratic equation")

  fmt.Println("Введи a:")
  fmt.Scan(&a)

  fmt.Println("Введи b:")
  fmt.Scan(&b)

  fmt.Println("Введи c:")
  fmt.Scan(&c)

  D := (b*b) - 4*(a*c)

  if D > 0 {
    var x1 float64
    var x2 float64

    x1 = (-b + math.Sqrt(D)) /(2*a)
    x2 = (-b - math.Sqrt(D)) /(2*a)

    fmt.Println("Ваше уровнение имеет 2 корня\nD = " + fmt.Sprint(D))
    fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
  } else if D == 0 {
      var x float64
      x = (-b) / (2*a)
      fmt.Println("Ваше уравнение имеет 1 корень\nD")
      fmt.Println("X: " + fmt.Sprint(x))
  } else if D < 0 {
      fmt.Println("Ваше уровнение не имеет корней")
  }



  // num := 3
  // if num < 0 {
  //   fmt.Println("Number is greater then 0")
  // } else if num < 0 {
  //   fmt.Println("Number less than 0")
  // } else if num == 3 {
  //   fmt.Println("Number equals 3!")
  // }

}
