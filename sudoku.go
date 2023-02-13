package main

import "fmt"

const size = 9

func main() {
    grid := [size][size]int{
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }
    if solveSudoku(grid) {
        printGrid(grid)
    } else {
        fmt.Println("Impossible de résoudre ce sudoku.")
    }
}

func solveSudoku(grid [size][size]int) bool {
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if grid[i][j] == 0 {
                for k := 1; k <= size; k++ {
                    if isValid(grid, i, j, k) {
                        grid[i][j] = k
                        if solveSudoku(grid) {
                            return true
                        }
                        grid[i][j] = 0
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(grid [size][size]int, row int, col int, k int) bool {
    for i := 0; i < size; i++ {
        if grid[row][i] == k || grid[i][col] == k {
            return false
        }
    }
    rowStart := (row / 3) * 3
    colStart := (col / 3) * 3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[rowStart+i][colStart+j] == k {
                return false
            }
        }
    }
    return true
}

func printGrid(grid [size][size]int) {
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            fmt.Print(grid[i][j], " ")
        }
        fmt.Println()
    }
}

