package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func pt1() {
	data, err := ioutil.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	// a slice of ints
	slice_first := []int{}
	slice_second := []int{}

	for _, line := range lines {
		// split by any whitespace
		pairs := strings.Fields(line)
		if len(pairs) == 0 {
			continue
		}
		first_int, err := strconv.Atoi(pairs[0])
		if err != nil {
			panic(err)
		}
		second_int, err := strconv.Atoi(pairs[1])
		if err != nil {
			panic(err)
		}

		slice_first = append(slice_first, first_int)
		slice_second = append(slice_second, second_int)
	}

	// sort
	sort.Ints(slice_first)
	sort.Ints(slice_second)

	sum := 0
	for i, first := range slice_first {
		second := slice_second[i]
		absolute_diff_int := int(math.Abs(float64(first - second)))
		sum += absolute_diff_int
	}

	fmt.Println(sum)
}

func count_occurances(slice []int, target int) int {
  count := 0
  for _, num := range slice {
    if num == target {
      count++
    }
  }
  return count
}

func pt2() {
	data, err := ioutil.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	// a slice of ints
	slice_first := []int{}
	slice_second := []int{}

	for _, line := range lines {
		// split by any whitespace
		pairs := strings.Fields(line)
		if len(pairs) == 0 {
			continue
		}
		first_int, err := strconv.Atoi(pairs[0])
		if err != nil {
			panic(err)
		}
		second_int, err := strconv.Atoi(pairs[1])
		if err != nil {
			panic(err)
		}

		slice_first = append(slice_first, first_int)
		slice_second = append(slice_second, second_int)
	}
  total_score := 0

  for _, first := range slice_first {
    count := count_occurances(slice_second, first)
    score:= count * first
    total_score += score
  }

  fmt.Println(total_score)
}

func main() {
	pt2()
}
