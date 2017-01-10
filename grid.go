package sudoku


import (
	"fmt"
	"strings"
)


type grid struct {
	grid [][]cell
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
func (myGrid *grid) setCellValue(rowNum int, colNum int, value uint, numTried int) error {
	myCell, err := myGrid.getCell(rowNum, colNum)
	myCell.setValue(value, numTried)
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
	for i, row := range myGrid.grid {
		for j, _ := range row {
			myCell, _ := myGrid.getCell(i, j)
			myCell.reset(guess)
		}
	}
	myGrid.updateAllOptions()
}

// Set the options for a cell
func (myGrid *grid) updateOptions(rowNum int, colNum int) {
	myCell, _ := myGrid.getCell(rowNum, colNum)
	if myCell.isKnown() {
		return
	}
	
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
	myCell, _ = myGrid.getCell(rowNum, colNum)
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

// Check that the sudoku meets requirements
func (myGrid *grid) isValid() bool {
	for i := 0; i <= 8; i++ {
		var rowSum uint
		var colSum uint
		var blockSum uint
		for j := 0; j <= 8; j++ {
			row, _ := myGrid.getCellValue(i, j)
			rowSum += row
			col, _ := myGrid.getCellValue(j, i)
			colSum += col
			blockRow := (j / 3) + 3 * (i % 3)
			blockCol := (j % 3) + 3 * (i / 3)
			block, _ := myGrid.getCellValue(blockRow, blockCol)
			blockSum += block
		}
		if !((blockSum == rowSum) && (rowSum == colSum) && (colSum == uint(45))) {
			return false
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
