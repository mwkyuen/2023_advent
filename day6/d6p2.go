package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func calculate_ways_to_win(race_time int, race_record int) int {
    sum := 0
    for i:=0; i<race_time; i++ {
        speed := i
        dist := (race_time - speed) * speed
        if dist > race_record {
            sum++
        }

    }
    return sum
}

func main() {

    filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    
    
    var input []int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        values := strings.Split(line, ":")[1]
        number, _ := strconv.Atoi(strings.ReplaceAll(values, " ", ""))
        input = append(input, number)
    }

    ways_to_win := calculate_ways_to_win(input[0], input[1])
    fmt.Println(ways_to_win)

}


