package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()

	fmt.Println("Current Date is : ", presentTime)
	fmt.Println("Current Date is : ", presentTime.Format("01-02-2006 Monday 15:04:05")) // default format which is used by time package in go

	createdDate := time.Date(2020, time.August, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created date is: ", createdDate)
	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 Monday 15:04:05"))
}
