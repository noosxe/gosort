package selection

// Sort does selection sort.
func Sort(nums []int) []int {
	size := len(nums)

	for p := 0; p < size; p++ {
		min := p
		for i := p + 1; i < size; i++ {
			if nums[i] < nums[min] {
				min = i
			}
		}
		nums[p], nums[min] = nums[min], nums[p]
	}

	return nums
}
