package sudoku

import (
	"fmt"
	"testing"
)

func intInSlice(values []uint, target uint) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

type testResult struct {
	test string
	expected string
	got string
}

func TestBitArray(t *testing.T) {
	var bitsTest = []struct {
		bits bitarray
		set []uint
		numSet int
		lowest uint
	}{
		{0, []uint {}, 0, 0},
		{1, []uint {1}, 1, 1},
		{5, []uint {1, 5}, 2, 1},
		{30, []uint {2, 3, 4}, 3, 2},
		{128, []uint {7}, 1, 7},
		{newBitArray(2, 5, 9), []uint {2, 5, 9}, 3, 2},
	}

	for _, bitTest := range bitsTest {

		var results = []testResult {
			{"number of values set", fmt.Sprint(bitTest.numSet), fmt.Sprint(bitTest.bits.numValuesSet())},
			{"lowest value set", fmt.Sprint(bitTest.numSet), fmt.Sprint(bitTest.bits.lowestValue())},
		}

		for i := uint(1); i <= 9; i++ {
			test := fmt.Sprint("%q is set", i)
			expected := fmt.Sprint(intInSlice(bitTest.set, i))
			got := fmt.Sprint(bitTest.bits.valueSet(i))

			result := testResult{test, expected, got}
			results = append(results, result)
		}

		for _, result := range results {
			t.Error("Bits: %q Test: %q Got: %q Expected: %q.",
				bitTest.bits, result.test, result.got, result.expected)
		}
	}
}
