package temparature

import (
	"fmt"
	"strconv"
	"testing"
)

type Input struct {
	input  string
	output int64
}

func TestGetTemp(t *testing.T) {
	inputs := map[string]Input{
		"Simple test case": {
			input:  "1 -2 -8 4 5",
			output: 1,
		},
		"-273 alone": {
			input:  "-273",
			output: -273,
		},
		"5526 alone": {
			input:  "5526",
			output: 5526,
		},
		"Only negative numbers": {
			input:  "-12 -5 -137",
			output: -5,
		},
		"Choose the right temperature": {
			input:  "42 -5 12 21 5 24",
			output: 5,
		},
		"Choose the right temperature 2": {
			input:  "42 5 12 21 -5 24",
			output: 5,
		},
		"Complex test case": {
			input:  "-5 -4 -2 12 -40 4 2 18 11 5",
			output: 2,
		},
		"No temperature": {
			input:  "",
			output: 0,
		},
	}

	for name, str := range inputs {
		fmt.Println(name)
		testTemp(t, str.input, str.output)
	}
}

func testTemp(t *testing.T, input string, expectedOutput int64) {
	result := GetTemp(input)
	if result != expectedOutput {
		t.Error("Expected output "+strconv.FormatInt(expectedOutput, 10)+" but output is ", result)
	}
	//fmt.Println("Output: " + strconv.Itoa(int(result)))
}
