package main

import (
	"fmt"
	"math"
	"time"
)

func leaveValidation(collectiveLeave int, joinDate, leaveDate time.Time, leaveDuration int) (bool, string) {
	// joinDate after period time
	passPeriodTime := joinDate.Add(time.Hour * 24 * 180)
	isPeriodTime := leaveDate.Sub(passPeriodTime).Hours() / 24
	if isPeriodTime <= 0 {
		return false, "\nAlasan: Karena belum 180 hari sejak tanggal join karyawan\n"
	}

	// end of year
	dateEndOfYear := time.Date(leaveDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	daysSinceJoin := dateEndOfYear.Sub(passPeriodTime).Hours() / 24

	// personal leave
	personalLeave := 14
	totalDaysInYear := 365
	personalLeave = int(math.Floor((float64(daysSinceJoin) / float64(totalDaysInYear)) * float64(personalLeave-collectiveLeave)))

	if leaveDuration > personalLeave {
		return false, fmt.Sprintf("\nAlasan: Karena hanya boleh mengambil %d hari cuti\n", personalLeave)
	}

	// Memeriksa durasi cuti berturut-turut
	if leaveDuration > 3 {
		return false, "\nAlasan: Karena hanya boleh mengambil 3 hari cuti berturut-turut\n"
	}

	return true, "\n"
}

func main() {
	// Test cases
	joinDate, _ := time.Parse("2006-01-02", "2021-05-01")
	leaveDate, _ := time.Parse("2006-01-02", "2021-07-05")
	// Test Case 1
	fmt.Println(leaveValidation(7, joinDate, leaveDate, 1)) // Output: false, Karena belum 180 hari sejak tanggal join karyawan

	// Test Case 2
	leaveDate, _ = time.Parse("2006-01-02", "2021-11-05")
	fmt.Println(leaveValidation(7, joinDate, leaveDate, 3)) // Output: false, Karena hanya boleh mengambil 1 hari cuti

	// Test Case 3
	joinDate, _ = time.Parse("2006-01-02", "2021-01-05")
	leaveDate, _ = time.Parse("2006-01-02", "2021-12-18")
	fmt.Println(leaveValidation(7, joinDate, leaveDate, 1)) // Output: true
	fmt.Println(leaveValidation(7, joinDate, leaveDate, 3)) // Output: true
}
