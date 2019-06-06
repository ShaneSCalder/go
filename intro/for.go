package main

import "fmt"

func main(){

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for b := 7; b <= 10; b++ {
    fmt.Println(b)
  }

  for X := 0; X <= 5; X++ {
    if X%2 == 0 {
      continue
    }
    fmt.Println(X)
  }
}
