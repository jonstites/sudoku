package sudoku


import (
	"fmt"
	"testing"
)


//defaultCell := Cell{bitArray: 511}

func TestSetValue(t *testing.T) {
	for i := uint(1); i <= 9; i++ {
		for j := uint(1); j <= 9; j++ {
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

			myCell.setTriedValue(j)

			for k := uint(1); k <= 9; k++ {
				got := myCell.tried.valueTrue(k)
				expected := (j == k)
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

