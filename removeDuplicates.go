//去除重复项
//基本思路:双指针法
package main

func removeDuplicates(nums []int) int {
	n:= len(nums)
	if n <= 1 {
		return n
	}
	i:= 1
	for j := 1; j<n ;j++ {
		if nums[j]!= nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}