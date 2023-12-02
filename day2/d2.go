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
        game_id := get_game_id(line)
        game_results := get_game_results(line)

        if valid_game(game_results) {
            total += game_id
        }
    }
    fmt.Println(total)
}

func valid_game(game_results []Result) bool {
    bag := map[string]int {
        "max_red": 12,
        "max_green": 13,
        "max_blue": 14,
    }

    for _, result := range game_results {
        if result.red > bag["max_red"] {
            return false
        } else if result.green > bag["max_green"] {
            return false
        } else if result.blue > bag["max_blue"] {
            return false
        }
    }
    return true
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

