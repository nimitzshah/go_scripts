package main

import "fmt"

func main() {
  var trials int
  var factorial int
  fmt.Scanf("%d",&trials)
  for trial:=0; trial<trials;trial++{
    fmt.Scanf("%d",&factorial)
    var multiples_of_five int;
    multiples_of_five =5
    var quotient int =1;
    var number_of_zeroes int =0;
    for quotient >=1{
      quotient = factorial / multiples_of_five
      multiples_of_five = multiples_of_five *5
      number_of_zeroes = number_of_zeroes + quotient
    }
    fmt.Println(number_of_zeroes)
  }
}
