package main

import (
        "fmt"
        "Go_Be_Geek/db"
        "Go_Be_Geek/log"
)

func main() {
  db.HelloDB()
  log.HelloLog()
  db.HelloBase()
  fmt.Println(db.Begeek)
}
