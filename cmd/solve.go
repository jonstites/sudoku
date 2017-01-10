package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"github.com/jonstites/sudoku"
)

func main() {
	sudokuFile := flag.String("filename", "", "File of sudoku puzzle to solve.")
	flag.Parse()

	if *sudokuFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	myGrid := sudoku.Read(*sudokuFile)
	err := sudoku.Solve(myGrid, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myGrid)
}
