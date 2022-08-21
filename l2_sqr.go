package main

import "fmt"

func main() {
  var wdh, hgh, sqr int

  fmt.Println("Введите длину:")
  fmt.Scanln(&wdh)

  fmt.Println("Введите ширину:")
  fmt.Scanln(&hgh)

  sqr = wdh * hgh
  fmt.Println("Площадь равна : ",sqr)
}
