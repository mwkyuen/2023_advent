package main
 
import (
	"bufio"
	"os"
    "fmt"
    "unicode"
    "strings"
)

func main() {
	viable_numbers := map[string]rune{
	    "one": '1',
	    "two": '2',
	    "three": '3',
	    "four": '4',
	    "five": '5',
	    "six": '6',
	    "seven": '7',
	    "eight": '8',
	    "nine": '9',
	}
	filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  	
  	var total int = 0
    for fileScanner.Scan() {
    	fmt.Println(fileScanner.Text())
    	var first rune = get_first_digit(fileScanner.Text(), viable_numbers)
		var last rune = get_last_digit(fileScanner.Text(), viable_numbers)
		var first_digit = int(first-'0') * 10
		var last_digit = int(last-'0')
		fmt.Println((first_digit + last_digit))
		total += (first_digit + last_digit)
    }
	
	
	fmt.Println(total)
	
}

func get_first_digit(text string, table map[string]rune) rune {
	for pos, char := range text {
		if unicode.IsDigit(char) {
			return char
		}
		for k, v := range table { 
			if strings.HasPrefix(text[pos:], k) {
				return v
			}
		}
	}
	return -1
}

func get_last_digit(text string, table map[string]rune) rune {
	for i := len(text)-1; i >= 0; i-- {
		if unicode.IsDigit(rune(text[i])) {
			return rune(text[i])
		}
		for k, v := range table { 
			if strings.HasPrefix(text[i:], k) {
				return v
			}
		}
	}
	return -1
}

