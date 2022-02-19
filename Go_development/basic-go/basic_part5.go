package main


import (
  "fmt"
)

func main() {

  // Масивы и слайсы
  // contactsList := []string{"Sam", "Tom", "Ann", "Jim"}
  //
  // fmt.Println("Contacts list:")
  // for index, value := range contactsList {
  //     fmt.Printf("%d. %s\n", index+1, value)
  // }
  //
  // contactsList = append(contactsList, "Luka", "Jon", "Andrey")
  //
  // for i := 0; i < len(contactsList); i ++{
  //   fmt.Printf("%s ", contactsList[i])
  // }
  // for i := 0; i <= 10; i++{
  //     fmt.Printf("%d ", i)
  // }

      contactsList := map[string]string{
        "Sam": "123",
        "Tom": "456",
        "Ann": "789",
        "Jim": "999",
         }
      fmt.Println("До delete()")
      for name, phoneNumber := range contactsList {
          fmt.Println(name, phoneNumber)
        }

        delete(contactsList, "Sam")

        fmt.Println("После delete()")
        for name, phoneNumber := range contactsList {
            fmt.Println(name, phoneNumber)
          }



}
