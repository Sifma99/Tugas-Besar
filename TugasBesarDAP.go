package main

import (
	"fmt"
)

type pelanggan struct {
	nama         string
	specMotor    motor
	dataWaktu    history
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

type ArrSparepart [100]sparePart
type ArrMotor [100]motor

var dataMotor ArrMotor
var dataSparepart ArrSparepart

func main() {
	var indeksData int
	var pilihan, choiceMenu string
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
					tambahSparePart(&dataSparepart)
				} else if pilihan == "2" {
					fmt.Print("Data yang diedit: ")
					fmt.Scan(&indeksData)
					editSparePart(indeksData, &dataSparepart)
				}
			} else if pilihan == "2" {
				menuTambahEdit()
				fmt.Scan(&pilihan)
				if pilihan == "1" {
					nambahMotor(&dataMotor)
				}

			}

		} else if pilihan == "3" {
			menuJenisData()
			fmt.Scan(&pilihan)
			if pilihan == "1" {
				listSparePart(dataSparepart)
			} else if pilihan == "2" {
				listMotor(dataMotor)
			}
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

func tambahSparePart(arr *ArrSparepart) {
	var kembali string
	kembali = "yes"
	for i := 0; i < len(arr) && kembali == "yes"; i++ {
		fmt.Println("Silahkan masukan spare-part yang ingin ditambahkan: ")
		fmt.Scan(&arr[i].nama)
		fmt.Println("Silahkan masukan harga: ")
		fmt.Scan(&arr[i].harga)
		fmt.Println("Silahkan masukan Stok yang tersedia: ")
		fmt.Scan(&arr[i].stokTersedia)
		fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
		fmt.Scan(&kembali)
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
	fmt.Print("Pilihan: ")
}

func menuTambahEdit() {
	fmt.Println("Silahkan pilih Edit / Tambah data: ")
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Print("Pilihan: ")

}
