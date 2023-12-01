package main
 
import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
 
    filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  	
  	var total = 0
    for fileScanner.Scan() {
        var numbers = find_digit_numbers(fileScanner.Text())
        // var numbers = find_all_numbers(fileScanner.Text())
        fmt.Println(fileScanner.Text())
        fmt.Println(numbers)
        var malformed_number = get_correct_number(numbers)

        // var wellformed_number = word_to_number(malformed_number)
        num, err := strconv.Atoi(strings.Join(wellformed_number, ""))
        fmt.Println(num)
        if err != nil {
	        fmt.Println(err)
	    }
        total += num
    }
  	fmt.Println(total)
    readFile.Close()
}


func find_digit_numbers(text string) []string {
	re := regexp.MustCompile(`\d`)
	numbers := re.FindAllString(text, -1)
	return numbers
}

func find_all_numbers(text string) []string {
	// Only returns non-overlapping occurrences of the regexp pattern!!
	// re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	// numbers := re.FindAllString(text, -1)
	// return numbers

	
}

func get_correct_number(numbers []string) []string {
	var correct_number []string
	correct_number = append(correct_number, numbers[0])
	correct_number = append(correct_number, numbers[len(numbers) - 1])
    return correct_number
}

func word_to_number(values []string) []string {
	translate_table := map[string]string{
	    // "zero": "0",
	    "one": "1",
	    "two": "2",
	    "three": "3",
	    "four": "4",
	    "five": "5",
	    "six": "6",
	    "seven": "7",
	    "eight": "8",
	    "nine": "9",
	}
	for i, v := range values {
    	translated_val, ok := translate_table[v]
    	if ok {
    		values[i] = translated_val
    	}
    }
    return values

}
