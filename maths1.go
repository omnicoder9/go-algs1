package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("maths go")
	var rr = 99
	var ww = strconv.Itoa(rr)
	fmt.Printf("\nthe number is %v\n", ww)
	mid := 44
	n := 23

if int64(mid)*int64(mid) <= int64(n) {
	//safe
}

//avoid floating point unless truly needed
}

/*
categories of maths problems
digits: reverse integer, palindrome number, sum of digits, count digits
divisibility/modular arithmetic: FizzBuzz, happy number, power checks
GCD/LCM: fraction simplification, array step sizes, syncing cycles
prime-related: count primes, prime factors, primality checks
powers/logarithmic decomposition: power of two/three/four
combinatorics: number of ways, handshakes, pairs
geometry/coordinates: slope, distance, area, collinearity
simulation with math shortcut: instead of simulating 1e9 steps, derive formula
number theory patterns: remainders, parity, digit cycles

Can I solve this with a formula or invariant instead of brute force?

*/

//use int64 if multiplication may get large
func square(x int) int64 {
	return int64(x) * int64(x)
}


//digit patterns: use % 10 and / 10
//reverse digits
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	res := 0
	for x > 0 {
		d := x % 10
		res = res*10 + d
		x /= 10
	}
	return sign * res
}

//palindrome number without string conversion
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	orig := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return orig == rev
}

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += n 
	}
	return count
}

//excel column number
func titleToNumber(columnTitle string) int {
	res := 0
	for _, ch := range columnTitle {
		res = res*26 + int(ch-'A'+1)
	}
	return res
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Return two values.
}