package sudoku


import (
	"fmt"
	"strings"
)


type grid struct {
	grid [][]cell
	triedCoords [][]int
}

// Create an empty 9x9 puzzle 
func newGrid() *grid {
	myGrid := new(grid)
	myGrid.grid = make([][]cell, 9)
	for i := 0; i < 9; i++ {
		myGrid.grid[i] = make([]cell, 9)
		for j := 0; j < 9; j++ {
			myGrid.grid[i][j] = *newCell()
		}
	}
	return myGrid
}

// Get a cell at particular coordinates  
func (myGrid *grid) getCell(rowNum int, colNum int) (*cell, error) {
	if rowNum < 0 || colNum < 0 || rowNum >= 9 || colNum >= 9 {
		return nil, fmt.Errorf("Tried to access out-of-bounds cell %d, %d",
			rowNum, colNum)
	}
	return &myGrid.grid[rowNum][colNum], nil
}

// Set cell value at given coordinates
func (myGrid *grid) setCellValue(rowNum int, colNum int, value uint) error {
	myCell, err := myGrid.getCell(rowNum, colNum)
	myCell.setValue(value, len(myGrid.triedCoords))
	return err
}

// Get cell value at given coordinates
func (myGrid *grid) getCellValue(rowNum int, colNum int) (uint, error) {
	myCell, err := myGrid.getCell(rowNum, colNum)
	return myCell.value, err
}

// Set the options the cell can be
func (myGrid *grid) setCellOptions(rowNum int, colNum int, options bitarray) error {
	myCell, err := myGrid.getCell(rowNum, colNum)
	myCell.setOptions(options)
	return err
}

// If a guess was wrong, undo all squares that depended on the guess
func (myGrid *grid) reset(guess int) {
	for _, row := range myGrid.grid {
		for _, cell := range row {
			cell.reset(guess)
		}
	}
	myGrid.updateAllOptions()
}

// Set the options for a cell
func (myGrid *grid) updateOptions(rowNum int, colNum int) {
	options := allTrue()

	// update within row
	for _, myCell := range myGrid.grid[rowNum] {
		if myCell.isKnown() {
			options = setBitFalse(options, myCell.value)
		}
	}

	// update within col
	for _, row := range myGrid.grid {
		myCell := row[colNum]
		if myCell.isKnown() {
			options = setBitFalse(options, myCell.value)
		}
	}

	// update within block
	blockRow := rowNum - (rowNum % 3)
	blockCol := colNum - (colNum % 3)
	for rowIndex := blockRow; rowIndex < blockRow + 3; rowIndex += 1 {
		for colIndex := blockCol; colIndex < blockCol + 3; colIndex += 1 {
			myCell, _ := myGrid.getCell(rowIndex, colIndex)
			if myCell.isKnown() {
				options = setBitFalse(options, myCell.value)
			}
		}
	}

	// update this cell
	myCell, _ := myGrid.getCell(rowNum, colNum)
	myCell.setOptions(options)
}

// Update all the options in the grid
func (myGrid *grid) updateAllOptions() {
	for i:= 0; i < 9; i++ {
		for j:= 0; j < 9; j++ {
			myGrid.updateOptions(i, j)
		}
	}
}

// Check if the grid is totally filled
func (myGrid *grid) isComplete() bool {
	for _, row := range myGrid.grid {
		for _, cell := range row {
			if !(cell.isKnown()) {
				return false
			}
		}
	}
	return true
}

// Overrwrite string option
func (myGrid *grid) String() string {
	var box []string
	for _, row := range myGrid.grid {
		var rowValues []string
		for _, col := range row {
			rowValues = append(rowValues, col.String())
		}
		box = append(box, strings.Join(rowValues, ""))
	}
	return strings.Join(box, "\n")
}
