package sudoku


import (
	"fmt"
	"log"
)


// Get the coordinates of the unfilled cell with fewest possible options
func getEasiestCell(myGrid *grid) (bestRowNum int, bestColNum int) {
	bestNumOptions := 10
	for rowNum, row := range myGrid.grid {
		for colNum, myCell := range row {
			if myCell.isKnown() {
				continue
			}
			numOptions := myCell.options.numValuesTrue()
			if numOptions < bestNumOptions {
				bestRowNum = rowNum
				bestColNum = colNum
				bestNumOptions = numOptions
			}
		}
	}
	return bestRowNum, bestColNum
}

// There is at least one square with no options
func notSolvable(myGrid *grid) bool {
	row, col := getEasiestCell(myGrid)
	myCell, _ := myGrid.getCell(row, col)
	numOptions := myCell.options.numValuesTrue()
	return numOptions == 0
}

// If the sudoku is stuck, undo the guesses made one at a time
func undoMistakes(myGrid *grid, tryCounter int) (int, error) {
	needUndo := notSolvable(myGrid)
	for needUndo == true {
		if tryCounter <= 0 {
			log.Fatal("Sudoku cannot be solved:", myGrid)
		}

		tryCounter -= 1
		myGrid.reset(tryCounter)
		needUndo = notSolvable(myGrid)
	}
	return tryCounter, nil
}

// Fill in one cell of the grid
func fillCell(myGrid *grid, tryCounter int) (int, error) {
	myGrid.reset(tryCounter)
	
	if myGrid.isComplete() {
		return tryCounter, fmt.Errorf("Sudoku already complete.")
	}

	tryCounter, err := undoMistakes(myGrid, tryCounter)
	if err != nil {
		return tryCounter, err
	}

	row, col := getEasiestCell(myGrid)
	myCell, _ := myGrid.getCell(row, col)
	value := myCell.options.lowestValue()

	if myCell.options.numValuesTrue() > 1 {
		tryCounter += 1
		myCell.setTriedValue(value)
	}

	myGrid.setCellValue(row, col, value, tryCounter)
	myGrid.reset(tryCounter)
	return tryCounter, nil
}



/*
func (myGrid *Puzzle) fillAllCells() error {
	isComplete, err := myGrid.isComplete()
	if err != nil {
		return err
	}
	for !(isComplete) {
		err := myGrid.fillOneCell() 
		if err != nil {
			return err
		}
		isComplete, err = myGrid.isComplete()
		if err != nil {
			return err
		}
	}
	isComplete, _ = myGrid.isComplete()
	fmt.Println(isComplete)
	return nil
}

func (myGrid *Puzzle) isComplete() (bool, error) {
	for _, row := range myGrid.puzzle {
		for _, cell := range row {
			if !cell.valueKnown {
				return false, nil
			}
			if cell.value == 0 {
				return false, fmt.Errorf("Error: Sudoku solution contains a zero.")
			}
		}
	}
	return true, nil
}


func main() {
	sudokuFile := flag.String("filename", "", "File of sudoku puzzle to solve.")
	flag.Parse()
	myGrid := puzzleFromFile(*sudokuFile)
	err := myGrid.fillAllCells()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myGrid)
}
*/
