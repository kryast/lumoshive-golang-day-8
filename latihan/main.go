// buatkan struct , lakukan inisialisasi struct tsb , kemudian ditampilkan nilai struct.

package main

import "fmt"

type Data struct {
	Nama, Kota, Negara string
}

func main() {
	pegawai := Data{"Ahmad", "Palembang", "Indonesia"}
	fmt.Println(pegawai)
	fmt.Println(pegawai.Nama)

}
