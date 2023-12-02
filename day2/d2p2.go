package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

type Result struct {
    red int
    green int
    blue int
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
        // game_id := get_game_id(line)
        game_results := get_game_results(line)

        minimum_set := get_minimum_set(game_results) 
        power := minimum_set.red * minimum_set.green * minimum_set.blue
        total += power
    }
    fmt.Println(total)
}

func get_minimum_set(game_results []Result) Result {
    minimum_set := Result{red: 0, green: 0, blue: 0}

    for _, result := range game_results {
        if result.red > minimum_set.red {
            minimum_set.red = result.red
        } 
        if result.green > minimum_set.green {
            minimum_set.green = result.green
        } 
        if result.blue > minimum_set.blue {
            minimum_set.blue = result.blue
        }
    }
    return minimum_set
}

func get_game_id(line string) int {
    game_id := strings.Split(line, ":")[0]
    id, _ := strconv.Atoi(strings.Split(game_id, " ")[1])
    return id
}

func get_game_results(line string) []Result {
    var results []Result
    game_result := strings.Split(line, ":")[1]
    games := strings.Split(game_result, ";")
    for _, game := range games {
        game = strings.TrimSpace(game)
        colours := strings.Split(game, ",")
        result := process_game_result(colours)
        results = append(results, result)
    }
    return results
}

func process_game_result(colours []string) Result {
    ordered_result := [3]int{0, 0, 0}
    for _, colour := range colours {
        balls := strings.Split(strings.TrimSpace(colour), " ")
        val, _ := strconv.Atoi(balls[0])
        if balls[1] == "red" {
            ordered_result[0] = val
        } else if balls[1] == "green" {
            ordered_result[1] = val
        } else {
            ordered_result[2] = val
        }
    }

    return Result{
        red: ordered_result[0], 
        green: ordered_result[1], 
        blue: ordered_result[2],
    }
}
