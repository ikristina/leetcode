package problems

func orangesRotting(grid [][]int) int {
	// 2 - rotten
	// 1 - fresh
	// 0 - empty

	// search 2s

	minutes := 0
	fresh := 0

    q := [][]int{}

    for i, arr := range grid {
        for j, n := range arr {
            if n==2 {
                q = append(q, []int{i, j})
            }
        }
    }


    // search fresh
    for _, arr := range grid {
        for _, n := range arr {
            if n==1 {
                fresh++
            }
        }
    }


    for len(q) > 0 {
        size := len(q)
        hasRot := false
        for i := 0; i < size; i++ {
            cs := q[0] // coordinates e.g. [0,1]
            q = q[1:]

            if cs[0] > 0 && grid[cs[0]-1][cs[1]] == 1{
                grid[cs[0]-1][cs[1]] = 2
                q = append(q, []int{cs[0]-1, cs[1]})
                fresh--
                hasRot = true
            }
            if cs[1] > 0 && grid[cs[0]][cs[1]-1] == 1 {
                grid[cs[0]][cs[1]-1] = 2
                q = append(q, []int{cs[0], cs[1]-1})
                fresh--
                hasRot = true
            }
            if cs[0] < len(grid)-1 && grid[cs[0]+1][cs[1]] == 1 {
                grid[cs[0]+1][cs[1]] = 2
                q = append(q, []int{cs[0]+1, cs[1]})
                fresh--
                hasRot = true
            }
            if cs[1] < len(grid[0])-1 && grid[cs[0]][cs[1]+1] == 1 {
                grid[cs[0]][cs[1]+1] = 2
                q = append(q, []int{cs[0], cs[1]+1})
                fresh--
                hasRot = true
            }
        }
        if hasRot == true {
            minutes ++
        }
    }

    if fresh > 0 {
        return -1
    }

	return minutes
}
