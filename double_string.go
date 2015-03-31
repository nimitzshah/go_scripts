package main

import "fmt"

func main() {
  var inputs int;
  var palindrome_size int
  fmt.Scan(&inputs)
  for input:=0;input<inputs;input++ {
    fmt.Scan(&palindrome_size)
    if palindrome_size%2!=0{
      palindrome_size = palindrome_size -1
    }
    fmt.Println(palindrome_size)
  }
}
