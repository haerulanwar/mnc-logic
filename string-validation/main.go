package main

import (
	"fmt"
)

func validate(str string) bool {
	stack := []rune{}
	// char mapping, close should follow open
	opening := map[rune]rune{'<': '>', '{': '}', '[': ']'}
	closing := map[rune]rune{'>': '<', '}': '{', ']': '['}

	for _, char := range str {
		if _, isOpen := opening[char]; isOpen {
			// push opening char to stack
			stack = append(stack, char)
		} else if _, isClose := closing[char]; isClose {
			if len(stack) == 0 {
				return false
			}
			lastOpen := stack[len(stack)-1]
			if closing[char] != lastOpen {
				return false
			}
			stack = stack[:len(stack)-1] // Pop stack
		}
	}
	// unfinish opening
	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(validate("{{[<>[{{}}]]}}"))                   // Output: true                                                                                                                                        // Output: true
	fmt.Println(validate("{<{[[{{[]<{{[{[]<>}]}}<>>}}]]}>}")) // Output: true
	fmt.Println(validate("{<{[[{{[]<{[{[]<>}]}}<>>}}]]}>}"))  // Output: false
	fmt.Println(validate("[{}<[>]"))                          // Output: false
}
