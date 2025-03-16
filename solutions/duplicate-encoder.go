package solutions

/*
link: https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go
The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.
*/
import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
  wordCounts := make(map[rune]int);
  
  word = strings.ToLower(word)
  
  for _, char := range word {
    _, exist := wordCounts[char]

    if exist {
      wordCounts[char] += 1;
    }else {
      wordCounts[char] = 1      
    }
  }
  
  fmt.Println(wordCounts)
  
  result := ""
  
  for _, char := range word {
    count, exist := wordCounts[char]
    
    if exist && count > 1 {
      result = result + ")"
    }else {
      result = result + "("
    }
  }
  
  
  return result
}