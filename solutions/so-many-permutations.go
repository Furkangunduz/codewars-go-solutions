package solutions

/*
link:https://www.codewars.com/kata/5254ca2719453dcc0b00027d/train/go
In this kata, your task is to create all permutations of a non-empty input string and remove duplicates, if present.

Create as many "shufflings" as you can!

Examples:

With input 'a':
Your function should return: ['a']

With input 'ab':
Your function should return ['ab', 'ba']

With input 'abc':
Your function should return ['abc','acb','bac','bca','cab','cba']

With input 'aabb':
Your function should return ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa']
*/


func Permutations(s string) []string {
	seen := make(map[string]bool)              
	return permute([]rune(s), 0, seen)         
}

func permute(chars []rune, index int, seen map[string]bool) []string {
	var result []string

	if index == len(chars)-1 {
		perm := string(chars)
		if !seen[perm] {            
			result = append(result, perm)
			seen[perm] = true       
		}
		return result
	}

	for i := index; i < len(chars); i++ {
		chars[index], chars[i] = chars[i], chars[index] 

		subPermutations := permute(chars, index+1, seen)
		result = append(result, subPermutations...)     

		chars[index], chars[i] = chars[i], chars[index]
	}

	return result
}
