package main

import "fmt"

func main() {
	fmt.Println("welcome to loop in golang")
	var days = []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for v := range days {
	// 	fmt.Println(days[v])
	// }

	for index, v := range days {
		fmt.Printf("index is %v and day is %v\n", index, v)
	}
}
