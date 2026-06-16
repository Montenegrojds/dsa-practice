/*
Longest Consecutive Sequence
Medium Topics Company Tags
Hints

Given an array of integers nums, return the length of the longest consecutive sequence of elements that can be formed.

A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element. The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [2,20,4,10,3,4,5]

Output: 4

Explanation: The longest consecutive sequence is [2, 3, 4, 5].

Example 2:

Input: nums = [0,3,2,5,4,6,1,1]

Output: 7

Constraints:

    0 <= nums.length <= 1000
    -10^9 <= nums[i] <= 10^9

*/

func counter(hmap map[int]bool,num int)int{
	maxCounter:=1
	for true{
		if hmap[num+1]==true{
			maxCounter++
			num++
		}else{
			return maxCounter
		}
	}
	return maxCounter
}
func longestConsecutive(nums []int) int {
hmap:=map[int]bool{}
maxCount:=0
for i:=0;i<len(nums);i++{
	hmap[nums[i]]=true
}

for i:=0;i<len(nums);i++{
	if hmap[nums[i]-1]==true{
		continue
	}else{
		tempCount:=counter(hmap,nums[i])
		if maxCount<tempCount{
			maxCount=tempCount
		}
	}
}
return maxCount
}
