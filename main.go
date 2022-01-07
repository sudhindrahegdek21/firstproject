package main

import "fmt"

var digitsum int

func addDigits(number int) int {
	digitsum = 0
	if number == 0 {
		return 0
	}
	digitsum = number%10 + addDigits(number/10)
	return digitsum
}

func main() {
	number := 0
	fmt.Println("helllo")
	fmt.Println("Enter Your number")
	fmt.Scanln(&number)
	fmt.Println(number)
	sum := addDigits(number)
	fmt.Println(sum)
	for sum > 9 {
		sum = addDigits(sum)
		fmt.Println(sum)
	}

}
