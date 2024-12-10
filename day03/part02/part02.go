package part02

import (
	"aoc2024/day03/part01"
	"fmt"
	"regexp"
)

func filterDoBlocks(data string) string {
	re, _ := regexp.Compile(`.?[-_\s\(\)]?don't\(\).*?[\s]?.*?do\(\)`)
	return re.ReplaceAllString(data, "")
}

func SolvePuzzleTwo(filePath string) {
	data := part01.GetInput(filePath)
	filterData := filterDoBlocks(data)
	muls := part01.FindMul(filterData)
	sum := 0
	for _, mul := range muls {
		numbers := part01.FilterNumbers(part01.GetNumbers(mul))
		fmt.Println(numbers, len(numbers))
		sum = sum + numbers[0]*numbers[1]
	}
	fmt.Println(sum)

}
