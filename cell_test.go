package sudoku


import (
	"fmt"
	"testing"
)


//defaultCell := Cell{bitArray: 511}

func TestSetValue(t *testing.T) {
	for i := uint(1); i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			myCell := newCell()
			myCell.setValue(i, j)
			if myCell.value != i {
				t.Errorf("Expected value %d but got %d.",
				i, myCell.value)
			}

			if myCell.guess != j {
				t.Errorf("Expected value %d but got %d.",
				i, myCell.value)
			}

			uint_j := uint(j)
			myCell.setTriedValue(uint_j)

			for k := uint(1); k <= 9; k++ {
				got := myCell.tried.valueTrue(k)
				expected := (uint_j == k)
				if got != expected {
					t.Errorf("Value %d was expected to be %t but got %t in %d.",
						j, expected, got, myCell.tried)
				}
			}
		}
	}
}
	
func TestConstructor(t *testing.T) {
	myCell := newCell()
	for i := uint(1); i <= 9; i++ {
		expected := true
		got := myCell.options.valueTrue(i)
		if got != expected {
			t.Errorf("Expected %q set to %q but got %q.",
				i, expected, got)
		}
	}
}

func TestCellString(t *testing.T) {
	for i := uint(1); i <= 9; i++ {
		myCell := cell{value: i}
		if myCell.String() != fmt.Sprint(i) {
			t.Errorf("Cell %s should print as %d.",
				myCell.String(), i)
		}
	}
}

func TestSetOptions(t *testing.T) {
	var testCells = []struct {
		cell cell
		value uint
		hasOption bool
	}{
		{cell{tried: 0}, 5, true},
		{cell{tried: 1}, 1, false},
		{cell{tried: 5}, 3, false},
		{cell{tried: 30}, 7, true},
		{cell{tried: 128}, 8, false},
	}

	for _, testCell := range testCells {
		testCell.cell.setOptions(511)
		expected := testCell.hasOption
		got := testCell.cell.options.valueTrue(testCell.value)
		if got != expected {
			t.Errorf("Expected %q set to %b but got %q.",
				testCell.value, expected, got)
		}
	}
}

func TestIsKnown(t *testing.T) {
	var testValues = []struct {
		value uint
		known bool
	}{
		{0, false},
		{1, true},
		{5, true},
		{9, true},
	}

	for _, testValue := range testValues {
		myCell := newCell()
		myCell.setValue(testValue.value, 0)
		got := myCell.isKnown()
		expected := testValue.known
		if got != expected {
			t.Errorf("Set value to %q, expected isKnown %t but got %t",
				testValue.value, expected, got)
		}
	}
}
