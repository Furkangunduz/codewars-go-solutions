package solutions

/*
link:https://www.codewars.com/kata/5659c6d896bc135c4c00021e/train/go

Write a function that takes a positive integer and returns the next smaller positive integer containing the same digits.

For example:

nextSmaller(21) == 12
nextSmaller(531) == 513
nextSmaller(2071) == 2017
Return -1 (for Haskell: return Nothing, for Rust: return None), when there is no smaller number that contains the same digits. Also return -1 when the next smaller number with the same digits would require the leading digit to be zero.

nextSmaller(9) == -1
nextSmaller(111) == -1
nextSmaller(135) == -1
nextSmaller(1027) == -1 // 0721 is out since we don't write numbers with leading zeros
*/
import (
	"sort"
	"strconv"
)

func NextSmaller(n int) int {
  if n < 10 {
    return -1
  } 
  
	strNum := strconv.Itoa(n)
  r := []rune(strNum)

	for i := len(r) - 1; i > 0; i-- {
    if r[i - 1] > r[i] {
      maxIndex := -1
      for j := len(r) - 1; j > i - 1; j-- {
        if r[j] < r[i - 1] {
          if maxIndex == -1 || r[j] > r[maxIndex] {
            maxIndex = j
          }
        }
      }
      
      if maxIndex == -1 {
        continue
      }

      r[i-1], r[maxIndex] = r[maxIndex], r[i-1]
      
      if r[0] == '0' {
         return -1
      }
      
      sort.Slice(r[i:], func(x,y int) bool {
        return r[i+x] > r[i+y]
      })
      
      result,err := strconv.Atoi(string(r)); if err != nil{
         panic(err)
      }
      return result
    }
	}
  return -1
}