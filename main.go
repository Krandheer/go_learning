package main

func fib(n int, ch chan int){
	x,y :=0,1
	for i:=0; i<n; i++{
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}


func main() {
	// channel keeps listening for value until it is closed, we have length 10 channel but we keep receiving values
	// so it is not getting overflown
	ch := make(chan int, 10)
	go fib(15, ch)
	for i := range ch{
		println(i)
	}
}