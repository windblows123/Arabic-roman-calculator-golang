package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var roman_num1, roman_num2, roman_result string

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	splited_line := strings.Split(line, " ")

	if len(splited_line) != 3 {
		err := errors.New("Error: incorrect input")
		fmt.Println(err)
		os.Exit(0)
	}

	num1, error1 := strconv.Atoi(splited_line[0])
	operator := splited_line[1]
	num2, error2 := strconv.Atoi(splited_line[2])

	if error1 == nil && error2 == nil {

		if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
			err := errors.New("Error: one ore both of given numbers are out of 1 ... 10")
			fmt.Println(err)
			os.Exit(0)
		}

		result := calculate(num1, operator, num2)
		fmt.Println(result)
	} else {
		if error1 != nil && error2 != nil {
			roman_num1 = splited_line[0]
			roman_num2 = splited_line[2]
			int_num1 := roman_to_arabic(roman_num1)
			int_num2 := roman_to_arabic(roman_num2)

			if int_num1 > 10 || int_num1 < 1 || int_num2 > 10 || int_num2 < 1 {
				err := errors.New("Error: one ore both of given numbers are out of 1 ... 10")
				fmt.Println(err)
				os.Exit(0)
			}

			arabic_result := calculate(int_num1, operator, int_num2)
			roman_result = arabic_to_roman(arabic_result)
			fmt.Println(roman_result)
		} else {
			err := errors.New("Error: numbers are of different type")
			fmt.Println(err)
		}
	}
}

func roman_to_arabic(num string) int {
	var int_result, int_number int
	roman_arabic_map := map[string]string{
		"I":  "1",
		"IV": "4",
		"V":  "5",
		"IX": "9",
		"X":  "10",
		"XL": "40",
		"L":  "50",
		"XC": "90",
		"C":  "100",
	}
	num_length := len(num)
	var shift int
	for current_position := 0; current_position < num_length; {
		if current_position+1 < num_length {
			shift = 2
		} else {
			shift = 1
		}
		number := roman_arabic_map[num[current_position:current_position+shift]]
		if len(number) > 0 {
			int_number, _ = strconv.Atoi(number)
			int_result += int_number
			current_position += 2
		} else {
			number := roman_arabic_map[num[current_position:current_position+1]]
			int_number, _ = strconv.Atoi(number)
			int_result += int_number
			current_position += 1
		}
	}
	return int_result
}

func arabic_to_roman(num int) string {
	var roman_result string
	if num > 0 {
		arabic_roman_map := map[int]string{
			1:   "I",
			4:   "IV",
			5:   "V",
			9:   "IX",
			10:  "X",
			40:  "XL",
			50:  "L",
			90:  "XC",
			100: "C",
		}
		numbers := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

		for _, int_number := range numbers {
			for int_number <= num {
				roman_result += arabic_roman_map[int_number]
				num -= int_number
			}
		}
	} else {
		err := errors.New("Error: roman result is less than 1")
		fmt.Println(err)
	}

	return roman_result
}

func calculate(num1 int, operator string, num2 int) int {
	var result int
	switch operator {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2

	case "*":
		result = num1 * num2

	case "/":
		result = num1 / num2

	}
	return result
}
