package main

import "fmt"

func topKFrequent(nums []int, k int)[]int {
	freq := make(map[int]int)
	for _, x := range nums {
		freq[x]++
	}

	buckets := make([][]int, len(nums)+1)
	for x, f := range freq {
		buckets[f] = append(buckets[f], x)
	}

	res := make([]int, 0, k)
	for f := len(buckets) - 1; f >= 0 && len(res) < k; f-- {
		res = append(res, buckets[f]...)
	}

	return res[:k]

}

//frequency counting, e.g. contains duplicate
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

//valid anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println("hash table go")

	//syntax
	m := make(map[int]int)
	s := make(map[string]bool)
	//check existence
	val, ok := m[key]
	if ok {
		fmt.Println("exists", val)
	}
	//increment count
	m[x]++
	//delete
	delete(m, key)
	//set pattern
	seen := make(map[int]struct{})
	seen[x] = struct{}{}
	if _, ok := seen[x]; ok {
		//exists
	}

}

