//Took heavy help from Aurum Potabile to understand how to convert
// byte to integer

package main
import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
  var inputs int;
  fmt.Scanf("%d", &inputs)
  list := make([]int,1000001)
  scanner := bufio.NewReader(os.Stdin)
  printer := bufio.NewWriter(os.Stdout)
  for input:=0;input <inputs;input++{
    value := input_to_number(scanner)
    list[value]++
  }
  for value,count := range list{
    for iterator:=0;iterator< count; iterator++{
      printer.WriteString(strconv.Itoa(value) + "\n")
    }
  }
  printer.Flush()
}

func input_to_number(scanner *bufio.Reader) int{
  var number int;
  for input_byte,_:= scanner.ReadByte();input_byte <= '9' && input_byte >= '0';input_byte,_ = scanner.ReadByte(){
    number = 10*number + int(input_byte-'0')
  }
  return number;
}
