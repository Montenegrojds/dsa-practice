/*
Valid Palindrome II
Easy Topics Company Tags

You are given a string s, return true if the s can be a palindrome after deleting at most one character from it.

A palindrome is a string that reads the same forward and backward.

Note: Alphanumeric characters consist of letters (A-Z, a-z) and numbers (0-9).

Example 1:

Input: s = "aca"

Output: true

Explanation: "aca" is already a palindrome.

Example 2:

Input: s = "abbadc"

Output: false

Explanation: "abbadc" is not a palindrome and can't be made a palindrome after deleting at most one character.

Example 3:

Input: s = "abbda"

Output: true

Explanation: "We can delete the character 'd'.

Constraints:

    1 <= s.length <= 100,000
    s is made up of only lowercase English letters.

*/

func valPalin(s string, l int,r int) bool{
	for  l<r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--	
	}
	return true
}

func validPalindrome(s string) bool {
	l:=0
	r:=len(s)-1
	for l<r{
		if s[l]!=s[r]{
			return valPalin(s,l+1,r)|| valPalin(s,l,r-1)
		}
		l++
		r--
	}
	return true
}
