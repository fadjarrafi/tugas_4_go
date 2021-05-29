package main

import "fmt"

func main() {
	mahasiswa := make(map[string]int)

	tampilMahasiswa(mahasiswa)
	for key, val := range mahasiswa {
		fmt.Println(key, ":", val, "cm")
	}
}

func tampilMahasiswa(mahasiswa map[string]int) {

	mahasiswa["Aldo"] = 182
	mahasiswa["Yosep"] = 178

}
