//旋转数组
package main

func rotate(nums []int, k int)  {
	k%=len(nums)
	swap(nums)
	swap(nums[:k])
	swap(nums[k:])

}
func swap(nums []int) {
	n:= len(nums)
	for i:= 0;i<n/2; i++ {
		nums[i], nums[n-1-i]= nums[n-1-i],nums[i]
	}
}