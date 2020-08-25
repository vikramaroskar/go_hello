package main

import (
	"fmt"
	
	
	"example.com/user/hello/morestrings"
	"example.com/user/hello/streamcode"
	
	"github.com/google/go-cmp/cmp"
	
	
)

func 1main() {
	fmt.Println("Hello, world.")
	
	
	/* reverse and diff */
	
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	
	
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	
	
	
	
	/*Stream code */
	fmt.Println("------------------------------")
	
	streamcode.Streamcode()	

	fmt.Println("------------------------------")	
	streamcode.LimitReaderFunc()
	
	
	fmt.Println("------------------------------")	
	fmt.Println("Bounds check")	
   list := []int{1, 2, 3}

   printList(list)
}

func printList(list []int) {
   println(list[2])
   println(list[3])
}
