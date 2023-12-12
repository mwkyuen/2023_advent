package main
 
import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "math"
)

type Group struct {
    lower_dest, lower_src, upper_dest, upper_src int
}

func convert_to_int(list_strings []string) []int {
    var nums []int
    for _, num_s := range list_strings {
        num, _ := strconv.Atoi(num_s)
        nums = append(nums, num)
    }
    return nums
}

func get_location(almanac *map[string][]Group, seed int) int {
    soil := get_next_val((*almanac)["seed-to-soil"], seed)
    fertilizer := get_next_val((*almanac)["soil-to-fertilizer"], soil)
    water := get_next_val((*almanac)["fertilizer-to-water"], fertilizer)
    light := get_next_val((*almanac)["water-to-light"], water)
    temp := get_next_val((*almanac)["light-to-temperature"], light)
    hum := get_next_val((*almanac)["temperature-to-humidity"], temp)
    loc := get_next_val((*almanac)["humidity-to-location"], hum)
    
    return loc

}

func get_next_val (groups []Group, origin int) int {
    next_val := -1
    for _, g := range groups {
        if origin >= g.lower_src && origin <= g.upper_src {
            next_val = origin - g.lower_src + g.lower_dest
        } 
    } 
    if next_val == -1 {
        next_val = origin
    } 
    return next_val
}

func set_group(num []int) Group {
    g := Group{lower_dest: num[0], lower_src: num[1], upper_dest: (num[0] + num[2] - 1), upper_src: (num[1] + num[2] - 1)}
    return g
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func get_range_of_seeds(nums []int) []int {
    var num_start int
    var num_range int
    var seed_range []int
    for i := 0; i < len(nums); i += 2 {
        var mrange []int
        num_start = nums[i]
        num_range = nums[i+1]
        mrange = makeRange(num_start, num_start + num_range - 1 )
        seed_range = append(seed_range, mrange...)
    }
    return seed_range
}

func main() {

    filePath := os.Args[1]
    b, err := os.ReadFile(filePath) 
    if err != nil {
        fmt.Print(err)
    }

    content := string(b) 

    lines := strings.Split(content, "\n")
    temp := strings.Split(lines[0], " ")
    seeds := get_range_of_seeds(convert_to_int(temp[1:len(temp)]))
    lines = lines[2:len(lines)]
    var temp_line []string
    almanac := make(map[string][]Group)
    var name string
    var nums []Group
    for _, line := range lines {
        temp_line = strings.Split(line, " ")
        var num []int
        if len(temp_line) == 2 {
            name = temp_line[0]
        } else if len(temp_line) == 3 {
            num = convert_to_int(temp_line)
            g := set_group(num)
            nums = append(nums, g)
        }else if len(temp_line) == 1 {
            almanac[name] = nums
            nums = nil   
        }
    }

    fmt.Println(almanac)
    fmt.Println(seeds)
    
    lowest := math.MaxInt32
    for _, seed := range seeds {
        location := get_location(&almanac, seed)
        if location < lowest {
            lowest = location
        }
    }
    
    fmt.Println(lowest)
}


