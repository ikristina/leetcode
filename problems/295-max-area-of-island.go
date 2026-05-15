package problems

func maxAreaOfIsland(grid [][]int) int {
	maxIsland := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 1 {
				area := getArea(grid, i, j)
				maxIsland = max(maxIsland, area)
			}
		}
	}

	return maxIsland
}

func getArea(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0
	if i > 0 {
		area = area + getArea(grid, i - 1, j)
	}
	if j > 0 {
		area = area + getArea(grid, i, j - 1)
	}
	if i < len(grid)-1 {
		area = area + getArea(grid, i + 1, j)
	}
	if j < len(grid[0])-1 {
		area = area + getArea(grid, i, j + 1)
	}
	return area
}
