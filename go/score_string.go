/*
Score of a String
Easy Topics Company Tags

You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.

Return the score of s.

Example 1:

Input: s = "code"

Output: 24

Explanation: The ASCII values of the characters in the given string are: 'c' = 99, 'o' = 111, 'd' = 100, and 'e' = 101. The score of s will be: |111 - 99| + |100 - 111| + |101 - 100|.

Example 2:

Input: s = "neetcode"

Output: 65

Constraints:

    2 <= s.length <= 100
    s is made up of lowercase English letters.
*/

func abs(num int) int {
	if num<0{
		return -num
	}
	return num
}

func scoreOfString(s string) int {
counter:=0

for i:=0;i<len(s)-1;i++{
	counter=counter+ abs(int (s[i])- int (s[i+1]))
}
return counter
}
