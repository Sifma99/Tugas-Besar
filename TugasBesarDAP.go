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

// ArrPelanggan Tipe untuk menampung array tipe bentukan pelanggan.
type ArrPelanggan [100]pelanggan

var dataMotor ArrMotor
var dataSparepart ArrSparepart
var dataPelanggan ArrPelanggan

func main() {
	var countArrSparepart, countArrPelanggan, countArrMotor int
	var pilihan, choiceMenu string
	choiceMenu = "Yes"
	for strings.ToLower(choiceMenu) == "yes" {
		menuUtama()
		switch pilihan {
		case "1":
			menuTambahEdit()
			fmt.Scan(&pilihan)
			switch pilihan {
			case "1":
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					tambahSparePart(&dataSparepart, &countArrSparepart)
				case "2":
					nambahMotor(&dataMotor, &countArrMotor)
				case "3":
					tambahPelanggan(&dataPelanggan, &countArrPelanggan)
				}
			case "2":
			case "3":
			}
		case "2":
		case "3":
			menuJenisData()
			fmt.Scan(&pilihan)
			switch pilihan {
			case "1":
				menuSortingSpareBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					sortSparePartNama(dataSparepart, countArrSparepart, &dataSparepart)
				case "2":
					sortSparePartHarga(dataSparepart, countArrSparepart, &dataSparepart)
				case "3":
					sortSparePartStok(dataSparepart, countArrSparepart, &dataSparepart)
				}
				listSparePart(dataSparepart)
			}

		}
		fmt.Println("Kembali ke menu ? (Yes / No): ")
		fmt.Scan(&choiceMenu)
		for strings.ToLower(choiceMenu) != "no" && strings.ToLower(choiceMenu) != "yes" {
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

func nambahMotor(arr *ArrMotor, nArr *int) {
	var (
		kembali string
	)
	kembali = "yes"
	for i := 0; i <= len(arr) && kembali == "yes"; i++ {
		if arr[i].tahunPabrikan == 0 && arr[i].merek == "" && arr[i].jenisMotor == "" && arr[i].stokTersedia == 0 {
			fmt.Print("Tahun pabrikan : ")
			fmt.Scan(&arr[i].tahunPabrikan)
			fmt.Print("Merek : ")
			fmt.Scan(&arr[i].merek)
			fmt.Print("Jenis motor : ")
			fmt.Scan(&arr[i].jenisMotor)
			fmt.Print("Stok tersedia : ")
			fmt.Scan(&arr[i].stokTersedia)
			fmt.Println("Apakah kembali ke menu? ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
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
	for i := 0; i < *nArr && kembali == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Println("=================================================")
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
			fmt.Println("========================================")
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
