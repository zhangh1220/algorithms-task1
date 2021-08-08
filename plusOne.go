//加一
//基本思路 遍历切片末尾加一 如果不为9直接加一  如果是9 改为0 下次循环再加一
package main

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{1}
	}
	for i := n-1 ;i>=0;  i++  {
		//如果不为9直接加1
		if digits[i] != 9 {
			digits[i]++
			return digits
		}else {
			digits[i] = 0
		}
	}
	return append([]int{1},digits...)
}