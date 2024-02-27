package main

import (
	"fmt"
	"os"
)

type Data struct {
	name    string
	address string
	job     string
	reason  string
}

var data = map[int]Data{
	1: {"Ilham", "Subang", "Mahasiswa", "Ingin mempelajari bahasa pemrograman baru"},
	2: {"Maulana", "Subang", "Mahasiswa", "Ingin mencoba hal baru"},
	3: {"Ilyasa", "Subang", "Mahasiswa", "Coba - coba"},
}

func showData(absen int) {
	teman, ok := data[absen]
	if !ok {
		fmt.Println("Nomor absen tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Nama: ", teman.name)
	fmt.Println("Alamat: ", teman.address)
	fmt.Println("Pekerjaan: ", teman.job)
	fmt.Println("Alasan memilih kelas Golang: ", teman.reason)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor absen]")
		fmt.Println("Contoh : go run biodata.go 1")
		return
	}

	// Get argument
	absen := os.Args[1]

	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	showData(absenInt)
}
