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
	ID        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var biodata = []Biodata{
		{Nama: "Olivia", Alamat: "Bogor", Pekerjaan: "Buruh Tulis", Alasan: "Alasan1"},
		{Nama: "Dian", Alamat: "Medan", Pekerjaan: "QA", Alasan: "Alasan2"},
		{Nama: "Putri", Alamat: "Malang", Pekerjaan: "Developer", Alasan: "Alasan3"},
	}
	inputKey := os.Args[1:]
	input_key := inputKey[0]
	var new_data = NewData{}
	bool_var := true
	for key, value := range biodata {
		keyStr := strconv.Itoa(key)
		if input_key == keyStr || strings.ToLower(input_key) == strings.ToLower(value.Nama) {
			new_data.ID = keyStr
			new_data.Nama = value.Nama
			new_data.Alamat = value.Alamat
			new_data.Alasan = value.Alasan
			new_data.Pekerjaan = value.Pekerjaan
			fmt.Printf("%+v\n", new_data)
			bool_var = false
		}
	}
	if bool_var {
		fmt.Printf("Data dengan nama/absen tsb tidak tersedia")
	}
}
