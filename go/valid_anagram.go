func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	 fstring :=make(map[string]int)
	 lstring :=make(map[string]int)
	
	
	si:= 0
	sf:= len(t)
	for i:=0; i<len(t);i++{
		fstring[string(s[i])]++
		lstring[string(t[i])]++
	} 
	
	return maps.Equal(fstring,lstring)
}
