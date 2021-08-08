//移动零
//基本思路:双指针法
package main

func moveZeroes(nums []int) {
	i, j, n := 0, 0, len(nums)
	for j < n {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
