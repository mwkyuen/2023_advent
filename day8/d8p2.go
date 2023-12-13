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

// Modified from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
    a := integers[0]
    b := integers[1]
    result := a * b / GCD(a, b)

    for i := 2; i < len(integers); i++ {
        temp := []int{result, integers[i]}
        result = LCM(temp)
    }

    return result
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
        if key[2:3] == "A" {
            locations = append(locations, key)
        }
    }

    var next_direction rune
    
    var steps []int
    for _, location := range locations {
        temp_step := 0
        for i:=0; ;i++{
            if i == len(directions){
                i=0
            }
            next_direction = directions[i]
            if location[2:3] == "Z" {
                steps = append(steps, temp_step)
                break
            }
            if next_direction == 'R' {
                location = signposts[location].right
                temp_step++
            } else {
                location = signposts[location].left
                temp_step++
            }
        }
    }
    
    total := LCM(steps)
    fmt.Println(total)
}


