func rotateTheBox(boxGrid [][]byte) [][]byte {
    m := len(boxGrid)
    n := len(boxGrid[0])

    //mMinOne := m - 1
    nMinOne := n - 1

    newGrid := make([][]byte, n)
    for row:= 0; row < n; row++ {
        newGrid[row] = make([]byte, len(boxGrid))
    }

    oldRow := m
    for col :=0; col < m; col++ {
        oldRow--

        row := nMinOne
        for oldCol := nMinOne; oldCol >=0; oldCol-- {
            switch boxGrid[oldRow][oldCol] {
                case '#':
                    newGrid[row][col] = '#'
                    row--
                case '*':
                    for row > oldCol{
                        newGrid[row][col] = '.'
                        row--
                    }
                    newGrid[row][col] = '*'
                    row--
            }
        }

        for ;row >=0; row-- {
            newGrid[row][col] = '.'
        }
    }

    return newGrid
}