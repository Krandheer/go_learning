package main

import (
	"fmt"
	"sort"
)

func main() {
	var intArr = []int{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(intArr[len(intArr)-1], len(intArr), cap(intArr))
	// intArr = append(intArr, 4)

	type Prize struct {
		ClaimType  string `json:"claim_type"`
		Amount     int    `json:"amount"`
		MaxWinners int    `json:"max_winners"`
	}

	prizes := []Prize{
		{"full_house", 1000.0, 2},
		{"full_house", 3000, 3},
		{"full_house", 800, 1},
		{"early_five", 400, 2},
		{"bottom_line", 300, 2},
		{"top_line", 300, 3},
		{"early_five", 300, 2},
		{"middle_line", 200, 2},
	}

	claimTypeOrder := map[string]int{
		"full_house":  1,
		"top_line":    2,
		"middle_line": 3,
		"bottom_line": 4,
		"early_five":  5,
	}

	// d, ok := claimTypeOrder["early_five"]
	// if ok {
	// 	fmt.Printf("value found %v \n", d)
	// } else {
	// 	fmt.Println("value was not found")
	// }

	sort.Slice(prizes, func(i, j int) bool {
		if prizes[i].Amount != prizes[j].Amount {
			return prizes[i].Amount > prizes[j].Amount
		}
		return claimTypeOrder[prizes[i].ClaimType] < claimTypeOrder[prizes[j].ClaimType]
	})
}
