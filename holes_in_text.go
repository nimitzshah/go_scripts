package main

import "fmt"
import "bufio"
import "os"

func main(){
  var trials int
  fmt.Scanf("%d",&trials)
  holes := map[rune]int{'A':1,'B':2,'D':1, 'O':1, 'P':1, 'R':1, 'Q':1}
  count_of_holes := 0
  scanner := bufio.NewReader(os.Stdin)
  for trial:=0; trial< trials;trial++ {
    count_of_holes = 0
    text,_,_ := scanner.ReadLine()
    runes := string(text[:])
    for _,runeValue :=range runes{
      count_of_holes += holes[runeValue]
    }
    fmt.Println(count_of_holes)
  }
}
