package sudoku

import (
	"fmt"
	"strconv"
)

// Cell represents one square of a Sudoku puzzle.
type Cell struct {
	value int
}

// Just print the value of a Cell 
func (c *Cell) String() string {
	return strconv.Itoa(c.value)
}

func main() {
	fmt.Printf("hello, world\n")
}
