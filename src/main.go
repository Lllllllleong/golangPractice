package main

import "fmt"

func main() {
    fmt.Println("!!")

}


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