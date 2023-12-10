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

func main() {

    filePath := os.Args[1]
    b, err := os.ReadFile(filePath) 
    if err != nil {
        fmt.Print(err)
    }

    content := string(b) 

    lines := strings.Split(content, "\n")
    directions := []rune(lines[0])
    lines = lines[2:len(lines)-1]
    signposts := make(map[string]Directions, len(lines))
    var s []string
    var key string
    var val string
    for _, line := range lines {
        s = strings.Split(line, "=")
        key = strings.TrimSpace(s[0])
        val = strings.TrimSpace(s[1])
        left_right := strings.Split(val[1:len(val)-1], ",")
        signposts[key] = Directions{left:strings.TrimSpace(left_right[0]), right:strings.TrimSpace(left_right[1])}
    }

    location := "AAA"
    var next_direction rune
    steps := 0
    for i:=0; ;i++{
        if i == len(directions){
            i=0
        }
        next_direction = directions[i]
        if location == "ZZZ" {
            break
        }
        if next_direction == 'R' {
            location = signposts[location].right
            steps++
        } else {
            location = signposts[location].left
            steps++
        }

        fmt.Println(next_direction)
        fmt.Println(location)
    }
    
    fmt.Println(steps)
}


