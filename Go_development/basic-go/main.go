package main


import (
  "fmt"
)

func main() {
  contactsList := [3]string{"Sam", "Tom", "Ann"}

  fmt.Println("Contacts list:")
  for index := range contactsList {
      fmt.Printf("%d. %s\n", index+1, contactsList[index])
  }
}
