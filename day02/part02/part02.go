package part02

import (
	"aoc2024/day02/part01"
	"fmt"
	"math"
)

func dampener(report []int) []int {

	index := getIndexBadLevel(report)
	if index < 0 {
		return report
	}
	newReport := make([]int, 0)
	newReport = append(newReport, report[:index]...)
	return append(newReport, report[index+1:]...)
}
func getIndexBadLevel(report []int) int {

	levelIndex := -1
	mag, grad := getLevelChanges(report)
	lmag := len(mag) != 0
	lgrad := len(grad) != 0
	if !lgrad && !lmag {
		return -1
	}
	if (lmag || lgrad) && !(lmag && lgrad) {
		if lmag {
			if mag[0] == 1 {
				return 0
			}
		}
		if lgrad {
			if grad[0] == 1 {
				return 0
			}
		}
	}
	if lgrad {
		levelIndex = grad[0]
	}
	if lmag {
		levelIndex = mag[0]
	}

	if lmag && lgrad {
		levelIndex = int(math.Min(float64(grad[0]), float64(mag[0])))
	}
	return levelIndex
}
func getLevelChanges(report []int) (mag, grad []int) {
	mag = make([]int, 0)
	posGradient := make([]int, 0)
	negGradient := make([]int, 0)
	//49 51 53 51 53 56 56
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 1 {
			negGradient = append(negGradient, i+1)
		} else {
			posGradient = append(posGradient, i+1)
		}
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			mag = append(mag, i+1)
		}
	}
	if len(posGradient) < len(negGradient) {
		grad = posGradient
	} else {
		grad = negGradient
	}
	return mag, grad
}

func SolvePuzzleTwo(filePath string) int {
	fileScanner := part01.InputReader(filePath)
	safeReportNumbers := 0
	for fileScanner.Scan() {

		report := part01.ParseReport(fileScanner.Text())

		dampedReport := dampener(report)
		if part01.CheckLevelsIncrements(dampedReport) == false {
			//fmt.Print(report)
			//fmt.Print(" UnSafe by levels")
			//fmt.Println(dampedReport)
			continue
		}
		if part01.CheckGradient(dampedReport) == false {
			//fmt.Print(report)
			//fmt.Print("UnSafe by gradient")
			//fmt.Println(dampedReport)
			continue
		}
		fmt.Print(report)
		fmt.Print(" Safe ")
		fmt.Println(dampedReport)
		safeReportNumbers++
	}
	return safeReportNumbers
}
