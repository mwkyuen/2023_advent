package main
 
import (
    "os"
    "fmt"
    "strings"
)

type Directions struct {
    left string
    right string
}

func check_cond (locations []string) bool {
    for _, loc:= range locations {
        if loc[2:3] != "Z" {
            return false
        } 
    }
    return true
}

func update_locations (locations []string, signposts map[string]Directions, d rune) {
    for i, loc := range locations{
        if d == 'R'{
            locations[i] = signposts[loc].right
        } else {
            locations[i] = signposts[loc].left
        }
    }
}

func main() {

    filePath := os.Args[1]
    b, err := os.ReadFile(filePath) 
    if err != nil {
        fmt.Print(err)
    }

    content := string(b) 

    lines := strings.Split(content, "\n")
    // directions := []rune(lines[0])
    lines = lines[2:len(lines)-1]
    signposts := make(map[string]Directions, len(lines))
    var locations []string
    var s []string
    var key string
    var val string
    for _, line := range lines {
        s = strings.Split(line, "=")
        key = strings.TrimSpace(s[0])
        val = strings.TrimSpace(s[1])
        left_right := strings.Split(val[1:len(val)-1], ",")
        signposts[key] = Directions{left:strings.TrimSpace(left_right[0]), right:strings.TrimSpace(left_right[1])}
        if key[2:3] == "Z" {
            locations = append(locations, key)
        }
    }
    fmt.Println(locations)
    fmt.Println(signposts)
    // var next_direction rune
    
    // steps := 0
    // for i:=0; ;i++{
    //     if i == len(directions){
    //             i=0
    //         }
    //     next_direction = directions[i]
    //     if next_direction == 'R' {
    //         update_locations(locations, signposts, 'R')
    //         steps++
    //     } else {
    //         update_locations(locations, signposts, 'L')
    //         steps++
    //     }
    //     if check_cond(locations) {
    //         break
    //     }
        
    // }
    
    // fmt.Println(steps)
}


