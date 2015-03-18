package main

import "math/big"
import "fmt"

func main(){
  var test_cases int;
  fmt.Scanf("%d",&test_cases)
  var integer int64;
  product := big.NewInt(2);
  for test_case:=0;test_case< test_cases;test_case++{
    fmt.Scanf("%d", &integer)
    product = factorial(integer)
    fmt.Println(product)
  }

}

func factorial(integer int64) (product *big.Int){
  product = big.NewInt(integer);
  for integer>1 {
    integer -=1
    product.Mul(product,big.NewInt(integer))
  }
  return
}
