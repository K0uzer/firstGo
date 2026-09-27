package main
import "fmt"
// import "time"

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

//     var time = 169
//     var rate = 8.6
//     var rej = "Кростофер Нолан"
//
//     fmt.Print("Фильм: ", "Интерстеллар\n")
//     fmt.Print("Интерстеллар\n")
//     fmt.Println("Фантастика, 2014", "Фантастика, 2014")
//     fmt.Printf("Рейтинг: %.1f | %d мин |\nРежиссер: %s", rate, time, rej)

// __________________________ ошибки __________________________

// type User struct {
// 	Name  string
// 	Email string
// 	Age   int
// }
//
// type TypeValidationError = [string]string
//
// var usersCatalog = map[int]User{
// 	1: {Name: "Аня", Email: "anya@mail.com", Age: 25},
// 	2: {Name: "Борис", Email: "boris@mail.com", Age: 30},
// 	3: {Name: "Вика", Email: "vika@mail.com", Age: 28},
// }
//
// var ErrorUserNotFound = errors.New("Пользователь не найден")
// var ValidationError = TypeValidationError {
//     Field: ""
//     Value: ""
// }
//
// func validationSubscription(id int, date string, usersCatalog map[int]User) {
//     if _, ok := usersCatalog[id]; !ok   {
//         return ErrorUserNotFound
//     }
//
//     dateMask := "2006-01-02"
//
//     if !time.Parse(dateMask, date) {
//         return ValidationError
//     }
// }

// const (
//     rate = "Premium",
//     price = 999.00
// )
// var months int = 12
// var discount = 15.0
// total := (price * float64(months)) * (1 - (discount / 100))
//
// fmt.Println("---------- ЧЕК GoFlix -----------")
// fmt.Printf("Тариф: %s\nМесяцев: %d\nБазовая цена: %.2f руб/мес\nСкидка: %.0f%%\nИтого: %.2f руб.", tariff, months, price, discount, total)



// var totalSec = 8520
//
// var hours = totalSec / (60 * 60)
// var minutes = (totalSec / 60) % 60
// var seconds = totalSec % 60
//
// fmt.Printf("%d ч %d мин %d сек", hours, minutes, seconds)



// totalBite := 4_831_838_208
// totalGigoBite := float64(totalBite / 1024) / 1024 / 1024
//
// fmt.Printf("%.2f ГБ", totalGigoBite)


func ticketPrice(price float64, isEvening bool, isDayOff bool) float64 {
    if isEvening && isDayOff {
       return price + 150.0
    }
    if isDayOff {
        return price + 100.0
    }
    if isEvening {
        return price + 50.0
    }
    return price
}

func formatReceipt(title string, countTickets int, priceForTicket float64) (string, float64) {
    totalPrice := float64(countTickets) * priceForTicket
    return title, totalPrice
}

func isAllower(age int, rateFilm int) bool {
    if age > rateFilm {
        return true
    }
    return false
}

func main() {
//     fmt.Println(validationSubscription())
//     fmt.Println(validationSubscription())
//     fmt.Println(validationSubscription())
//
// fmt.Println(formatReceipt("Гарри потер", 5, 1002.5))
// fmt.Println(formatReceipt("Вперед!", 2, 150.0))
// fmt.Println(formatReceipt("Хакеры", 9, 40.5))
//
// fmt.Println(isAllower(18, 18))   // false — ровно 18 НЕ проходит (возраст > рейтинг)
// fmt.Println(isAllower(19, 18))   // true
// fmt.Println(isAllower(17, 18))   // false
// fmt.Println(isAllower(0, 0))     // false
// fmt.Println(isAllower(100, 18))  // true

fmt.Println(ticketPrice(100, true, true))
fmt.Println(ticketPrice(100, false, false))
fmt.Println(ticketPrice(100, false, true))
fmt.Println(ticketPrice(100, true, false))


}


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
