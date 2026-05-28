/*
3Sum
Medium Topics Company Tags
Hints

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] where nums[i] + nums[j] + nums[k] == 0, and the indices i, j and k are all distinct.

The output should not contain any duplicate triplets. You may return the output and the triplets in any order.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]

Output: [[-1,-1,2],[-1,0,1]]

Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].

Example 2:

Input: nums = [0,1,1]

Output: []

Explanation: The only possible triplet does not sum up to 0.

Example 3:

Input: nums = [0,0,0]

Output: [[0,0,0]]

Explanation: The only possible triplet sums up to 0.

*/
func threeSum(nums []int) [][]int {
sort.Ints(nums)
lst:=[][]int{}
for i:=0;i<len(nums);i++{
    if i>0 && nums[i]==nums[i-1]{
        continue
    }
    l:=i+1
    r:=len(nums)-1
    for l<r{
        sum:= nums[i]+nums[l]+nums[r]
        if sum>0{
            r--
        }else if sum<0 {
            l++
        }else{
            lst=append(lst,[]int{nums[i],nums[l],nums[r]})
            l++
             for l<r && nums[l]==nums[l-1]{
                l++
            }
        }    
    }

}
return lst

}
