// O(n^2)
// func twoSum(nums []int, target int) []int {
// 	var numAns []int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				numAns = []int{i, j}
// 				return numAns
// 			}
// 		}
// 	}

// 	return nil
// }

// O(n log n) use map
func twoSum(nums []int, target int) []int {
	var mapNum map[int]int = make(map[int]int)
	for currentIndex, num := range nums {
		pairNum := target - num
		pairNumIndex, isFound := mapNum[pairNum]
		if isFound == true {
			return []int{pairNumIndex, currentIndex}
		}
		mapNum[num] = currentIndex  // store num as key to the map, it will be used when the next index want to search a pairNum on the mapNum
        // for example, [2, 7, 11, 15] -> we are trying to get 18 (it will be 7+11)
        // for the first iteration, 2 doesn't have any pairNum to get 18 (18-2 = 16 doesn't exist on the mapNum), store 2 to the mapNum -> 2
        // then second iteration start, 18 - 7 = 11, check 11 on the mapNum (doesn't exist), store 7 to the mapNum -> 2, 7
        // third iteration start, 18 - 11 = 7, check 7 on the mapNum (exist), then return the pairNumIndex and the current index iteration
	}
	return nil
}