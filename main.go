package main
import "fmt"

// func ticketPrice(price int, session bool, weekend bool) int {
//     if weekend == true && session == true {
//         return price + 150
//     }
//     if session == true {
//         return price + 50
//     }
//     if weekend == true {
//         return price + 100
//     }
//
//     return price
// }
//
// func formatReceipt(titleFilm string, quantityTicket int, cost int) (string, int) {
//     return titleFilm, quantityTicket * cost
// }
//
// func isAllowed(age int, rate int) bool {
//     if age < rate {
//         return false
//     }
//     return true
// }

// var titleFilms = []string{"Гарри Поттер", "Паук", "Начало", "Наруто"}
// var times = []int{10, 12, 15, 17}

// func sum(nums []int) int {
//     total := 0
//     for _, r := range nums {
//         total += r
//     }
//     return total
// }

// func removeDuplicates(st []string) []string {
//     seen := make(map[string] struct{})
//     result := make([]string, 0, len(st))
//
//     for _, el := range st {
//         if _, ok := seen[el]; !ok {
//             seen[el] = struct{}{}
//             result = append(result, el)
//         }
//     }
//
//     return result
// }

// func countChars(s string) map[rune]int {
//     seen := make(map[rune]int)
//
//     for _, r := range s {
//         seen[r]++
//     }
//
//     return seen
// }

func hasDuplicates(nums []int) bool {
    seen := make(map[int]struct{})
    for _, el := range nums {
        if _, ok := seen[el]; ok {
            return true
        }
        seen[el] = struct{}{}
    }
    return false
}

func main() {

    fmt.Println(hasDuplicates([]int{1,2,3,1}))

//         st := []string { "apple", "banana", "apple", "orange", "banana" }
//         fmt.Println(removeDuplicates(st))

//     	fmt.Println(ticketPrice(10, true, true))   // 160
//     	fmt.Println(ticketPrice(10, true, false))  // 60
//     	fmt.Println(ticketPrice(10, false, true))  // 110
//     	fmt.Println(ticketPrice(10, false, false)) // 10
//
//     	fmt.Println(formatReceipt("Гарри Поттер", 30, 10))
//
//     	fmt.Println(isAllowed(18, 18))
//     	fmt.Println(isAllowed(12, 18))

//         fmt.Println("Рассписание кинотеатра")

//         for zal := 1; zal <= 3; zal++ {
//         	fmt.Printf("Зал %d:\n", zal)
//         	for i := range titleFilms {
//         		fmt.Printf("%d - %s\n", times[i], titleFilms[i])
//         	}
//         }

//         fmt.Println("Поиск. Найти начало")

//     for _, r := range titleFilms {
//        fmt.Println([]rune(r))
//     }

}