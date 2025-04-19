package main

import "fmt"

// func fib(n int, ch chan int){
// 	x,y :=0,1
// 	for i:=0; i<n; i++{
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }

func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("Quit")
			return
		}
	}
}

func main() {

	// channel keeps listening for value until it is closed, we have length 10 channel but we keep receiving values
	// so it is not getting overflown
	// ch := make(chan int, 10)
	// go fib(15, ch)
	// for i := range ch{
	// 	println(i)
	// }

	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 15; i++ {
	// 		println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fib2(c, quit)

	// whomAmI := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		println("int")
	// 	case string:
	// 		println("string")
	// 	default:
	// 		println("unknown")
	// 	}
	// }
	// whomAmI("randheer")

	nums := make([]int, 0)
	fmt.Println(nums)
    nums2 := []int{}
    var nums3 []int
	fmt.Println(nums == nil)
	fmt.Println(nums2 == nil)
	fmt.Println(nums3 == nil)

}
