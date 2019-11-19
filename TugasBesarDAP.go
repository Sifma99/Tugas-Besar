package main

import "fmt"
import "strings"

type pelanggan struct {
	nama         string
	specMotor    motor
	jumTransaksi int
}

type service struct {
	jenisService string
	hargaService int
}
type transaksi struct {
	pembeli      pelanggan
	arrService   []service
	arrSparePart []sparePart
	dataWaktu    history
	totalHarga   int
}

type sparePart struct {
	nama         string
	harga        int
	stokTersedia int
}

type history struct {
	tglTransaksi int
	blnTransaksi int
	thnTransaksi int
}

type motor struct {
	tahunPabrikan int
	merek         string
	jenisMotor    string
	stokTersedia  int
}

// ArrSparepart Tipe untuk menampung Array tipe bentukan Sparepart.
type ArrSparepart [100]sparePart

// ArrMotor Tipe untuk menampung Array tipe bentukan motor.
type ArrMotor [100]motor

type ArrPelanggan [100]pelanggan

var dataMotor ArrMotor
var dataSparepart ArrSparepart
var dataPelanggan ArrPelanggan

func main() {
	var indeksData, countArr int
	var pilihan, choiceMenu string
	dataMotor[0].jenisMotor = "Beat"
	dataMotor[0].merek = "Yamaha"
	dataMotor[0].stokTersedia = 2
	dataMotor[0].tahunPabrikan = 2019
	dataMotor[1].jenisMotor = "CBR"
	dataMotor[1].merek = "Honda"
	dataMotor[1].stokTersedia = 1
	dataMotor[1].tahunPabrikan = 2001
	dataSparepart[0].nama = "Rem"
	dataSparepart[0].harga = 45000
	dataSparepart[0].stokTersedia = 2
	dataSparepart[1].nama = "Ban"
	dataSparepart[1].harga = 394000
	dataSparepart[1].stokTersedia = 1
	dataPelanggan[0].nama = "Rusep"
	dataPelanggan[0].specMotor.merek = "Yamaha"
	dataPelanggan[0].specMotor.jenisMotor = "cbr"
	dataPelanggan[0].specMotor.tahunPabrikan = 2018
	countArr = 2
	choiceMenu = "Yes"
	for choiceMenu == "Yes" || choiceMenu == "yes" {
		menuUtama()
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			menuJenisData()
			fmt.Scan(&pilihan)
			if pilihan == "1" {
				menuTambahEdit()
				fmt.Scan(&pilihan)
				if pilihan == "1" {
					tambahSparePart(&dataSparepart, &countArr)
				} else if pilihan == "2" {
					fmt.Print("Data yang diedit: ")
					fmt.Scan(&indeksData)
					editSparePart(indeksData, &dataSparepart)
				} else {
					fmt.Println("Maaf input tidak valid, Kembali ke menu utama.")
				}
			} else if pilihan == "2" {
				menuTambahEdit()
				fmt.Scan(&pilihan)
				if pilihan == "1" {
					nambahMotor(&dataMotor)
				} else {
					fmt.Println("Maaf input tidak valid, Kembali ke menu utama.")
				}
			} else if pilihan == "3" {
				menuTambahEdit()
				fmt.Scan(&pilihan)
				if pilihan == "1" {
					tambahPelanggan(&dataPelanggan, &countArr)
				}
			}

		} else if pilihan == "3" {
			menuJenisData()
			fmt.Scan(&pilihan)
			if pilihan == "1" {
				menuSortingSpareBy()
				fmt.Scan(&pilihan)
				if pilihan == "1" {
					sortSparePartNama(dataSparepart, countArr, &dataSparepart)
					listSparePart(dataSparepart)
				} else if pilihan == "2" {
					sortSparePartHarga(dataSparepart, countArr, &dataSparepart)
					listSparePart(dataSparepart)
				} else if pilihan == "3" {
					sortSparePartStok(dataSparepart, countArr, &dataSparepart)
					listSparePart(dataSparepart)
				}
			} else if pilihan == "2" {
				listMotor(dataMotor)
			} else if pilihan == "3" {
				listPelanggan(dataPelanggan, countArr)
			}
		} else {
			fmt.Println("Maaf input tidak valid, Kembali ke menu utama.")
		}
		fmt.Print("Kembali ke menu? (Yes/No): ")
		fmt.Scan(&choiceMenu)
		for choiceMenu != "No" && choiceMenu != "no" && choiceMenu != "Yes" && choiceMenu != "yes" {
			fmt.Println("Maaf input anda tidak valid, silahkan masukkan (Yes) atau (No)")
			fmt.Scan(&choiceMenu)

		}
	}
}

