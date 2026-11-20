package main

import (
	"fmt"
)

func main() {
	var namaMahasiswa string = "Jaka Kurnia"
	var umurMahasiswa int = 21
	var semesterMahasiswa int = 4
	var isActive bool = true

	ipk := 4.0
	kampus := "Institut Teknologi Bandung"

	fmt.Println("--- Data Mahasiswa ---")
	fmt.Println("Nama :", namaMahasiswa)
	fmt.Println("Umur :", umurMahasiswa)
	fmt.Println("Semester :", semesterMahasiswa)
	fmt.Println("IPK :", ipk)
	fmt.Println("Kampus :", kampus)
	fmt.Println("Status Aktif :", isActive)

	fmt.Printf("\nHI, %s kamu semester %d dengan IPK %.2f.\n", namaMahasiswa, semesterMahasiswa, ipk)

	// var programStudi string = "Teknik Informatika"
	// var totalSks int = 144
	// var lulusSls bool = false

	// if lulusSls {
	// 	fmt.Printf("Selamat, kamu lulus program studi %s dengan total SKS %d.\n", programStudi, totalSks)
	// } else {
	// 	fmt.Printf("Maaf, kamu belum lulus program studi %s dengan total SKS %d.\n", programStudi, totalSks)
	// }

	// switch lulusSls {
	// case true:
	// 	fmt.Printf("Selamat, kamu lulus program studi %s dengan total SKS %d.\n", programStudi, totalSks)
	// case false:
	// 	fmt.Printf("Maaf, kamu belum lulus program studi %s dengan total SKS %d.\n", programStudi, totalSks)
	// }

	// resaultTest := 70

	// if resaultTest >= 85{
	// 	fmt.Println("A")
	// }else if resaultTest >= 70{
	// 	fmt.Println("B")
	// }else{
	// 	fmt.Println("Perlu belajar lagi")
	// }

}
