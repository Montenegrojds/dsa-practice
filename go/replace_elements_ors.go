/*
Replace Elements With Greatest Element On Right Side
Easy Topics Company Tags

You are given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.

Example 1:

Input: arr = [2,4,5,3,1,2]

Output: [5,5,3,2,2,-1]

Example 2:

Input: arr = [3,3]

Output: [3,-1]

Constraints:

    1 <= arr.length <= 10,000
    1 <= arr[i] <= 100,000

*/

func replaceElements(arr []int) []int {
maxNum:=0
pastNum:=arr[len(arr)-1]
fmt.Println(arr)
for i:=len(arr)-2;i>=0;i--{
	if maxNum<pastNum{
		maxNum=pastNum
	}
	pastNum=arr[i]
	arr[i]=maxNum
}

arr[len(arr)-1]=-1

return arr
}
