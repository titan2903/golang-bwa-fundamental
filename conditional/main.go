package main

import "fmt"

func main() {
	// ! CONDITION AND FOR LOOP
	// age := 10

	// if age > 10 {
	// 	fmt.Println("main game GTA")
	// } else {
	// 	fmt.Println("kamu blm boleh")
	// }

	// score := 60
	// var grade string

	// if score <= 60 {
	// 	grade = "D"
	// } else if score <= 70 {
	// 	grade = "C"
	// } else if score <= 80 {
	// 	grade = "B"
	// } else {
	// 	grade = "No Grade"
	// }

	// fmt.Println(grade)

	// number := 5
	// switch number {
	// case 1:
	// 	fmt.Println("satu")
	// case 2:
	// 	fmt.Println("dua")
	// case 3:
	// 	fmt.Println("tiga")
	// default:
	// 	fmt.Println("tidak ada")
	// }

	// for i := 1; i <= 100; i += 1 {
	// 	fmt.Println("saya sedang belajar go :", i)
	// }
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("saya sedang belajar go :", i)
	// 	i++
	// }

	// title := "Golang the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index: ", index, " latter: ", string(letter))
	// 	}
	// }

	// for index, letter := range title {
	// 	letterString := string(letter)

	// 	if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
	// 		fmt.Println("index: ", index, " latter: ", string(letter))
	// 	}

	// 	switch letterString {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println("index: ", index, " latter: ", string(letter))
	// 	}
	// }

	// ! ARRAY

	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Py"
	// languages[2] = "Ruby"
	// languages[3] = "JS"
	// languages[4] = "C#"

	// languages := [...]string{
	// 	"Ruby",
	// 	"go",
	// 	"python",
	// 	"js",
	// 	"c#",
	// 	"kotlin",
	// }
	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	// for index, lang := range languages {
	// 	fmt.Println("Index: ", index, "languages: ", lang)
	// }

	// ! SLICE
	// var gamingConsoles []string
	// gamingConsoles = append(gamingConsoles, "ps4")
	// gamingConsoles = append(gamingConsoles, "ps3")
	// gamingConsoles = append(gamingConsoles, "ps5")
	// gamingConsoles = append(gamingConsoles, "ps2")
	// gamingConsoles = append(gamingConsoles, "ps1")
	// // fmt.Println(gamingConsoles)

	// for _, console := range gamingConsoles {
	// 	fmt.Println(console)
	// }

	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["GO"] = 8

	// fmt.Println(myMap["Ruby"])

	// myMap := map[string]string{"ruby": "is oke", "go": "is fast", "javascript": "hype"}
	// fmt.Println(myMap)
	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	// delete(myMap, "ruby")

	// fmt.Println("=======\n")

	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }
	// value := myMap["python"]
	// fmt.Println(value)

	// value, isAvailable := myMap["python"]
	// fmt.Println(isAvailable)
	// fmt.Println(value)
	// students := []map[string]string{
	// 	{"name": "titan", "skore": "a"},
	// 	{"name": "budi", "skore": "b"},
	// 	{"name": "eko", "skore": "c"},
	// }

	// for _, student := range students {
	// 	fmt.Println(student["name"], " = ", student["skore"])
	// }
	// fmt.Println(students)

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// var total int
	// for _, score := range scores {
	// 	total += score
	// }

	// length := len(scores)
	// average := float64(total) / float64(length)

	// fmt.Println(average)

	var godScores []int

	for _, score := range scores {
		if score >= 90 {
			godScores = append(godScores, score)
		}
	}

	fmt.Println(godScores)
}
