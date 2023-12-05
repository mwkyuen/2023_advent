package main
 
import (
    "fmt"
    "strings"
    "strconv"
)

type number struct {
    row_idx, xStart, xEnd, val int
    is_adjacent             bool
}

func markAdjacent(col_idx_of_symbol int, numbers []*number) {
    for _, n := range numbers {
        if col_idx_of_symbol >= n.xStart-1 && col_idx_of_symbol <= n.xEnd+1 {
            n.is_adjacent = true
        }
    }
}

func char_is_symbol(char rune) bool {
    return (char != '.' && !(char >= '0' && char <= '9'))
}

func main() {

    const sample = 
`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

    lines := strings.Split(sample, "\n")
    lines = lines[1:len(lines)-1]
    fmt.Println(len(lines))
    matrix := make([][]rune, len(lines))  // rune representation of each line
    numbersPerRow := make([][]*number, len(matrix))
    sum := 0

    for row_idx, line := range lines {
        rLine := []rune(line)
        matrix[row_idx] = rLine
        numbersPerRow[row_idx] = []*number{}
        curStart := -1
        curEnd := -1
        for col_idx, char := range rLine {
            if char >= '0' && char <= '9' {
                if curStart == -1 {
                    curStart = col_idx
                }
                curEnd = col_idx
                // number is at the end of the line
                if col_idx == len(rLine)-1 {
                    val, _ := strconv.Atoi(line[curStart : curEnd+1])
                    numbersPerRow[row_idx] = append(numbersPerRow[row_idx], &number{row_idx: row_idx, xStart: curStart, xEnd: curEnd, val: val})
                }
            } else {
                // end of number, but not end of line
                if curEnd != -1 {
                    val, _ := strconv.Atoi(line[curStart : curEnd+1])
                    numbersPerRow[row_idx] = append(numbersPerRow[row_idx], &number{row_idx: row_idx, xStart: curStart, xEnd: curEnd, val: val})
                    curStart = -1
                    curEnd = -1
                }
                // else: Not a number
            }
        }
    }
    
    for row_idx, rLine := range matrix {
        for col_idx, char := range rLine {
            if char_is_symbol(char) {
                if row_idx > 0 {
                    // mark above row
                    markAdjacent(col_idx, numbersPerRow[row_idx-1])
                }
                // mark current row
                markAdjacent(col_idx, numbersPerRow[row_idx])
                if row_idx < len(numbersPerRow)-1 {
                    // mark below row
                    markAdjacent(col_idx, numbersPerRow[row_idx+1])
                }
            }
        }
    }

    for _, numbers := range numbersPerRow {
        for _, number := range numbers {
            if number.is_adjacent {
                sum += number.val
            }
        }
    }

    fmt.Println(sum)

}

// func get_symbols_index(filePath string) [142][142]int {
//     re := regexp.MustCompile(`[^0-9.]`)
//     m := [142][142]int{}

//     readFile, err := os.Open(filePath)
  
//     if err != nil {
//         fmt.Println(err)
//     }
//     fileScanner := bufio.NewScanner(readFile)
 
//     fileScanner.Split(bufio.ScanLines)
    
//     row_idx := 1
//     for fileScanner.Scan() {
//         line := []byte(fileScanner.Text())
//         symbol_index := re.FindAllIndex(line, -1)

//         if symbol_index != nil{
//             for _, symbol := range symbol_index {
//                 col_idx := symbol[0]
//                 m[row_idx][col_idx] = 1
//             }
//         }
//         row_idx++
//     }
//     return m
// }


