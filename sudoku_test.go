package sudoku

import (
	"testing"
)

func TestCellPrint(t *testing.T) {
	myCell := Cell{3, map[int]bool{}, true}
	if !(myCell.String() == "3") {
		t.Error("Cell{3} should print as 3.")
	}
}

func TestCellOptions(t *testing.T) {
	myCell := Cell{3, map[int]bool{3: true, 5: true}, true}
	if !(myCell.valueOptions[3]) {
		t.Error("Cell should be allowed to have value 3")
	}

}

func TestCellNumOptions(t *testing.T) {
	myCell := Cell{3, map[int]bool{3: true, 5: true}, true}
	if (myCell.numValueOptions() != 2) {
		t.Error("Cell should have 2 options")
	}

}

func TestCellConstructor(t *testing.T) {
	myCell := newCell()
	for i := 1; i <= 9; i++ {
		if !(myCell.valueOptions[i]) {
			t.Error("Cell should have options 1-9")
		}
	}
}
