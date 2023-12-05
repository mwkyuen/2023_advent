package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func get_numbers(numbers string) []int {
    var res_num []int
    for _, number := range strings.Split(numbers, " ") {
        num, err := strconv.Atoi(number)
        if err == nil {
            res_num = append(res_num, num)
        }
    }
    return res_num
}

func intersect_slice(s1 *[]int, s2 *[]int) []int {
    var intersect []int
    for _, n1 := range *s1 {
        for _, n2 := range *s2 {
            if n1 == n2 {
                intersect = append(intersect, n2)
            }
        }
    }
    return intersect
}

func expand_card_set(row []int) [][]int {
    var temp_matrix [][]int
    for j, val := range row {
        if val != 0 {
            new_col := row[j+1:j+val+1]
            temp_matrix = append(temp_matrix, new_col)
        }
    }
    return temp_matrix
}

func score_game(card_matrix *[][]int) int {
    total := 0
    for _, elem := range *card_matrix {
        total += len(elem)
    }
    return total
}



func main() {

    filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    
    row_idx := 0
    
    var temp_slice []int 
    var temp_matrix [][]int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        card := strings.Split(strings.Split(line, ":")[1], "|")
        winning_numbers := get_numbers(strings.TrimSpace(card[0]))
        numbers_i_have := get_numbers(strings.TrimSpace(card[1]))
        intersect := intersect_slice(&winning_numbers, &numbers_i_have)

        number_of_scratchcards_won := len(intersect)
        temp_slice = append(temp_slice, number_of_scratchcards_won)
    
        row_idx++
    }

    temp_matrix = append(temp_matrix, temp_slice)
    
    var card_matrix [][]int
    for len(temp_matrix) != 0 {
        x := temp_matrix[0]
        temp_matrix = temp_matrix[1:]
        card_matrix = append(card_matrix, x)
        m := expand_card_set(x)
        for _, n := range m {
            temp_matrix = append(temp_matrix, n)
        }

    }
    
    total := score_game(&card_matrix)
    fmt.Println(total)

}


