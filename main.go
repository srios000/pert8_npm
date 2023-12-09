package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const PORT = 28080 //ubdah port sesuai dengan NPM kalian masing - masing. contoh: jika npm 54422423 maka portnya 22423

func main() {
	http.HandleFunc("/api/mahasiswa", user) // buat route untuk api/mahasiswa
	http.Handle("/", http.FileServer(http.Dir("polymer_npm"))) // buat route untuk file statis
	fmt.Printf("web berjalan pada Port %d. url adalah http://[::1]:%d/", PORT, PORT)
	fmt.Println()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil) // jalankan server web
}

//TODO: buatlah struct lepkom dengan atribut Nama, Kursus, dan Foto

//TODO: buatlah slice lepkom dengan data 3 mahasiswa

func user(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-type", "application/json") // set header content-type menjadi json
	if r.Method == http.MethodGet {

		result, err := json.Marshal(data_mahasiswa) // ubah variabel data_mahasiswa menjadi json

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // jika terjadi error, tampilkan error 500
			return
		}

		w.Write(result)	// tampilkan data json
		return
	}
}
