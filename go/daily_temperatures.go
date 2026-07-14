/*
Daily Temperatures
Medium Topics Company Tags
Hints

You are given an array of integers temperatures where temperatures[i] represents the daily temperatures on the ith day.

Return an array result where result[i] is the number of days after the ith day before a warmer temperature appears on a future day. If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

Example 1:

Input: temperatures = [30,38,30,36,35,40,28]

Output: [1,4,1,2,1,0,0]

Example 2:

Input: temperatures = [22,21,20]

Output: [0,0,0]

Constraints:

    1 <= temperatures.length <= 1000.
    1 <= temperatures[i] <= 100

*/


type Data struct {
    Index int
    Value int
}

func dailyTemperatures(temperatures []int) []int {
	endstack:=0
	rpta:=make([]int,len(temperatures))
	stack:=[]Data{}
	for i:=0;i<len(temperatures)-1;i++{
		stack=append(stack,Data{i,temperatures[i]})
		endstack=len(stack)-1
		if stack[endstack].Value<temperatures[i+1]{
			for endstack>=0 {
			if	stack[endstack].Value<temperatures[i+1]{
			rpta[stack[endstack].Index]=i+1-stack[endstack].Index
			stack=stack[:endstack]
			}else{
				break
			}
			endstack--
		}}
	}
	return rpta
}
