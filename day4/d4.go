package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "math"
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

func calculate_score(number_of_winning_numbers int) int {
    return 1 * int(math.Pow(float64(2), float64(number_of_winning_numbers-1)))
}

func main() {

    filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    
    total := 0
    for fileScanner.Scan() {
        line := fileScanner.Text()
        card := strings.Split(strings.Split(line, ":")[1], "|")
        winning_numbers := get_numbers(strings.TrimSpace(card[0]))
        numbers_i_have := get_numbers(strings.TrimSpace(card[1]))
        intersect := intersect_slice(&winning_numbers, &numbers_i_have)
        score := calculate_score(len(intersect))
        total += score
    }
    fmt.Println(total)
}


