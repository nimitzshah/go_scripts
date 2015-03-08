package main

import "fmt"

func main() {
  var number int
  for number!= 42 {
    fmt.Scanf("%d", &number)
    if number!=42{fmt.Println(number)}
  }
}
