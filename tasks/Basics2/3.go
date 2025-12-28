package main

import "fmt"

type order struct {
	sum    int
	status string
}

func main() {
	orderList := []order{
		{
			sum:    100,
			status: "completed",
		},
		{
			sum:    50,
			status: "pending",
		},
		{
			sum:    13123,
			status: "pending",
		},
		{
			sum:    32,
			status: "completed",
		},
		{
			sum:    11231,
			status: "completed",
		},
	}

	conditionCompleteList := make([]order, 0)
	var boolStatus bool

	for _, value := range orderList {
		if value.sum > 60 && value.status == "completed" {
			conditionCompleteList = append(conditionCompleteList, order{sum: value.sum, status: value.status})
			boolStatus = true
		} else {
			boolStatus = false
		}
	}
	
	if boolStatus == true {
		fmt.Println(conditionCompleteList)
	} else {
		fmt.Println("Заказа нет")
	}
}
