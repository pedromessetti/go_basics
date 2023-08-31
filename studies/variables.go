package main

import "fmt"

func main() {

	// Unitialized declared variables
	// var smsSendingLimit int // this will be 0
	// var costPerSms float64 // this will be 0.0
	// var hasPermission bool // this will be false
	// var senderName string // this will be ""

	// Declaring and initializing variables using short assignment operator
	smsSendingLimit := 1000
	costPerSms := 0.25
	hasPermission := true
	senderName := "Pedro Messetti"

	fmt.Printf(
		"SMS Sending Limit: %d\nCost Per SMS: %.2f\nHas Permission: %t\nSender Name: %s\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		senderName,
	)

	// Declaring multiple variables of same type
	var a, b, c int // all of then will be initializes as 0
	fmt.Printf("\na: %d\nb: %d\nc: %d\n", a, b, c)

	// Attributing values to multiple variables
	a, b, c = 1, 2, 3

	fmt.Printf("\na: %d\nb: %d\nc: %d\n", a, b, c)
}
