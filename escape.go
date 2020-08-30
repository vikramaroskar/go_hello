package main


import (
	"fmt"
)

func main() {
   var number *int
   for i := 0; i < 3; i++ {
      number = new(int)
      *number = i
   }
   fmt.Println(*number)
   
}