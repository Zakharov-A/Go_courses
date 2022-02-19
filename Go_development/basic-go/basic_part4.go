package main


import (
  "fmt"
)

func main() {
  contactsList := [4]string{"Sam", "Tom", "Ann", "Jim"}

  fmt.Println("Contacts list:")
  for index, value := range contactsList {
      fmt.Printf("%d. %s\n", index+1, value)
  }

  for i := 0; i < len(contactsList); i ++{
    fmt.Printf("%s ", contactsList[i])
  }
  // for i := 0; i <= 10; i++{
  //     fmt.Printf("%d ", i)
  // }
}
