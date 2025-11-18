package easy

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
============================================================

============================================================
Time Complexity: O()
Space Complexity: O()
*/

/*
============================================================
3541. Find Most Frequent Vowel and Consonant
============================================================
Time Complexity: O(n)
Space Complexity: O(k) where k is the number of unique characters
*/
func maxFreqSum(s string) int {
	vowelSet := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	var vowelCount, consonantCount int
	runeMap := make(map[rune]int)
	for _, rune := range s {
		runeMap[rune]++
		if vowelSet[rune] && runeMap[rune] > vowelCount {
			vowelCount = runeMap[rune]
		}
		if !vowelSet[rune] && runeMap[rune] > consonantCount {
			consonantCount = runeMap[rune]
		}
	}
	return vowelCount + consonantCount
}

/*
============================================================
3731. Find Missing Elements
============================================================
Time Complexity: O(n log n)
Space Complexity: O(k) where k is the number of missing elements
*/
func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	output := []int{}
	counter := nums[0]
	for _, v := range nums {
		for counter < v {
			output = append(output, counter)
			counter++
		}
		counter++
	}
	return output
}

/*
============================================================
Two Sum
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func twoSum(nums []int, target int) []int {
	numberMap := make(map[int]int)
	for currentIndex, currentNumber := range nums {
		requiredNumber := target - currentNumber
		if secondIndex, exists := numberMap[requiredNumber]; exists {
			return []int{secondIndex, currentIndex}
		}
		numberMap[currentNumber] = currentIndex
	}
	return nil
}

/*
============================================================
FizzBuzz
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func fizzBuzz(n int) []string {
	outputSlice := []string{}
	for i := 1; i <= n; i++ {
		condition1 := (i % 3) == 0
		condition2 := (i % 5) == 0
		switch {
		case condition1 && condition2:
			outputSlice = append(outputSlice, "FizzBuzz")
		case condition1:
			outputSlice = append(outputSlice, "Fizz")
		case condition2:
			outputSlice = append(outputSlice, "Buzz")
		default:
			outputSlice = append(outputSlice, fmt.Sprintf("%d", i))
		}
	}
	return outputSlice
}

/*
============================================================
Score of a String
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func scoreOfString(s string) int {
	output := 0
	for i := 0; i < len(s)-1; i++ {
		difference := int(s[i]) - int(s[i+1])
		if difference < 0 {
			output += -difference
		} else {
			output += difference
		}
	}
	return output
}

/*
============================================================
Minimum Depth of Binary Tree
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	return 1 + min(left, right)
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}

/*
============================================================
Counting Words With a Given Prefix
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func prefixCount(words []string, pref string) int {
	output := 0
	prefixLength := len(pref)
	for _, word := range words {
		if (len(word) >= prefixLength) && (word[:prefixLength] == pref) {
			output++
		}
	}
	return output
}

/*
============================================================
Symmetric Tree
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(root1, root2 *TreeNode) bool {
	condition1 := root1 == nil
	condition2 := root2 == nil
	if condition1 && condition2 {
		return true
	}
	if (condition1 && !condition2) || (!condition1 && condition2) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return (isSymmetric2(root1.Right, root2.Left)) && (isSymmetric2(root1.Left, root2.Right))
}

/*
============================================================
Valid Palindrome II
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	if right == 0 {
		return true
	}
	flag := false
	for left <= right {
		if s[left] != s[right] {
			if flag {
				return false
			}
			flag = true
		}
		left++
		right--
	}
	return true
}

/*
============================================================
Make Two Arrays Equal by Reversing Subarrays
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	n := len(target)
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}

/*
============================================================
Divide Array Into Arrays With Max Difference
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	output := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		currentSlice := []int{}
		for j := 0; j < 3; j++ {
			currentSlice = append(currentSlice, nums[i+j])
		}
		if (currentSlice[2] - currentSlice[0]) > k {
			return [][]int{}
		}
		output = append(output, currentSlice)
	}
	return output
}
