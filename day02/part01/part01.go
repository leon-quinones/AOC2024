package part01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func InputReader(filePath string) (fileScanner *bufio.Scanner) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner = bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func ParseReport(reportData string) (report []int) {
	report = make([]int, 0)
	for _, level := range strings.Fields(reportData) {
		num, _ := strconv.Atoi(level)
		report = append(report, num)
	}
	return report
}

func CheckLevelsIncrements(report []int) bool {
	if len(report) == 1 {
		return true
	}
	diff := math.Abs(float64(report[0]) - float64(report[1]))
	if diff > 3 || diff == 0 {
		return false
	}
	return CheckLevelsIncrements(report[1:])
}

func CheckGradient(report []int) bool {
	levelChange := report[1] - report[0]
	for i := 1; i < len(report)-1; i++ {
		if levelChange*(report[i+1]-report[i]) < 0 {
			return false
		}
		levelChange = report[i+1] - report[i]
	}
	return true
}

func SolvePuzzle(filePath string) int {
	fileScanner := InputReader(filePath)
	safeReportNumbers := 0
	for fileScanner.Scan() {
		report := ParseReport(fileScanner.Text())
		if CheckLevelsIncrements(report) == false {
			continue
		}
		if CheckGradient(report) == false {
			continue
		}
		safeReportNumbers++
	}
	return safeReportNumbers
}
