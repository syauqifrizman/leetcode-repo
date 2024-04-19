func twoSum(nums []int, target int) []int {
	var numAns []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				numAns = []int{i, j}
				return numAns
			}
		}
	}

	return nil
}