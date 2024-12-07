package part01

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./day01/04.input")
	if err != nil {
		log.Fatal(err)
	}
	pairs := strings.Fields(string(f))
	firstList := make([]int, len(pairs)/2)
	secondList := make([]int, len(pairs)/2)

	for i := 0; i < len(pairs); i++ {
		num, err := strconv.Atoi(pairs[i])
		if err != nil {
			log.Fatal(err)
		}
		if i%2 == 1 {
			firstList = append(firstList, num)
		} else {
			secondList = append(secondList, num)
		}
	}
	fmt.Println(pairs)
	sort.Ints(firstList)
	sort.Ints(secondList)
	//fmt.Println(firstList)
	//fmt.Println(secondList)

	var sum float64

	for i := 0; i < len(firstList); i++ {
		sum = sum + math.Abs(float64(firstList[i]-secondList[i]))
	}
	fmt.Println("list")
	fmt.Println(sum)

}
