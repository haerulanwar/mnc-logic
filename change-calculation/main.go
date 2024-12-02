package main

import "fmt"

func calculateChange(total, paid int) string {
	denoms := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	change := ((paid - total) / 100) * 100 // rounding down

	if change < 0 {
		return "False, kurang bayar\n"
	}

	// detail of change
	changeDetail := ""
	for _, p := range denoms {
		if change >= p {
			sumOf := change / p
			var kindOf string
			if p/500 > 1 {
				kindOf = "lembar"
			} else {
				kindOf = "koin"
			}
			change = change % p
			changeDetail += fmt.Sprintf("%d %s %d\n", sumOf, kindOf, p)
		}
	}

	return fmt.Sprintf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d\nPecahan uang:\n%s", paid-total, paid-total-((paid-total)%100), changeDetail)
}

func main() {
	// Test cases
	fmt.Println(calculateChange(700649, 800000))
	fmt.Println(calculateChange(657650, 600000))
	fmt.Println(calculateChange(575650, 580000))
}
