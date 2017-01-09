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
		{5, []uint {1, 3}, 2, 1},
		{30, []uint {2, 3, 4, 5}, 4, 2},
		{128, []uint {8}, 1, 8},
		{newBitArray(2, 5, 9), []uint {2, 5, 9}, 3, 2},
	}

	for _, bitTest := range bitsTest {

		var results = []testResult{
			{"number of values set", fmt.Sprint(bitTest.numSet), fmt.Sprint(bitTest.bits.numValuesTrue())},
			{"lowest value set", fmt.Sprint(bitTest.lowest), fmt.Sprint(bitTest.bits.lowestValue())},
		}

		for i := uint(1); i <= 9; i++ {
			test := fmt.Sprint(i, " is set")
			expected := fmt.Sprint(intInSlice(bitTest.set, i))
			got := fmt.Sprint(bitTest.bits.valueTrue(i))

			result := testResult{test, expected, got}
			results = append(results, result)
		}

		for _, result := range results {
			if result.expected != result.got {
				t.Errorf("Bits: %d Test: %q Got: %q Expected: %q.",
					bitTest.bits, result.test, result.got, result.expected)
			}
		}
	}
}

func TestBitArraySet(t *testing.T) {
	var bitsTest = []struct {
		bits bitarray
		testBit uint
		addBit bitarray
		subBit bitarray
	}{
		{newBitArray(2, 5, 9), 3, newBitArray(2, 3, 5, 9), newBitArray(2, 5, 9)},
		{newBitArray(1, 2, 5, 9), 1, newBitArray(1, 2, 5, 9), newBitArray(2, 5, 9)},
		{newBitArray(2, 5), 9, newBitArray(2, 5, 9), newBitArray(2, 5)},
		{newBitArray(5, 6, 7), 5, newBitArray(5, 6, 7), newBitArray(6, 7)},
		{newBitArray(), 1, newBitArray(1), newBitArray()},
		{newBitArray(1, 2, 3, 4, 5, 6, 7, 8, 9), 9, newBitArray(1, 2, 3, 4, 5, 6, 7, 8, 9), newBitArray(1, 2, 3, 4, 5, 6, 7, 8)},
	}

	for _, bitTest := range bitsTest {
		var results = []testResult{}

		result := testResult {
			fmt.Sprint("set bit %q to true", bitTest.testBit),
			fmt.Sprint(bitTest.addBit),
			fmt.Sprint(setBitTrue(bitTest.bits, bitTest.testBit)),
		}

		results = append(results, result)

		result = testResult {
			fmt.Sprint("set bit %q to false", bitTest.testBit),
			fmt.Sprint(bitTest.subBit),
			fmt.Sprint(setBitFalse(bitTest.bits, bitTest.testBit)),
		}

		results = append(results, result)

		for _, result := range results {
			if result.got != result.expected {
				t.Error("Bits: %q Test: %q Got: %q Expected: %q.",
					bitTest.bits, result.test, result.got, result.expected)
			}
		}
	}
}
