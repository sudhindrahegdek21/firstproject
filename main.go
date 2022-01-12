package main

import (
	"fmt"
	"strconv"
	"time"
)

var digitsum int

//adding digits in integer first method
func addDigits(number int) int {
	digitsum = 0
	if number == 0 {
		return 0
	}
	digitsum = number%10 + addDigits(number/10)
	return digitsum
}

//getting sum of Digits in a integer string
func AddDigitsString(str string) int {
	sum, i := 0, 0
	// Traversing through the string
	for i < len(str) {
		// Since ascii value of
		// numbers starts from 48
		// so we subtract it from sum
		sum = sum + int(str[i]) - 48
		i++
	}
	return sum
}

// To check String a valid integer
func IsDigitsOnly(s string) bool {
	if s[0] == '0' {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func SliceDisplay(str string) {
	fmt.Println("string Display...")
	for _, s := range str {
		fmt.Printf("%c, ", s)
	}
}

func main() {

	fmt.Println()
	//first method...
	number := 0

	fmt.Println("Enter Your number")
START:
	_, err := fmt.Scanln(&number)
	if err != nil {
		goto START
	}

	sum := addDigits(number)
	start := time.Now()
	fmt.Println(time.Since(start))
	fmt.Println(sum)
	for sum > 9 {
		sum = addDigits(sum)
		fmt.Println(sum)
	}
	//second method...
	fmt.Println("string method...")
	digitnum := ""
	fmt.Println("integer as string")
AGAIN:
	fmt.Scanln(&digitnum)
	if _, err := strconv.ParseInt(digitnum, 10, 64); err == nil {
		fmt.Printf("%q looks like a number.\n", digitnum)
		sum = AddDigitsString(digitnum)
		fmt.Println(sum)
		for sum > 9 {
			sum = addDigits(sum)
			fmt.Println(sum)
		}
	} else {
		goto AGAIN
	}
	//third method
	snumber := ""
REPEAT:
	fmt.Println("Enter a valid number..")
	fmt.Scanln(&snumber)
	if !IsDigitsOnly(snumber) {
		goto REPEAT
	}
	sum = AddDigitsString(snumber)
	start = time.Now()
	fmt.Println(time.Since(start))
	fmt.Println(sum)
	for sum > 9 {
		sum = addDigits(sum)
		fmt.Println(sum)
	}
	// Display in slices
	snumber = ""
CYCLE:
	fmt.Println("Enter a valid number to Display..")
	fmt.Scanln(&snumber)
	if !IsDigitsOnly(snumber) {
		goto CYCLE
	}
	SliceDisplay(snumber)

}
