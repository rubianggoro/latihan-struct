package main

import "fmt"

type Sepeda struct {
	jumlahBan  int
	jumlahGigi int
	warna      string
}

func (e *Sepeda) waktuTempuh(jarak float32) float32 {
	return jarak * 2.5
}

func main() {
	arrSepeda := make(map[int]Sepeda)
	arrSepeda[0] = Sepeda{2, 4, "Merah"}
	arrSepeda[1] = Sepeda{4, 6, "Hijau"}
	arrSepeda[2] = Sepeda{1, 3, "Abu-abu"}
	arrSepeda[3] = Sepeda{2, 5, "Pink"}
	arrSepeda[4] = Sepeda{5, 6, "Biru"}

	fmt.Println(arrSepeda)

	for _, value := range arrSepeda {
		fmt.Printf("Jumlah Ban : %d\n", value.jumlahBan)
		fmt.Printf("Jumlah Gigi : %d\n", value.jumlahGigi)
		fmt.Println("warna :", value.warna)
		fmt.Println("waktu yang dibutuhkan :", value.waktuTempuh(20), "Menit")
	}
}