func listMotor(arr ArrMotor) {
	for i := 0; i < len(arr); i++ {
		if arr[i].tahunPabrikan != 0 && arr[i].stokTersedia != 0 && arr[i].jenisMotor != "" && arr[i].merek != "" {
			fmt.Println(arr[i].tahunPabrikan)
			fmt.Println(arr[i].merek)
			fmt.Println(arr[i].jenisMotor)
			fmt.Println(arr[i].stokTersedia)
		}
	}
}

func listSparePart(arr ArrSparepart) {
	for i := 0; i < len(arr); i++ {
		if arr[i].nama != "" && arr[i].harga != 0 && arr[i].stokTersedia != 0 {
			fmt.Printf("Spare-part: %s \n", arr[i].nama)
			fmt.Printf("Harga: %d \n", arr[i].harga)
			fmt.Printf("Stok: %d \n", arr[i].stokTersedia)
		}
	}

}

func nambahMotor(arr *ArrMotor) {
	var (
		kembali string
	)
	kembali = "yes"
	for i := 0; i <= len(arr) && kembali == "yes"; i++ {
		if arr[i].tahunPabrikan == 0 && arr[i].merek == "" && arr[i].jenisMotor == "" && arr[i].stokTersedia == 0 {
			fmt.Print("Tahun pabrikan : ")
			fmt.Scanln(&arr[i].tahunPabrikan)
			fmt.Print("Merek : ")
			fmt.Scanln(&arr[i].merek)
			fmt.Print("Jenis motor : ")
			fmt.Scanln(&arr[i].jenisMotor)
			fmt.Print("Stok tersedia : ")
			fmt.Scanln(&arr[i].stokTersedia)
			fmt.Println("Apakah kembali ke menu? ")
			fmt.Scanln(&kembali)
		}
	}
}

func tambahSparePart(arr *ArrSparepart, nArr *int) {
	var kembali string
	kembali = "yes"
	for i := 0; i < len(arr) && kembali == "yes"; i++ {
		if arr[i].nama == "" && arr[i].harga == 0 && arr[i].stokTersedia == 0 {
			fmt.Println("Silahkan masukan spare-part yang ingin ditambahkan: ")
			fmt.Scan(&arr[i].nama)
			fmt.Println("Silahkan masukan harga: ")
			fmt.Scan(&arr[i].harga)
			fmt.Println("Silahkan masukan Stok yang tersedia: ")
			fmt.Scan(&arr[i].stokTersedia)
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
		}
		*nArr = *nArr + 1
	}
}

func editSparePart(n int, arr *ArrSparepart) {
	fmt.Print("Silahkan masukan nama spare-part: ")
	fmt.Scan(&arr[n].nama)
	fmt.Print("Silahkan masukan harga: ")
	fmt.Scan(&arr[n].harga)
	fmt.Print("Silahkan masukan Stok yang tersedia: ")
	fmt.Scan(&arr[n].stokTersedia)
}

// sortSparePartHarga sorting dari harga terkecil ke harga terbesar.
func sortSparePartHarga(arr ArrSparepart, nArr int, arrOut *ArrSparepart) {
	var j int
	var temp int
	for i := 0; i < nArr; i++ {
		temp = i
		j = i + 1
		for j < nArr {
			if arr[j].harga < arr[temp].harga {
				temp = j
			}
			j++
		}
		arr[i], arr[temp] = arr[temp], arr[i]
		(*arrOut)[i] = arr[i]
		(*arrOut)[temp] = arr[temp]
	}
}

