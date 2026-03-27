package main

import "fmt"


//1d dp template
func solve(nums []int) int {
	n := len(nums)
	if n ==0 {
		return 0
	}

	dp := make([]int, n)
	//base case
	//dp[0]= ...
	for i:= 1; i<n; i++ {
		//dp[i] = ...
	}
	return dp[n-1]
}
//2d dp template
func solve(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
}

func main(){
	fmt.Println("dp go")
}