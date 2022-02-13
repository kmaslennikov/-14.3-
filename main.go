package main

import (
	"fmt"
)

func main() {
  var x,k,c int
	fmt.Println("Введите число:")
  fmt.Scan(&x)
  fmt.Println("Введите коэффициент:")
  fmt.Scan(&k)
  fmt.Println("Введите слагаемое:")
  fmt.Scan(&c)

  result:=add(multiply(x,k),c)
  fmt.Println("Результат:", result)

}

func multiply(x,k int) (a int){
  a=k*x
  return
}
func add(x,c int)(a int){
  a=x+c
  return
}