package main

import "fmt"

func main() {
  name := "Sam"
  fmt.Println(name)

  setName(&name)

  fmt.Println(name)

  var emptyPointer *string
  fmt.Println(emptyPointer)
  // pName := &name
  //
  // fmt.Println(name)
  // fmt.Println(pName)
  //
  // *pName = "Tom"
  // fmt.Println(name)
}

func setName(name *string) {
  *name = "Tom"
}
