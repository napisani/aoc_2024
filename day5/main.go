package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func parseRuleSet(rulesStr string) map[Rule]bool {
	rules := make(map[Rule]bool)
	for _, line := range strings.Split(rulesStr, "\n") {
		rule := Rule{}
		fmt.Sscanf(line, "%d|%d", &rule.before, &rule.after)
		rules[rule] = true
	}
	return rules
}

func inCorrectOrder(rules map[Rule]bool, printOrders []int, idx int) bool {
	num := printOrders[idx]
	for printInstructionIdx, printInstruction := range printOrders {
		if printInstructionIdx == idx {
			continue
		}

		_, mustBeBefore := rules[Rule{before: num, after: printInstruction}]
		_, mustBeAfter := rules[Rule{before: printInstruction, after: num}]
		isBefore := idx < printInstructionIdx
		isAfter := idx > printInstructionIdx
		if mustBeBefore && !isBefore {
			return false
		}
		if mustBeAfter && !isAfter {
			return false
		}

	}

	return true
}

func allCorrectOrder(rules map[Rule]bool, printOrders []int) bool {
	for i := range printOrders {
		if !inCorrectOrder(rules, printOrders, i) {
			return false
		}
	}
	return true
}

func getCorrectOrder(rules map[Rule]bool, printOrders []int) []int {
	if allCorrectOrder(rules, printOrders) {
		return printOrders
	}

	for {
		for idx, num := range printOrders {
			for printInstructionIdx, printInstruction := range printOrders {
				if printInstructionIdx == idx {
					continue
				}

				_, mustBeBefore := rules[Rule{before: num, after: printInstruction}]
				_, mustBeAfter := rules[Rule{before: printInstruction, after: num}]
				isBefore := idx < printInstructionIdx
				isAfter := idx > printInstructionIdx
				if mustBeBefore && !isBefore {
					temp := printOrders[idx]
					printOrders[idx] = printOrders[printInstructionIdx]
					printOrders[printInstructionIdx] = temp
					if allCorrectOrder(rules, printOrders) {
						return printOrders
					}

				}
				if mustBeAfter && !isAfter {
					temp := printOrders[idx]
					printOrders[idx] = printOrders[printInstructionIdx]
					printOrders[printInstructionIdx] = temp
					if allCorrectOrder(rules, printOrders) {
						return printOrders
					}
				}
			}
		}
	}
}

func pt2() {
	data, err := ioutil.ReadFile("input/day5.txt")
	if err != nil {
		panic(err)
	}

	splitData := strings.Split(string(data), "\n\n")
	rulesStr := splitData[0]
	instructsStr := splitData[1]

	fmt.Println(instructsStr)
	rules := parseRuleSet(rulesStr)

	for rule := range rules {
		fmt.Println(rule)
	}

	splitInstructs := strings.Split(instructsStr, "\n")

	sum := 0

	for _, instruct := range splitInstructs {
		fmt.Println("instruct: ", instruct)
		if instruct == "" {
			continue
		}

		printOrdersStr := strings.Split(instruct, ",")
		fmt.Println("printOrdersStr: ", printOrdersStr)
		printOrders := make([]int, len(printOrdersStr))
		for i, printOrderStr := range printOrdersStr {
			printOrders[i], err = strconv.Atoi(printOrderStr)
			if err != nil {
				panic(err)
			}
		}

		allCorrect := allCorrectOrder(rules, printOrders)
		fmt.Println("allCorrect: ", allCorrect)
		if !allCorrect {
			corrected := getCorrectOrder(rules, printOrders)
			fmt.Println("corrected: ", corrected)
			if corrected != nil {
				mid := (len(corrected) / 2)
				fmt.Println("mid: ", mid, corrected)
				sum += corrected[mid]
			}
		} else {
			fmt.Println("-----all correct", printOrders)

		}

		fmt.Println("----")
	}
	fmt.Println(sum)

}

func pt1() {
	data, err := ioutil.ReadFile("input/day5.txt")
	if err != nil {
		panic(err)
	}

	splitData := strings.Split(string(data), "\n\n")
	rulesStr := splitData[0]
	instructsStr := splitData[1]

	fmt.Println(instructsStr)
	rules := parseRuleSet(rulesStr)

	for rule := range rules {
		fmt.Println(rule)
	}

	splitInstructs := strings.Split(instructsStr, "\n")

	sum := 0

	for _, instruct := range splitInstructs {
		if instruct == "" {
			continue
		}

		printOrdersStr := strings.Split(instruct, ",")
		printOrders := make([]int, len(printOrdersStr))
		for i, printOrderStr := range printOrdersStr {

			printOrders[i], err = strconv.Atoi(printOrderStr)
			if err != nil {
				panic(err)
			}
		}

		allCorrect := allCorrectOrder(rules, printOrders)
		if allCorrect {
			mid := (len(printOrders) / 2)
			fmt.Println("mid: ", mid, printOrders)
			sum += printOrders[mid]
		}

		fmt.Println("----")
	}
	fmt.Println(sum)

}
func main() {
	pt2()
}

