package tools

import (
	"strconv"
	"strings"
)

func ASCIItoIntArr(ascii []uint8) ([]int, error){
	var withoutBrackets []uint8 = ascii[1:len(ascii)-1]
	convertedToString := string(withoutBrackets)
	StrNums := strings.Split(convertedToString, ",")
	var result []int
	for _, val := range StrNums {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
