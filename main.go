package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const PORT int = 1782 //ubdah port sesuai dengan NPM kalian masing - masing. contoh: jika npm 54422423 maka portnya 22423

func main() {
	http.HandleFunc("/api/mahasiswa", user)                         // buat route untuk api/mahasiswa
	http.Handle("/", http.FileServer(http.Dir("polymer_51422292"))) // buat route untuk file statis
	fmt.Printf("web berjalan pada Port %d. url adalah http://[::1]:%d/", PORT, PORT)
	fmt.Println()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil) // jalankan server web
}

//TODO: buatlah struct lepkom dengan atribut Nama, Kursus, dan Foto
type lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

//TODO: buatlah slice lepkom dengan data 3 mahasiswa
var dataMahasiswa = []lepkom{
	{
		Nama:   "Pradipa Adicandra",
		Kursus: "Fundamental web",
		Foto:   "img/gambar1.jpg",
	},
	{
		Nama:   "Pradipa Adicandra",
		Kursus: "Golang for Beginner",
		Foto:   "img/gambar2.png",
	},
	{
		Nama:   "Pradipa Adicandra",
		Kursus: "Golang for Intermediate",
		Foto:   "img/gambar3.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header content-type menjadi json
	if r.Method == http.MethodGet {

		result, err := json.Marshal(dataMahasiswa) // ubah variabel data_mahasiswa menjadi json

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // jika terjadi error, tampilkan error 500
			return
		}

		w.Write(result) // tampilkan data json
		return
	}
}
