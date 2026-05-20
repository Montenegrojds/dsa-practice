
func isPalindrome(s string) bool {
 pi:=0
 pf:=len(s)-1
 s=strings.ToLower(s)
 

for pi<pf{

	if !unicode.IsLetter(rune(s[pi])) && !unicode.IsDigit(rune(s[pi])){
		pi++
		continue
	}
	if !unicode.IsLetter(rune(s[pf]))&& !unicode.IsDigit(rune(s[pf])){
		pf--
		continue
	}

	if s[pi]!=s[pf]{
		return false
	}
	pi=pi+1
	pf=pf-1
}
 return true
}
