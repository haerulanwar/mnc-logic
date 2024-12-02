package main

import (
	"fmt"
	"strings"
)

func findMatchStrings(n int, strs []string) string {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// return found first set
			if strings.EqualFold(strs[i], strs[j]) {
				return fmt.Sprintf("%d %d", i+1, j+1)
			}
		}
	}
	return "false"
}

func main() {
	// test case
	fmt.Println(findMatchStrings(4, []string{"abcd", "acbd", "aaab", "acbd"}))
	fmt.Println(findMatchStrings(5, []string{"pisang", "goreng", "enak", "sekali", "rasanya"}))
	// contoh case 3 jawaban salah, karena set pertama yang ditemukan adalah sate pada posisi 2 dan 6
	fmt.Println(findMatchStrings(11, []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"}))
}
