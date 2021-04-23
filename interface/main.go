package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int //! Contract dalam sebuah interface
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 10001
}

func main() {

	bangunDatar1 := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(bangunDatar1)

	fmt.Println("Luas persegi panjang: ", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas persegi: ", luas)

	asal := Asal{Angka: 5}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas asal: ", luas)

}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
