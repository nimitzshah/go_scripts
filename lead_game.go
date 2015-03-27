package main
import "fmt"
// import "math"
// import "bufio"
// import "os"
// import "strconv"
func main(){
  var inputs int;
  fmt.Scanf("%d",&inputs)
  var lead int =0;
  var blue_team int;
  var red_team int;
  var blue_team_total int = 0;
  var red_team_total int =0;
  var winner int;
  for input:=0;input<inputs;input++{
    fmt.Scanf("%d %d",&blue_team,&red_team);
    blue_team_total += blue_team
    red_team_total  += red_team
    if(blue_team_total>red_team_total && blue_team_total-red_team_total > lead){
      lead = blue_team_total-red_team_total;
      winner =1;
      }else if(red_team_total>blue_team_total && red_team_total-blue_team_total > lead){
      lead = red_team_total - blue_team_total;
      winner =2;
    }
  }
  fmt.Println(winner,lead)
}
