package main

import (
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Name string `json:"name"`
  Age uint16 `json:"age"`

}

func main() {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goland")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  //Добавление данных, конкретно пользователя
  // insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Pipa', 18)")
  // if err != nil {
  //   panic(err)
  // }
  // defer insert.Close()
  // fmt.Println("Подключено к MySQL")

// Выборка данных

  res, err := db.Query("SELECT `name`, `age` FROM `users`")
  if err != nil {
    panic(err)
  }

  for res.Next() {
    var user User
    err = res.Scan(&user.Name, &user.Age)
    if err != nil {
      panic(err)
    }

    fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))

  }
}
