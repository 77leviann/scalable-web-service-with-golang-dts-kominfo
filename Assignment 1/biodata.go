package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataTeman = []Teman{
	{1, "El", "Jakarta", "Software Engineer", "Suka bahasa pemrograman Go."},
	{2, "Ann", "Bandung", "Digital Marketing", "Ingin switch career."},
	{3, "Nur", "Surabaya", "Web Developer", "Mendapat rekomendasi dari teman untuk belajar Go."},
	{4, "Bryan", "Lampung", "Front End Web Dev", "Ingin belajar Back End dengan Go."},
	{5, "Adel", "Bali", "Back End Developer", "Ingin switch dari Java ke Go."},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go [nomor_absen]")
		return
	}

	nomorAbsenStr := os.Args[1]
	nomorAbsen, err := strconv.Atoi(nomorAbsenStr)
	if err != nil {
		fmt.Println("Error: input tidak valid")
		return
	}

	teman, err := getTemanByAbsen(nomorAbsen)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}

func getTemanByAbsen(nomor int) (Teman, error) {
	for _, teman := range dataTeman {
		if teman.Absen == nomor {
			return teman, nil
		}
	}
	return Teman{}, fmt.Errorf("tidak ada teman dengan nomor absen %d", nomor)
}
