package part01

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FindMul(data string) []string {
	re, _ := regexp.Compile(`mul\(\d*,\d*\)`)
	return re.FindAllString(data, -1)
}

func GetNumbers(data string) []string {
	re, _ := regexp.Compile(`\d*`)
	return re.FindAllString(data, -1)
}

func GetInput(filePath string) string {
	f, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(f)
}

func FilterNumbers(data []string) (numbers []int) {
	numbers = make([]int, 0)
	for _, str := range data {
		if strings.EqualFold(str, "") {
			continue
		}
		number, _ := strconv.Atoi(str)
		numbers = append(numbers, number)
	}
	return numbers
}

func SolvePuzzle(filePath string) {
	data := GetInput(filePath)
	muls := FindMul(data)
	sum := 0
	for _, mul := range muls {
		numbers := FilterNumbers(GetNumbers(mul))
		fmt.Println(numbers, len(numbers))
		sum = sum + numbers[0]*numbers[1]
	}
	fmt.Println(sum)

}
