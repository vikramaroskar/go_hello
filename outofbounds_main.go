package main


import "math/rand"

func 4main() {
   list := []int{1, 2, 3}

   printList(list)
   
   
   num := getRandom()
   println(*num)
}

func printList(list []int) {
   println(list[2])
   println(list[1])
   
}


//go:noinline
func getRandom() *int {
   tmp := rand.Intn(100)

   return &tmp
}