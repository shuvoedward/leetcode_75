package main

import (
	"strconv"
	"strings"
)

// grid = [[3,2,1],[1,7,6],[2,7,7]]

func equalPairs(grid [][]int) int {
	n := len(grid)

	rowFrequencies := make(map[string]int)

	for r := 0; r < n; r++ {
		rowStr := buildStringFromSlice(grid[r])
		rowFrequencies[rowStr]++
	}

	count := 0
	for c := 0; c < n; c++ {

		currentCol := make([]int, n)
		for r := 0; r < n; r++ {
			currentCol[r] = grid[r][c]
		}

		colStr := buildStringFromSlice(currentCol)

		if freq, ok := rowFrequencies[colStr]; ok {
			count += freq
		}
	}

	return count

}

func buildStringFromSlice(s []int) string {
	strBuilder := strings.Builder{}
	for i, val := range s {
		strBuilder.WriteString(strconv.Itoa(val))
		// add a comma if it is not the last element
		if i < len(s)-1 {
			strBuilder.WriteByte(',')
		}
	}
	return strBuilder.String()
}

func hash(hash, val int) int {
	return hash*37 + val
}

func equalPairsHash(grid [][]int) int {
	rows := make([]int, 0)
	cols := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		row_hash := 0
		col_hash := 0
		for j := 0; j < len(grid[0]); j++ {
			// row
			row_hash = row_hash*37 + grid[i][j]
			col_hash = col_hash*37 + grid[j][i]
		}

		rows = append(rows, row_hash)
		cols = append(cols, col_hash)
	}

	result := 0
	for _, row := range rows {
		for _, col := range cols {
			if row == col {
				result += 1
			}
		}
	}

	return result

}
