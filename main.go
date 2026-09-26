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
//     mapa := make(map[rune]int)
//
//     for _, el := range s {
//         if _, ok := mapa[el]; !ok {
//             mapa[el] = 1
//         } else {
//             mapa[el] += 1
//         }
//     }
//
//     return mapa
// }

// func hasDuplicates(arr []int) bool {
//     mapa := make(map[int]struct{}, len(arr))
//
//     for _, el := range arr {
//         if _, ok := mapa[el]; ok {
//             return true
//         }
//         mapa[el] = struct{}{}
//     }
//
//     return false
// }

// var catalog = map[string]float64 {
//                                  	"Побег из Шоушенка": 9,
//                                  	"Крёстный отец":     9,
//                                  	"Тёмный рыцарь":     9,
//                                  	"Криминальное чтиво": 9,
//                                  	"Форрест Гамп":      9,
//                                  }
//
// var film = map[string]float64 {
//     "Форрест Гамп 2":      9.9,
// }
//
// func addFilm(catalog map[string]float64, film map[string]float64 ) map[string]float64 {
//     for title, rating := range film {
//         catalog[title] = rating
//     }
//
//     return catalog
// }
//
// func topFilms(catalog map[string]float64, min float64) []string {
//     slice := []string{}
//
//     for title, rating := range catalog {
//         if rating >= min {
//             slice = append(slice, title)
//         }
//     }
//
//     return slice
// }
//
// func avgRating(catalog map[string]float64) float64 {
//     number := 0.0
//
//     if len(catalog) == 0 {
//         return 0.0
//     }
//
//     for _, rating := range catalog {
//         number += rating
//     }
//
//     avgRating := number / float64(len(catalog))
//
//     return avgRating
// }

func main() {
    fmt.Println()
    fmt.Println()
    fmt.Println()

//       st := []string { "apple", "banana", "apple", "orange", "banana" }
//       fmt.Println(removeDuplicates(st))

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