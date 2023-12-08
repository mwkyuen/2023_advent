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
    
    
    var m [][]int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        values := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
        var t []int
        for _, v := range values {
            if len(v) > 0  {
                num, _ := strconv.Atoi(v)
                t = append(t, num)
            }   
        }
        m = append(m, t)
    }

    margin := 1
    for i:=0; i<len(m[0]); i++ {
        race_time := m[0][i]
        race_record := m[1][i]
        ways_to_win := calculate_ways_to_win(race_time, race_record)
        margin = margin * ways_to_win
    }

    fmt.Println(margin)

}


