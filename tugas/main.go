package main

import (
	"fmt"
)

type Kendaraan interface {
	testKecepatan(bensin int) int
	Tipe() string
}

// - buatlah 3 struct kendaraan ( mobil , motor , bajai)
// - masing2 struct memiliki properti kecepatan

type Mobil struct {
	Kecepatan int
}

type Motor struct {
	Kecepatan int
}

type Bajaj struct {
	Kecepatan int
}

// - buatkan function masing2 struct dengan spesifikasi :
// 	1 liter = 1 km (mobil)
// 	1 liter = 3 km (motor)
// 	1 liter = 4 km (bajaj)

func (m Mobil) testKecepatan(bensin int) int {
	return bensin * m.Kecepatan
}

func (m Motor) testKecepatan(bensin int) int {
	return bensin * m.Kecepatan
}

func (b Bajaj) testKecepatan(bensin int) int {
	return bensin * b.Kecepatan
}
func (m Mobil) Tipe() string {
	return "Mobil"
}

func (m Motor) Tipe() string {
	return "Motor"
}

func (b Bajaj) Tipe() string {
	return "Bajaj"
}

//   - buatkan satu func untuk menentukan mana yg lebih efisien dari 3 kendaraan tsb,
//     func memiliki parameter bensin dan object kendaraan (bisa langsung dimasukkan 3 object)
func KendaraanPalingEfisien(bensin int, kendaraan ...Kendaraan) string {
	var efisien Kendaraan
	terjauh := 0

	for _, k := range kendaraan {
		jarakTempuh := k.testKecepatan(bensin)
		if jarakTempuh >= terjauh {
			terjauh = jarakTempuh
			efisien = k
		}
	}
	return efisien.Tipe() + " Adalah kendaraan yang paling efisien"
}

func main() {
	mobil := Mobil{1}
	motor := Motor{3}
	bajaj := Bajaj{4}

	// - jika memiliki 10 liter bensin
	bensin := 10

	fmt.Println(KendaraanPalingEfisien(bensin, mobil, motor, bajaj))
}
