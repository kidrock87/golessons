package main

import "fmt"
import "math"

func main() {
  var sqr, rad, lgh float64
  pi := math.Pi

  fmt.Println("Введите площадь круга:")
  fmt.Scanln(&sqr)

  rad = math.Sqrt(sqr/pi)
  fmt.Println("Радиус равен: ",rad)

  lgh = math.Sqrt(sqr*4*pi)
  fmt.Println("Длина окружности равна: ",lgh)

}
