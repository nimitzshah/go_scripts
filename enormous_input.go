package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
  var divisor int;
  var inputs int;
  total :=0
  fmt.Scanf("%d %d",&inputs, &divisor)
  scanner := bufio.NewReader(os.Stdin)
  for input:= 0;input <inputs;input++ {
    input_string,_,_ := scanner.ReadLine()
    dividend,_:= strconv.Atoi(string(input_string))
    if(dividend%divisor ==0){
      total= total +1
    }
  }
  fmt.Println(total)
}
