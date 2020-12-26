package tools

import (
	"strconv"
	"strings"
)

func ASCIItoIntArr(ascii []uint8) ([]int, error) {
	var result []int
	if len(ascii) == 0 {
		return result, nil
	}
	var withoutBrackets []uint8 = ascii[1 : len(ascii)-1] //to {51, 44, 53} after 51, 44, 53
	convertedToString := string(withoutBrackets)          //to 51, 44, 53 after 3,6
	StrNums := strings.Split(convertedToString, ",")      //to 3,6 after ["3", "6"]
	for _, val := range StrNums {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result = append(result, i) // [3, 6]
	}
	return result, nil
}
