/*Two Sum
Easy Topics Company Tags
Hints

Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first. */

func twoSum(nums []int, target int) []int {
    numdic := make(map[int]int)
	for i:=0;i<len(nums);i++{
		
		num2:=target-nums[i]
		index2, ok := numdic[num2]
		if ok {
		return []int{index2,i}
		}
		numdic[nums[i]]=i
	}
	return []int{1, 2}
}
