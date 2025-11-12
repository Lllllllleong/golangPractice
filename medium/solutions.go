package medium

/*
============================================================

============================================================
Time Complexity: O()
Space Complexity: O()
*/

/*
============================================================
3675. Minimum Operations to Transform String
============================================================
Time Complexity: O(n)
Space Complexity: O(1)
*/
func minOperations(s string) int {
	numOperations := 0
	for _, rune := range s {
		currentCost := int(26-(rune-'a')) % 26
		if currentCost > numOperations {
			numOperations = currentCost
		}
	}
	return numOperations
}
