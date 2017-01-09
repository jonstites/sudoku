package sudoku


import (
	"fmt"
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

func (myGrid *grid) setCellOptions(rowNum int, colNum int, options bitarray) error {
	myCell, err := myGrid.getCell(rowNum, colNum)
	myCell.setOptions(options)
	return err
}

/*
func (myGrid *Puzzle) addGuess(rowNum int, colNum int, value int) {
	myCell := myGrid.getCell(rowNum, colNum)
	myCell.addGuess(value)
}

// If a guess was wrong, undo all squares that depended on the guess
func (myGrid *Grid) reset(guess uint) {
	for _, row := range myGrid.grid {
		for _, cell := range row {
			cell.reset(guess)
		}
	}
}
/*
func (myGrid *Puzzle) String() string {
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

/*
func calcValueOptions(cells []Cell) (valueSet, error) {
	var valueOptions valueSet
	valueOptions = fullSet
	for _, cell := range cells {
		if cell.valueKnown {
			valueOptions -= 1 << (uint(cell.value) - 1)
		}
	}
	return valueOptions, nil
}

/*
func (myGrid *Puzzle) updateRow(rowNum int) error {
	cells := make([]Cell, len(myGrid.puzzle))
	for colNum, cell := range myGrid.puzzle[rowNum] {
		cells[colNum] = cell
	}

	valueOptions, err := calcValueOptions(cells)

	for colNum, _ := range myGrid.puzzle[rowNum] {
		myGrid.setValueOptions(rowNum, colNum, valueOptions)
	}
	
	return err
}

func (myGrid *Puzzle) updateCol(colNum int) error {
	cells := make([]Cell, len(myGrid.puzzle))
	for rowNum, row := range myGrid.puzzle {
		cells[rowNum] = row[colNum]
	}

	valueOptions, err := calcValueOptions(cells)
	for rowNum, _ := range myGrid.puzzle {
		myGrid.setValueOptions(rowNum, colNum, valueOptions)
	}

	return err
}


func (myGrid *Puzzle) updateBlock(rowNum int, colNum int) error {
	cells := make([]Cell, len(myGrid.puzzle))
	blockRow := rowNum - (rowNum % 3)
	blockCol := colNum - (colNum % 3)
	i := 0
	for rowIndex := blockRow; rowIndex < blockRow + 3; rowIndex += 1 {
		for colIndex := blockCol; colIndex < blockCol + 3; colIndex += 1 {
			cells[i] = *myGrid.getCell(rowIndex, colIndex)
			i++
		}
	}

	valueOptions, err := calcValueOptions(cells)

	for rowIndex := blockRow; rowIndex < blockRow + 3; rowIndex += 1 {
		for colIndex := blockCol; colIndex < blockCol + 3; colIndex += 1 {
			myGrid.setValueOptions(rowIndex, colIndex, valueOptions)
		}
	}
	return err
		
}

func (myGrid *Puzzle) updateAll() error {
	fmt.Println("Updating all")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			myCell := myGrid.getCell(i, j)
			if myCell.guessNum > len(myGrid.guesses) {
				myGrid.puzzle[i][j].valueKnown = false
				myGrid.puzzle[i][j].valueOptions = fullSet
			}
		}
	}
	
	for i := 0; i < 9; i++ {
		err := myGrid.updateRow(i)
		if err != nil {
			return err
		}
		err = myGrid.updateCol(i)
		if err != nil {
			return err
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			err := myGrid.updateBlock(i, j)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (myGrid *Puzzle) fillOneCell() error {
	bestRowNum := 0
	bestColNum := 0
	bestNumOptions := 10
	bestCellValue := 0
	isComplete := true

	for rowNum, row := range myGrid.puzzle {
		for colNum, cell := range row {
			if cell.valueKnown {
				continue
			}

			isComplete = false
			numOptions := cell.numValueOptions()

			if (numOptions >= bestNumOptions) {
				continue
			}

			bestNumOptions = numOptions
			bestRowNum = rowNum
			bestColNum = colNum
			bestCellValue = cell.chooseValue()
		}
	}

	if isComplete {
		return fmt.Errorf("Can't fill a cell: sudoku complete.")
	}


	if bestNumOptions == 0 {
		guessSize := len(myGrid.guesses)
		fmt.Println("Problem cell is: ", bestRowNum, bestColNum)
		if guessSize == 0 {
			return fmt.Errorf("This sudoku puzzle cannot be solved.",
			myGrid.guesses, "\n", myGrid)
		}

		fmt.Println("Removed a guess at: ", myGrid.guesses[guessSize - 1])
		fmt.Println("Was:\n", myGrid)
		myGrid.guesses = myGrid.guesses[:guessSize - 1]
		myGrid.updateAll()
		fmt.Println("Is:\n", myGrid)
	}
	

	if bestNumOptions > 1 {
		coordinates := []int {bestRowNum, bestColNum}


		myGrid.guesses = append(myGrid.guesses, coordinates)
		myGrid.addGuess(bestRowNum, bestColNum, bestCellValue)
		fmt.Println("Made a guess at: ", bestRowNum, bestColNum, bestCellValue)
		fmt.Println("Was:\n", myGrid)
	}

	myGrid.setValue(bestRowNum, bestColNum, bestCellValue)
	myGrid.updateRow(bestRowNum)
	myGrid.updateCol(bestColNum)
	myGrid.updateBlock(bestRowNum, bestColNum)
	if bestNumOptions > 1 {
		fmt.Println("Is:\n", myGrid)
	}
	return nil
}


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

func (myGrid *Puzzle) insertRow(rowNum int, row string) {
	for colNum, char := range row {
		value := int(char-'0')
		if value != 0 {
			myGrid.setValue(rowNum, colNum, value)
		}
	}
}



func isNumeric(row string) bool {
	for _, char := range row {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func validateRowFormat(row string) error {
	if !(isNumeric(row)) {
		return fmt.Errorf("File contains non-numeric characters in row: ", row)
	}

	if len(row) != 9 {
		return fmt.Errorf("Row does not contain 9 digits: ", row) 
	}

	return nil
}

func puzzleFromFile(filename string) *Puzzle {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myGrid := newPuzzle()
	
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i >= 9 {
			log.Fatal("Files contains more than 9 rows.")
		}
		err := validateRowFormat(row)
		if err != nil {
			log.Fatal(err)
		}

		myGrid.insertRow(i, row)
		i++
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	myGrid.updateAll()
	return myGrid
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
