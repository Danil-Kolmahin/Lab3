package tools

import (
	"strconv"
	"strings"
)

func ASCIItoIntArr(ascii []uint8) ([]int, error){
    var result []int
    if len(ascii) == 0 {return result, nil}
	var withoutBrackets []uint8 = ascii[1:len(ascii)-1]
	convertedToString := string(withoutBrackets)
	StrNums := strings.Split(convertedToString, ",")
	for _, val := range StrNums {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
