package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func max_delta(numbers []int) (int, bool) {
	max := 0
	var direction = 0

	for i, num := range numbers {
		if i == len(numbers)-1 {
			break
		}
		next_num := numbers[i+1]
		delta := next_num - num
		if delta == 0 {
			return 0, false
		}
		if i == 0 {
			if delta > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else if delta > 0 && direction == -1 {
			return 0, false
		} else if delta < 0 && direction == 1 {
			return 0, false
		}

		delta = int(math.Abs(float64(delta)))
		if delta > max {
			max = delta
		}

	}
	return max, true
}

func is_valid(nums []int) bool {
	c, ok := max_delta(nums)
	if c <= 3 && ok {
		return true
	}
	for i, _ := range nums {
		new_nums := make([]int, len(nums)-1)
		copy(new_nums, nums[:i])
		copy(new_nums[i:], nums[i+1:])
		c, ok := max_delta(new_nums)
		if c <= 3 && ok {
			return true
		}
	}
	return false

}

func pt1() {
	data, err := ioutil.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	// a slice of ints

	total := 0
	for _, line := range lines {
		// split by any whitespace
		strs := strings.Fields(line)
		if len(strs) == 0 {
			continue
		}
		nums := make([]int, len(strs))
		for i, s := range strs {
			nums[i], err = strconv.Atoi(s)
		}

		if is_valid(nums) {
			total += 1
		}
	}
	fmt.Println(total)
}
func main() {
	pt1()
}
