package main

import (
	"fmt"
	"os"
	"syscall"
	"os/signal"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	time.AfterFunc(time.Second, func() {
	
		fmt.Println("done")
	
	})
	
	<-sigs
}