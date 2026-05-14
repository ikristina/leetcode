package problems

func pacificAtlantic(heights [][]int) [][]int {
    result := [][]int{}
    m := len(heights)
    n := len(heights[0])

    pacificQueue := [][]int{}
    for i := range m {
        pacificQueue = append(pacificQueue, []int{i, 0})
    }
    for i := range n {
        pacificQueue = append(pacificQueue, []int{0, i})
    }

    pacificVisited := make([][]bool, m)
    for i, _ := range pacificVisited {
        pacificVisited[i] = make([]bool, n)
    }

    for len(pacificQueue) > 0 {

        r, c := pacificQueue[0][0], pacificQueue[0][1] // row-column pair
        pacificQueue = pacificQueue[1:]

        if pacificVisited[r][c] == true {
            continue
        }

        pacificVisited[r][c] = true

        if r - 1 >= 0 && heights[r-1][c] >= heights[r][c]{
            pacificQueue = append(pacificQueue, []int{r-1, c})
        }
        if r + 1 < m && heights[r+1][c] >= heights[r][c]{
            pacificQueue = append(pacificQueue, []int{r+1, c})
        }
        if c - 1 >= 0 && heights[r][c-1] >= heights[r][c]{
            pacificQueue = append(pacificQueue, []int{r, c-1})
        }
        if c + 1 < n && heights[r][c+1] >= heights[r][c]{
            pacificQueue = append(pacificQueue, []int{r, c+1})
        }
    }


    atlanticQueue := [][]int{}
    for i := range m {
        atlanticQueue = append(atlanticQueue, []int{i, n-1})
    }
    for i := range n {
        atlanticQueue = append(atlanticQueue, []int{m-1, i})
    }

    atlanticVisited := make([][]bool, m)
    for i := range atlanticVisited {
        atlanticVisited[i] = make([]bool, n)
    }
    for len(atlanticQueue) > 0 {

        r, c := atlanticQueue[0][0], atlanticQueue[0][1] // row-column pair
        atlanticQueue = atlanticQueue[1:]

        if atlanticVisited[r][c] == true {
            continue
        }

        atlanticVisited[r][c] = true

        if r - 1 >= 0 && heights[r-1][c] >= heights[r][c] {
            atlanticQueue = append(atlanticQueue, []int{r-1, c})
        }
        if r + 1 < m && heights[r+1][c] >= heights[r][c]{
            atlanticQueue = append(atlanticQueue, []int{r+1, c})
        }
        if c - 1 >= 0 && heights[r][c-1] >= heights[r][c]{
            atlanticQueue = append(atlanticQueue, []int{r, c-1})
        }
        if c + 1 < n && heights[r][c+1] >= heights[r][c]{
            atlanticQueue = append(atlanticQueue, []int{r, c+1})
        }
    }

    for r := range pacificVisited {
        for c := range pacificVisited[0] {
            if pacificVisited[r][c] == true && atlanticVisited[r][c] == true {
                result = append(result, []int{r, c})
            }
        }
    }

    return result
}
