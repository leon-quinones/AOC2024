package part02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func reduceNumber(x int) int {
	k := 0
	for x >= 10 {
		x = x / 10
		k++
	}
	if k == 0 {
		return k
	}
	reduced := x * int(math.Pow(float64(10), float64(k)))
	return reduced
}

func sortInput(filePath string) (dict map[int][]int, numberList []int) {
	dict = make(map[int][]int)
	numberList = make([]int, 0)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		numbers := strings.Fields(fileScanner.Text())
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		numberList = append(numberList, num1)
		key := reduceNumber(num2)
		dict[key] = append(dict[key], num2)
	}
	return dict, numberList
}

func count(x int, group []int) (times int) {
	times = 0
	for _, number := range group {
		if number == x {
			times++
		}
	}
	return times
}
func SolvePuzzle(filePath string) int {
	sum := 0
	dict, numbers := sortInput(filePath)
	for _, number := range numbers {
		sum = sum + number*count(number, dict[reduceNumber(number)])
	}
	return sum
}
