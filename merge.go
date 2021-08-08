//合并两个数组
//基本思路：//从nums1数组后面向前插入元素  特别注意处理 nums1 = [0]  nums2 = [1] 的情况
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		if i == -1 {
			nums1[k] = nums2[j]
			j--
		} else if j == -1 {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
}
