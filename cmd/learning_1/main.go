package main

import "fmt"

func main() {
	var intArr = []int{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(intArr[len(intArr)-1], len(intArr), cap(intArr))
    intArr = append(intArr, 4)
	fmt.Println(intArr)
}
