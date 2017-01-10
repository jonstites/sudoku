package sudoku


import (
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
func undoMistakes(myGrid *grid, tried [][]int) error {
	needUndo := notSolvable(myGrid)
	for needUndo == true {
		if len(tried) == 0 {
			log.Fatal("Sudoku cannot be solved.")
		}

		tried = tried[:len(tried) - 1]
		myGrid.reset(len(tried))
		needUndo = notSolvable(myGrid)
	}
	return nil
}

// Fill in one cell of the grid
func fillCell(myGrid *grid, tried [][]int) error {
	if myGrid.isComplete() {
		log.Fatal("Sudoku already complete.")
	}

	err := undoMistakes(myGrid, tried)
	if err != nil {
		log.Fatal(err)
	}

	row, col := getEasiestCell(myGrid)
	myCell, _ := myGrid.getCell(row, col)
	value := myCell.options.lowestValue()
	myGrid.setCellValue(row, col, value)
	myGrid.reset(len(tried))
	return nil
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
