package temparature

import (
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func GetTemp(str string) int64 {
	inputs := strings.Split(str," ")
	var min, max, result int64 = 5526, -273, 0
	for i := 0; i < len(inputs); i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t,_ := strconv.ParseInt(inputs[i],10,32)

		if t > 0 && t < min {
			min = t
		} else if t > max && t != 0 {
			max = t
		}
	}

	if len(inputs) > 1 {
		if max < 0 && min <= -max || max > 0 && min < max  {
			result = min
		} else {
			result = max
		}
	} else if str != "" {
		result, _ = strconv.ParseInt(str, 10, 32)
	}

	return result
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	// fmt.Println(result)// Write answer to stdout
}