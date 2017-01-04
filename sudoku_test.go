package sudoku

import (
	"testing"
)

func TestCellPrint(t *testing.T) {
	myCell := Cell{3}
	if !(myCell.String() == "3") {
		t.Error("Cell{3} should print as 3.")
	}
}
