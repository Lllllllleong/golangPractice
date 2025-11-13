package medium

/*
============================================================

============================================================
Time Complexity: O()
Space Complexity: O()
*/

/*
============================================================
3597. Partition String
============================================================
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
func partitionString(s string) []string {
	runes := []rune(s)
	n, i := len(runes), 0
	seenSet := make(map[string]struct{})
	output := []string{}
	for i < n {
		j := i + 1
		for j <= n {
			currentSubstring := string(runes[i:j])
			if _, ok := seenSet[currentSubstring]; ok {
				j++
			} else {
				seenSet[currentSubstring] = struct{}{}
				output = append(output, currentSubstring)
				break
			}
		}
		i = j
	}
	return output
}

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
