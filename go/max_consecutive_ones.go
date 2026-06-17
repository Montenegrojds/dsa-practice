/*
Max Consecutive Ones
Easy Topics Company Tags

You are given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:

Input: nums = [1,1,0,1,1,1]

Output: 3


Example 2:

Input: nums = [1,0,1,1,0,1]

Output: 2

Constraints:

    1 <= nums.length <= 100,000
    nums[i] is either 0 or 1.


*/

func findMaxConsecutiveOnes(nums []int) int {
	counter:=0
	maxCounter:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==1{
			counter++
		}else{
			fmt.Println(maxCounter,counter,i)
			if maxCounter<counter{
				maxCounter=counter
			}
			counter=0
		}
	}
	if counter>maxCounter{
		maxCounter=counter
	}
	return maxCounter
}
