package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	k := []int{}
	sc.Scan()
	text := sc.Text()
	for _, v := range strings.Fields(text) {
		num, _ := strconv.Atoi(v)
		k = append(k, num)
	}
	fmt.Println(solution(k))
}

func solution(k []int) int {
	baseWeight := findWeight(k)
	// fmt.Println(baseWeight)
	lengths := []int{}
	for i := 0; i < len(k)-1; i++ {
		for j := i + 1; j < len(k); j++ {
			subArray := k[i : j+1]
			weight := findWeight(subArray)
			if weight == baseWeight {
				lengths = append(lengths, len(subArray))
				// fmt.Println(subArray, len(subArray), weight)
			}
		}
	}
	minLength := lengths[0]
	for _, v := range lengths[1:] {
		if v < minLength {
			minLength = v
		}
	}
	return minLength
}

func findWeight(k []int) int {
	weight := 0
	noOfOccurrences := make(map[int]int)
	for _, v := range k {
		noOfOccurrences[v]++
	}

	for _, numberCount := range noOfOccurrences {
		if numberCount > weight {
			weight = numberCount
		}
	}
	return weight
}
