package main

import "fmt"

/*
two pointers: use when
 - array is sorted
 - need pairs or triples
 - need to compress or remove in place
 - compare left/right ends
linear time (compared to quadratic for nested for loops)
*/
func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		}else{
			right--
		}
	}
	return nil
}

/*
sliding window:
 - problem asks about a subarray
 - often says longest, shortest, at most, sum/product, condition
 - usually for contiguous ranges
*/
func maxSumSubarrayK(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	windowSum := 0 
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for right := k; right < len(nums); right++ {
		windowSum += nums[right]
		windowSum -= nums[right-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func moveZeroes(nums []int) {
	write := 0

	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}

	for write < len(nums) {
		nums[write] = 0
		write++
	}
}

func main(){
	fmt.Println(twoSumSorted([]int{2,7,11,15}, 9))
	fmt.Println(maxSumSubarrayK([]int{2,1,5,1,3,2}, 3))
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

/* 
common array patterns: most array problems can be reduced to one of these
1. two pointers
2. sliding window
3. prefix sum
4. hash map or frequency counting 
5. sorting + scanning
6. in-place modification
7. binary search on array
8. greedy scan
9. kadane's algorithm
10. intervals or merging
*/
/*binary search
left, right = 0, len(nums)-1
for left <= right {
	mid := left + (right - left)/2
}
*/
/**/
/**/
/**/
/**/
/**/