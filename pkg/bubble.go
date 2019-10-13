package bubble

// Sort does bubble sort.
func Sort(nums []int) []int {
	size := len(nums)
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < size-1; i++ {
			if nums[i] > nums[i+1] {
				sorted = false
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		}
	}

	return nums
}
