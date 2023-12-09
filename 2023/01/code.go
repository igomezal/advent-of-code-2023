package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	scanner := bufio.NewScanner(strings.NewReader(input))

	stringNumbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	count := 0

	if part2 {
		re := regexp.MustCompile("\\d|zero|one|two|three|four|five|six|seven|eight|nine")

		for scanner.Scan() {
			var numbers []string
			line := scanner.Text()
			for i := range line {
				numberInString := re.FindString(line[i:])
				if numberInString != "" {
					numbers = append(numbers, numberInString)
				}
			}
			firstNumber := "0"
			lastNumber := "0"
			firstNumber, ok := stringNumbers[numbers[0]]

			if !ok {
				firstNumber = numbers[0]
			}

			lastNumber, ok = stringNumbers[numbers[len(numbers)-1]]

			if !ok {
				lastNumber = numbers[len(numbers)-1]
			}

			numberFound, _ := strconv.Atoi(firstNumber + lastNumber)
			count = count + numberFound
		}
		return count
	}

	re := regexp.MustCompile("\\d")

	for scanner.Scan() {
		numbers := re.FindAllString(scanner.Text(), -1)

		firstNumber := numbers[0]
		lastNumber := numbers[len(numbers)-1]

		numberFound, _ := strconv.Atoi(firstNumber + lastNumber)
		count = count + numberFound
	}
	// solve part 1 here
	return count
}