func sortSparePartStok(arr ArrSparepart, nArr int, arrOut *ArrSparepart) {
	var temp, j int
	for i := 0; i < nArr; i++ {
		temp = arr[i].stokTersedia
		j = i
		for j > 0 && arr[j-1].stokTersedia > temp {
			arr[j].stokTersedia = arr[j-1].stokTersedia
			j--
		}
		arr[j].stokTersedia = temp
	}
	*arrOut = arr
}

func sortSparePartNama(arr ArrSparepart, nArr int, arrOut *ArrSparepart) {
	var j int
	var temp int
	for i := 0; i < nArr; i++ {
		temp = i
		j = i + 1
		for j < nArr {
			if strings.ToLower(arr[j].nama) < strings.ToLower(arr[temp].nama) {
				temp = j
			}
			j++
		}
		arr[i], arr[temp] = arr[temp], arr[i]
		(*arrOut)[i] = arr[i]
		(*arrOut)[temp] = arr[temp]
	}
}

// Misalkan Data sudah di sort
func searchSparePart(arr ArrSparepart, key string, nArr int) bool {
	var awal, tengah, akhir int
	var ketemu bool
	awal = 0
	akhir = nArr - 1
	tengah = (awal + akhir) / 2
	ketemu = false
	for awal < akhir && arr[tengah].nama != key {
		if arr[tengah].nama > key {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
	}
	if arr[tengah].nama == key {
		ketemu = true
	}
	return ketemu
}

func tambahPelanggan(arr *ArrPelanggan, nArr *int) {
	var kembali string
	kembali = "yes"
	for i := 0; i < len(arr) && kembali == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Println("Silahkan masukan nama pelanggan: ")
			fmt.Scan(&arr[i].nama)
			fmt.Println("Silahkan masukan Spesifikasi Motor ")
			fmt.Println("Silahkan masukan merek motor: ")
			fmt.Scan(&arr[i].specMotor.merek)
			fmt.Println("Silahkan masukan Tahun Pabrikan Motor: ")
			fmt.Scan(&arr[i].specMotor.tahunPabrikan)
			fmt.Println("Silahkan masukan Jenis motor: ")
			fmt.Scan(&arr[i].specMotor.jenisMotor)
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
		}
		*nArr = *nArr + 1
	}
}

func listPelanggan(arr ArrPelanggan, nArr int) {
	for i := 0; i < nArr; i++ {
		if arr[i].nama != "" {
			fmt.Printf("Nama: %s \n", arr[i].nama)
			fmt.Printf("Merek Motor: %s \n", arr[i].specMotor.merek)
			fmt.Printf("Jenis Motor: %s \n", arr[i].specMotor.jenisMotor)
			fmt.Printf("Tahun Pabrikan Motor: %d \n", arr[i].specMotor.tahunPabrikan)
		}
	}
}
func menuUtama() {
	fmt.Println("Selamat datang di bengkel onlen")
	fmt.Println("Silahkan pilih menu: ")
	fmt.Println("1. Edit Data")
	fmt.Println("2. Transaksi")
	fmt.Println("3. Lihat Data")
	fmt.Print("Pilihan: ")
}

func menuJenisData() {
	fmt.Println("Silahkan pilih jenis data: ")
	fmt.Println("1. Spare-part")
	fmt.Println("2. Motor")
	fmt.Println("3. Pelanggan")
	fmt.Print("Pilihan: ")
}

func menuTambahEdit() {
	fmt.Println("Silahkan pilih Edit / Tambah data: ")
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Print("Pilihan: ")
}

func menuSortingSpareBy() {
	fmt.Println("Silahkan pilih Sorting berdasarkan: ")
	fmt.Println("1. Nama")
	fmt.Println("2. Harga")
	fmt.Println("3. Stok")
	fmt.Print("Pilihan: ")

}
