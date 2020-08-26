package datarace


import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int;

}

func Datarace() {

	var v atomic.Value

	
	
	//writer
	go func() {
		var i int
		for {
			i++
			//cfg.a = []int{i, i+1, i+2, i+3, i+4, i+5}
			
			cfg := &Config{
				a: []int{i, i+1, i+2, i+3, i+4, i+5},
			}
			v.Store(cfg)
		}
		
	
	} ()
	
	
	// reader
	var wg sync.WaitGroup
	for n:=0; n< 4; n++ {
		wg.Add(1)
		go func() {
			for n:=0; n<10; n++ {
				cfg := v.Load()
				//fmt.Printf("%#cfg\n")
				fmt.Printf("%+v\n", cfg)
			
			}
			wg.Done()
		}()
	}
	
	wg.Wait()
	

}