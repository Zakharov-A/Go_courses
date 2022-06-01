package main
import "fmt"

// type person struct{
//   name string
//   surname string
//   age int
// }
//
// func main()  {
//   var sam = person {"Sam", "Disel", 25}
//   var li = person {"Li", "Sin", 45}
//   fmt.Println(sam.name)
//   fmt.Println(sam.surname)
//   fmt.Println(sam.age)
//   sam.surname = "Vanhellsing"
//   fmt.Println(sam)
//   fmt.Println(li)
//
//
// }

// ----

type person struct{
  name string
  surname string
  age int
}

func main() {
  red := person {"Zed", "Tramp", 44}
  var redPointer *person = &red
  redPointer.age = 39
  fmt.Println(red.age)
  (*redPointer).age = 33
  fmt.Println(red.age)
  fmt.Println(red)
  var tom *person = &person{name:"Tom", age:23}
  fmt.Println(tom)
  var bob *person = new(person)
  fmt.Println(bob)


}
