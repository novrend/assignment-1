package main

import (
	"fmt"
	"os"
)

type biodata struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	if len(os.Args) > 1 {
		inputNama := os.Args[1]
		dataSemuaTeman := getDataTeman()
		dataTeman := cariDataTeman(inputNama, dataSemuaTeman)
		if dataTeman == nil {
			dataTeman = generateData(inputNama)
		}
		printBuilder(dataTeman)
	}
}

func getDataTeman() []biodata {
	return []biodata{
		{id: 0, nama: "Ihza", alamat: "Jakarta", pekerjaan: "programmer", alasan: "wfh"},
		{id: 1, nama: "Aldi", alamat: "alamat aldi", pekerjaan: "pekerjaan aldi", alasan: "alasan aldi"},
		{id: 2, nama: "Tashya", alamat: "alamat tashya", pekerjaan: "pekerjaan tashya", alasan: "alasan tashya"},
		{id: 3, nama: "Syarif", alamat: "alamat syarif", pekerjaan: "pekerjaan syarif", alasan: "alasan syarif"},
		{id: 4, nama: "Shan", alamat: "alamat shan", pekerjaan: "pekerjaan shan", alasan: "alasan shan"},
		{id: 5, nama: "Husin", alamat: "alamat husin", pekerjaan: "pekerjaan husin", alasan: "alasan husin"},
	}
}

func generateData(nama string) *biodata {
	data := getDataTeman()
	id := len(data)
	return &biodata{
		id:        id,
		nama:      nama,
		alamat:    "alamat " + nama,
		pekerjaan: "pekerjaan " + nama,
		alasan:    "alasan " + nama,
	}
}

func cariDataTeman(nama string, data []biodata) *biodata {
	for _, value := range data {
		if value.nama == nama {
			return &value
		}
	}
	return nil
}

func printBuilder(data *biodata) {
	fmt.Printf("ID : %d\n", data.id)
	fmt.Printf("Nama : %s\n", data.nama)
	fmt.Printf("Alamat : %s\n", data.alamat)
	fmt.Printf("Pekerjaan : %s\n", data.pekerjaan)
	fmt.Printf("Alasan : %s\n", data.alasan)
}
