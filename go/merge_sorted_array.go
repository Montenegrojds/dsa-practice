func merge(nums1 []int, m int, nums2 []int, n int) {
	freespace:=len(nums1)-1
	m--
	n--
	for freespace>=0 {
		if   m>=0 && n>=0 &&  nums1[m]>=nums2[n] {
			nums1[freespace]=nums1[m]
			m--
		}else if n>=0{
			nums1[freespace]=nums2[n]
			n--
		}
		freespace--
	}

}
