package main

import (
	"fmt"
	"log"
	"solution-go/solution"
)

func main() {
	/*
		Solution 1
	*/
	a := 50
	b := 63
	fmt.Printf("Before: a = %d after b = %d\n", a, b)

	solution.SwapVariable1(&a, &b)
	fmt.Printf("After: a = %d after b = %d\n", a, b)

	/*
		Solution 2
	*/
	orgStr := "jatis"
	fmt.Printf("Original: %s, Reversed: %s\n", orgStr, solution.ReverseString(orgStr))

	/*
		Solution 3
	*/
	orgStr = "dani Maulana"
	res, err := solution.GenerateNumberString(orgStr)
	if err != nil {
		log.Fatalf("Error : %+v", err)
	}
	fmt.Printf("Original: %s, Result: %s\n", orgStr, res)

	/*
		Solution 4
	*/
	solution.PrintPrime()
}
