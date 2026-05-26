/*
Container With Most Water
Medium Topics Company Tags
Hints

You are given an integer array heights where heights[i] represents the height of the ithith bar.

You may choose any two bars to form a container. Return the maximum amount of water a container can store.

Example 1:

Input: height = [1,7,2,5,4,7,3,6]

Output: 36

Example 2:

Input: height = [2,2,2]

Output: 4
*/


func maxArea(heights []int) int {
    l:=0
    r:=len(heights)-1
    area:=0
    for l<r{
        altura:=min(heights[l],heights[r])
        ancho:=r-l
        tempa:= altura* ancho
        if area<tempa{
            area=tempa
        }
        if altura==heights[l]{
            l++
        }else{
            r--
        }
    }
    return area
}
