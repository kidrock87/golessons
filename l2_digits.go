package main

import "fmt"

func main() {
  var dgt int

  fmt.Println("Введите число:")
  fmt.Scanln(&dgt)
  fmt.Println("Сотен : ", dgt / 100)
  fmt.Println("Десятков : ", (dgt % 100) / 10)
  fmt.Println("Единиц : ", dgt % 10)
}
