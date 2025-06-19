package main

import (
    "fmt"
    "math"
)

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


func scoreOfString(s string) int {
    output := 0
    for i := 0; i < len(s) - 1; i++ {
        difference := int(s[i]) - int(s[i+1])
        if difference < 0 {
            output += -difference
        } else {
            output += difference
        }
    }
    return output
}