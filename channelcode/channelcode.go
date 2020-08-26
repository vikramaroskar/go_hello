package channelcode


import (
	"fmt"
	"time"
)


func SendValue (c chan int	) {
	c <- 8
}

func SendValueString (c chan string) {

	fmt.Println("executing stiring send values")
	time.Sleep(2 *time.Second)
	c <- "first string"
	fmt.Println("Done  stiring send values")
}

func Channelcode() {
	
	values := make (chan int, 5)
	
	defer close(values)
	
	go SendValue(values)
	
	
	//blocking until channel gets values
	
	receivedvalue := <-values
	
	fmt.Println("receivedvalue", receivedvalue)
	
	
	schannel := make (chan string, 5)
	
	go SendValueString(schannel)
	
	
	//blocking until channel gets values
	
	receivedvaluestring := <-schannel
	
	fmt.Println("receivedvaluestring", receivedvaluestring)
	

}