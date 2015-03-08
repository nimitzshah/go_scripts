package main

import "fmt"

func main() {
  var transaction int
  var bank_balance float64
  fmt.Scanf("%d %f", &transaction, &bank_balance)
  if transaction%5 ==0 && bank_balance >= (float64(transaction) + 0.50){
    bank_balance = bank_balance - (float64(transaction) + 0.50)
  }
  fmt.Printf("%.2f",bank_balance)
}
