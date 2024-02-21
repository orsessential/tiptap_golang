package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type NewData struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func isEqual(key1, key2 string) bool {
	return strings.EqualFold(key1, key2)
}

func main() {
	var biodata = []Biodata{
		{Nama: "Olivia", Alamat: "Bogor", Pekerjaan: "Buruh Tulis", Alasan: "Alasan1"},
		{Nama: "Dian", Alamat: "Medan", Pekerjaan: "QA", Alasan: "Alasan2"},
		{Nama: "Putri", Alamat: "Malang", Pekerjaan: "Developer", Alasan: "Alasan3"},
	}
	var new_data = NewData{}
	bool_var := true
	if len(os.Args[1:]) == 0 {
		bool_var = false
		fmt.Println("Tolong Masukkan Nama atau Nomor Absen.\nContoh: go run main.go Olivia atau go run main.go 3")
	} else {
		inputKey := os.Args[1:]
		input_key := inputKey[0]
		for key, value := range biodata {
			num, err := strconv.Atoi(input_key)
			if (num == key && err == nil) || isEqual(input_key, value.Nama) {
				new_data.ID = key
				new_data.Nama = value.Nama
				new_data.Alamat = value.Alamat
				new_data.Alasan = value.Alasan
				new_data.Pekerjaan = value.Pekerjaan
				bool_var = false

				fmt.Printf("ID: %d\n", new_data.ID)
				fmt.Printf("Nama: %s\n", new_data.Nama)
				fmt.Printf("Alamat: %s\n", new_data.Alamat)
				fmt.Printf("Pekerjaan: %s\n", new_data.Pekerjaan)
				fmt.Printf("Alasan: %s\n", new_data.Alasan)

				new_data = NewData{}
			}
		}
	}
	if bool_var {
		fmt.Printf("Data dengan nama/absen tsb tidak tersedia\n")
	}
}
