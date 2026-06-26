/*
Can Place Flowers
Easy Topics Company Tags

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

You are given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1

Output: true

Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2

Output: false

Constraints:

    1 <= flowerbed.length <= 20,000
    flowerbed[i] is 0 or 1.
    There are no two adjacent flowers in flowerbed.
    0 <= n <= flowerbed.length
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
    counter:=0
    lenValue:=len(flowerbed)-1
    if len(flowerbed)>1{
        if flowerbed[0]==0 && flowerbed[1]==0{
            flowerbed[0]=1
            counter++
        }   
        if flowerbed[lenValue]== 0 && flowerbed[lenValue-1]==0{
            flowerbed[lenValue]=1
            counter++
        }
    }else{
        if lenValue == 0 && flowerbed[0]==0{
            flowerbed[0]=1
            counter++
        }
    }
    i:=1    
    for i<lenValue{
        if flowerbed[i-1]==0 && flowerbed[i]==0 && flowerbed[i+1]==0{
            flowerbed[i]=1
            counter++
            }
        i++
    }
    return counter>=n

}