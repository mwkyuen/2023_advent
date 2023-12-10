package main
 
import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func is_zeros(nums []int) bool {
    for _, z := range nums {
        if z != 0 {
            return false
        }
    }
    return true
}

func get_differences(nums []int) []int {
    var diff []int 
    var d int
    for i:=0; i<len(nums)-1; i++ {
        d = nums[i+1] - nums[i]
        diff = append(diff, d)
    }
    return diff
}

func get_tree(nums []int) [][]int {
    var tree [][]int
    tree = append(tree, nums)
    diff := get_differences(nums)
    tree = append(tree, diff)
    for {
        if is_zeros(diff) {
            break
        } 
        diff = get_differences(diff)
        tree = append(tree, diff)
    }
    
    return tree
}

func get_next_value(tree *[][]int) int {
    diff := 0

    for i := len(*tree)-2; i >= 0; i-- {
        nums := (*tree)[i]
        val := nums[len(nums)-1]
        next_val := val + diff
        diff = next_val
    }
    return diff
}

func get_prev_value(tree *[][]int) int {
    diff := 0

    for i := len(*tree)-2; i >= 0; i-- {
        nums := (*tree)[i]
        val := nums[0]
        next_val := val - diff
        diff = next_val
    }
    return diff
}

func main() {

    filePath := os.Args[1]
    b, err := os.ReadFile(filePath) 
    if err != nil {
        fmt.Print(err)
    }

    content := string(b) 

    lines := strings.Split(content, "\n")
    lines = lines[:len(lines)-1]
    
    var vals []string
    sum := 0
    for _, line := range lines {
        vals = strings.Split(line, " ")
        var nums []int
        for _, v := range vals {
            num, _ := strconv.Atoi(v)
            nums = append(nums, num)
        }
        tree := get_tree(nums)
        // part 1
        // next_val := get_next_value(&tree)
        // sum += next_val
        // part 2
        prev_val := get_prev_value(&tree)
        sum += prev_val
    }
    fmt.Println(sum)
}


