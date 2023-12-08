package main
 
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

type handValue int
type cardValue int

const (
    high_card handValue = iota  
    one_pair handValue = iota  
    two_pair handValue = iota  
    three_of_a_kind handValue = iota
    full_house handValue = iota
    four_of_a_kind handValue = iota
    five_of_a_kind handValue = iota
)
const (
    c2 cardValue = iota  
    c3 cardValue = iota  
    c4 cardValue = iota  
    c5 cardValue = iota
    c6 cardValue = iota
    c7 cardValue = iota
    c8 cardValue = iota
    c9 cardValue = iota
    c10 cardValue = iota
    cJ cardValue = iota
    cQ cardValue = iota
    cK cardValue = iota
    cA cardValue = iota
)

func get_card_val(char rune) cardValue {
    if char == '2' {
        return c2
    } else if char == '3' {
        return c3
    } else if char == '4' {
        return c4
    } else if char == '5' {
        return c5
    } else if char == '6' {
        return c6
    } else if char == '7' {
        return c7
    } else if char == '8' {
        return c8
    } else if char == '9' {
        return c9
    } else if char == 'T' {
        return c10
    } else if char == 'J' {
        return cJ
    } else if char == 'Q' {
        return cQ
    } else if char == 'K' {
        return cK
    } else {
        return cA
    } 
}

func get_hand_value(card string) handValue {
    counter := make(map[rune]int)    
    for _, char := range card {
        counter[char]++
    }
    if len(counter) == 5 {
        return high_card
    } else if len(counter) == 1 {
        return five_of_a_kind
    } else if len(counter) == 4 {
        return one_pair
    } else if len(counter) == 3 {
        for _, v := range counter {
            if v == 3 {
                return three_of_a_kind
            } 
        }
        return two_pair
    } else {
        for _, v := range counter {
            if v == 4 {
                return four_of_a_kind
            } 
        }
        return full_house
    }
}

func is_still_bigger(card1 string, card2 string) bool {
    r1 := []rune(card1)
    r2 := []rune(card2)
    for i:=0; i < len(r1); i++ {
        card_val_1 := get_card_val(r1[i])
        card_val_2 := get_card_val(r2[i])
        if (card_val_1 != card_val_2) {
            return card_val_1 > card_val_2 
        } 
    }
    return false
}

func is_bigger(card1 string, card2 string) bool {
    hand_value_1 := get_hand_value(card1)
    hand_value_2 := get_hand_value(card2)
    if (hand_value_1 == hand_value_2) {
        return is_still_bigger(card1, card2)
    } else {
        return hand_value_1 > hand_value_2 
    }
}

func bubble_sort(cards *[]string) {
    for i := 0; i < len(*cards)-1; i++ {
        for j := 0; j < len(*cards)-i-1; j++ {
            if is_bigger((*cards)[j], (*cards)[j+1]) {
                (*cards)[j], (*cards)[j+1] = (*cards)[j+1], (*cards)[j]
            }
        }
    }
}

func main() {

    filePath := os.Args[1]
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    
    cardMap := make(map[string]int)
    var cardList []string
    for fileScanner.Scan() {
        line := fileScanner.Text()
        card_bid_pair := strings.Split(line, " ")
        card := card_bid_pair[0]
        bid, _ := strconv.Atoi(card_bid_pair[1])
        cardMap[card] = bid
        cardList = append(cardList, card)
    }

    bubble_sort(&cardList)
    fmt.Println(cardList)

    total := 0
    for i, card := range cardList {
        score := (i+1) * cardMap[card]
        total += score
    }
    fmt.Println(total)
}


