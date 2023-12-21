package stuff

import (
	"strconv"
)

// Uppercase variable name means so it can be exported
var Name string = "Derek"

// Function name is uppercase so it can be exported
func IntArrToStrArr(intArr []int) []string {
	var strArr []string

	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}

	return strArr
}
