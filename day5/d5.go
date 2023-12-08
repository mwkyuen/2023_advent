package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "math"
)

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
   
    }
    fmt.Println(total)
}


