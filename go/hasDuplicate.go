func hasDuplicate(nums []int) bool {
	myMap := make(map[int]int)

	for index := range nums {
		fmt.Println(nums[index])
		_,exist:=myMap[nums[index]]
		if exist{
			return true
		}else{
            myMap[nums[index]]=1
		}
	}
	return false
}
