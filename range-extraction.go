package solutions

/*
link:https://www.codewars.com/kata/51ba717bb08c1cd60f00002f/train/go
A format for expressing an ordered list of integers is to use a comma separated list of either

individual integers
or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'. The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"
Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.

Example:

solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
*/

import (
	"strconv"
)

func AbsDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func Solution(list []int) string {
	resultStr := ""
  skippedItems := []int{} 
  
	for i := 0; i < len(list)-1; i++ {
		diff := list[i+1] - list[i]
    
		if diff == 1 {
      skippedItems = append(skippedItems,list[i])
		} else {
      if len(skippedItems) >= 2 {
        resultStr += strconv.Itoa(skippedItems[0])
        resultStr += string('-')  
        resultStr += strconv.Itoa(list[i])
        resultStr += string(',')  
      } else {
        for _, num := range skippedItems {
          resultStr += strconv.Itoa(num)
          resultStr += string(',')  
        }
        resultStr += strconv.Itoa(list[i])
        resultStr += string(',')  
      }
      skippedItems = []int{}
		}
	}
  
  if len(skippedItems) >= 2 {
    resultStr += strconv.Itoa(skippedItems[0])
    resultStr += string('-')  
    resultStr += strconv.Itoa(list[len(list)-1])
  } else if len(skippedItems) == 1 {
      resultStr += strconv.Itoa(skippedItems[0])
      resultStr += string(',')  
      resultStr += strconv.Itoa(list[len(list)-1])
  } else {
      resultStr += strconv.Itoa(list[len(list)-1]) 
  }

	return resultStr 
}
