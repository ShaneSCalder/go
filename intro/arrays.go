package main

import "fmt"

func main() {
  var a [5]int
  fmt.Println("emp ", a)

  fmt.Println("get ", a[4])

  b := [5]int{1,2,3,4,5}
  fmt.Println("Dir ", b)

  fmt.Println("get Dir ", b[3])
}
