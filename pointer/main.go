package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

type Gamer struct {
	Name  string
	Games []string
}

// func (student *Student) graduate() { //! Pointer struct sebagai parameter
// 	student.Name = student.Name + " S.T"
// }

func main() {
	// numberA := 10
	// numberB := &numberA //! Menyimpan alamat memori

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 30

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// var numberA int = 5
	// var numberB *int = &numberA //! & = Referencing / Reference alamat memori

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB) //! Proses direferencing atau direference, mengembalikan value sebenarnya

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	// number := 5
	// fmt.Println("alamat memori: ", &number)
	// fmt.Println("nilai awal: ", number)
	// change(&number, 100)
	// fmt.Println("nilai akhir: ", number)
	// fmt.Println("alamat memori: ", &number)

	// student := Student{1, "titan", 3.72}

	// fmt.Println(student.Name)
	// student.graduate()
	// fmt.Println(student.Name)

	gamer := Gamer{Name: "Titan"}

	gamer.Addgame("Fifa")
	gamer.Addgame("God of Wars")
	gamer.Addgame("Mario")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}

func (gamer *Gamer) Addgame(game string) {
	gamer.Games = append(gamer.Games, game)
}

// func change(old *int, new int) {
// 	*old = new
// 	fmt.Println("Alamat memory: ", old)
// 	fmt.Println("Di dalam function: ", *old)
// }
