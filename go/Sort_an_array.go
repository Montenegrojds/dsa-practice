/*
Sort an Array
Medium Topics Company Tags

You are given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Example 1:

Input: nums = [10,9,1,1,1,2,3,1]

Output: [1,1,1,1,2,3,9,10]

Example 2:

Input: nums = [5,10,2,1,3]

Output: [1,2,3,5,10]
*/



func sortArray(nums []int) []int {
    if len(nums)<=1{
        return nums
    }
    pivot := nums[len(nums)/2]
    l,m,r:= []int{},[]int{},[]int{}
    for _,val:= range(nums){
        if val<pivot{
            l=append(l,val)
        }
    }

    for _,val:=range(nums){
        if val>pivot{
            r=append(r,val)
        }
    }
    for _,val:=range(nums){
        if val==pivot{
            m=append(m,val)
        }
    }
    re := append(sortArray(l), m...)
    return append(re, sortArray(r)...)
}
