/*
Valid Parentheses
Easy Topics Company Tags
Hints

You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

    Every open bracket is closed by the same type of close bracket.
    Open brackets are closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

Return true if s is a valid string, and false otherwise.

Example 1:

Input: s = "[]"

Output: true

Example 2:

Input: s = "([{}])"

Output: true

Example 3:

Input: s = "[(])"

Output: false

Explanation: The brackets are not closed in the correct order.

Constraints:

    1 <= s.length <= 1000

*/

func isValid(s string) bool {
    lst:= []string{}
	hmap := map[string]string{
		")":"(",
		"]":"[",
		"}":"{",
	}
	fmt.Println(lst)
	if len(s)==1 {
		return false
	}
	if _, ok := hmap[string(s[0])]; ok {
    return false
	}

	lst=append(lst,string(s[0]))
	for i:=1;i<len(s);i++{
		actual_value:=string(s[i])
		if closing_value,ok:=hmap[actual_value];ok{
			if len(lst) == 0 || closing_value != lst[len(lst)-1]{
				return false
			}else{
				lst=lst[:len(lst)-1]
			}
		}else{
			lst=append(lst,string(s[i]))
		}
	}
	if len(lst)!=0 {
		return false
	}
	return true
}
