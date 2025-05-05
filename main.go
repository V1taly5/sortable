package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	count, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || count <= 0 {
		return
	}

	mainMas := make([][]int, count)
	for i := range count {
		mainMas[i] = make([]int, count)
	}

	for i := range count {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for j := range count {
			mainMas[i][j], _ = strconv.Atoi(line[j])
		}
	}
	if isSortable(count, mainMas) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}

func isSortable(n int, matrix [][]int) bool {
	containerSums := make([]int, n)
	colorSums := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			containerSums[i] += matrix[i][j]
			colorSums[j] += matrix[i][j]
		}
	}

	sort.Ints(containerSums)
	sort.Ints(colorSums)

	for i := 0; i < n; i++ {
		if containerSums[i] != colorSums[i] {
			return false
		}
	}
	return true
}
