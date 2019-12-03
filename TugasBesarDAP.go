package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Pelanggan Tipe bentukan pelanggan yang berisi nama, specMotor, jumTransaksi.
type Pelanggan struct {
	nama         string
	specMotor    Motor
	jumTransaksi []Transaksi
	countTrans   int
	totalHarga   int
}

// Service Tipe bentukan service yang berisi nama, hargaService.
type Service struct {
	nama         string
	hargaService int
}

// Transaksi Tipe bentukan Transaksi yang berisi namaTrans, hargaTrans.
type Transaksi struct {
	namaTrans  string
	hargaTrans int
}

// SparePart Tipe bentukan SparePart yang berisi nama, harga, stokTersedia.
type SparePart struct {
	nama         string
	harga        int
	stokTersedia int
}

// History Tipe bentukan history yang berisi tglTransaksi, blnTransaksi, thnTransaksi.
type History struct {
	tglTransaksi int
	blnTransaksi int
	thnTransaksi int
}

// Motor Tipe bentukan motor yang berisi tahunPabrikan, merek, jenisMotr, stokTersedia.
type Motor struct {
	tahunPabrikan int
	merek         string
	jenisMotor    string
	stokTersedia  int
}

// IsiArray Panjang Array
const IsiArray = 1000

// ArrSparepart Tipe untuk menampung Array tipe bentukan Sparepart.
type ArrSparepart [IsiArray]SparePart

// ArrMotor Tipe untuk menampung Array tipe bentukan motor.
type ArrMotor [IsiArray]Motor

// ArrPelanggan Tipe untuk menampung array tipe bentukan pelanggan.
type ArrPelanggan [IsiArray]Pelanggan

// ArrService Tipe untuk menampung array tipe bentukan Service.
type ArrService [IsiArray]Service

var inputreader = bufio.NewReader(os.Stdin)

func main() {
	var countArrSparepart, countArrPelanggan, countArrMotor, countArrService int
	var idxSpare, idxMotor, idxPelanggan, idxService int
	var pilihan, choiceMenu, spareDicari, motorDicari, pelangganDicari, serviceDicari string
	var dataMotor ArrMotor
	var dataSparepart ArrSparepart
	var dataPelanggan ArrPelanggan
	var dataService ArrService

	//var pendapatanTotal int
	choiceMenu = "Yes"
	for strings.ToLower(choiceMenu) == "yes" {
		menuUtama()
		fmt.Scan(&pilihan)
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
				case "4":
					tambahService(&dataService, &countArrService)
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
			case "2":
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					fmt.Print("Sparepart yang ingin diedit: ")
					fmt.Scan(&spareDicari)
					idxSpare = searchSparePart(dataSparepart, spareDicari, countArrSparepart)
					if idxSpare != -1 {
						editSparePart(dataSparepart, idxSpare, &dataSparepart)
					} else {
						fmt.Printf("Maaf tidak terdapat %s.", spareDicari)
						fmt.Println("")
					}
				case "2":
					fmt.Print("Motor yang ingin diedit: ")
					fmt.Scan(&motorDicari)
					idxMotor = searchMotor(dataMotor, motorDicari, countArrMotor)
					if idxMotor != -1 {
						editMotor(dataMotor, idxMotor, &dataMotor)
					} else {
						fmt.Printf("Maaf tidak terdapat %s.\n", motorDicari)
					}
				case "3":
					fmt.Print("Pelanggan yang ingin diedit: ")
					fmt.Scan(&pelangganDicari)
					idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
					if idxPelanggan != -1 {
						editPelanggan(dataPelanggan, idxPelanggan, &dataPelanggan)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", pelangganDicari)
					}
				case "4":
					fmt.Print("Jenis service yang ingin diedit: ")
					fmt.Scan(&serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						editService(dataService, idxSpare, &dataService)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", serviceDicari)
					}
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
			case "3":
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					fmt.Print("Sparepart yang ingin dihapus: ")
					fmt.Scan(&spareDicari)
					idxSpare = searchSparePart(dataSparepart, spareDicari, countArrSparepart)
					if idxSpare != -1 {
						deleteSparepart(dataSparepart, idxSpare, &dataSparepart)
					}
				case "2":
					fmt.Print("Motor yang ingin dihapus: ")
					fmt.Scan(&motorDicari)
					idxMotor = searchMotor(dataMotor, motorDicari, countArrMotor)
					if idxMotor != -1 {
						deleteMotor(dataMotor, idxMotor, &dataMotor)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", motorDicari)
					}
				case "3":
					fmt.Print("Pelanggan yang ingin dihapus: ")
					fmt.Scan(&pelangganDicari)
					idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
					if idxPelanggan != -1 {
						deletePelanggan(dataPelanggan, idxPelanggan, &dataPelanggan)
					} else {
						fmt.Printf("Maaf idak terdapat %s.\n", pelangganDicari)
					}
				case "4":
					fmt.Print("Service yang ingin dihapus: ")
					fmt.Scan(&serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						deleteService(dataService, idxService, &dataService)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", serviceDicari)
					}
				}
			default:
				fmt.Println("Maaf pilihan tidak valid.")
			}
		case "2":
			menuTransaksi()
			fmt.Scan(&pilihan)
			fmt.Println("Silahkan masukan nama pelanggan: ")
			fmt.Scan(&pelangganDicari)
			idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
			if idxPelanggan != 1 {
				switch pilihan {
				case "1":
					fmt.Println("Silahkan masukkan jenis service yang ingin dilakukan: ")
					sortServiceHarga(dataService, countArrService, &dataService)
					listService(dataService, countArrService)
					fmt.Print("Pilihan: ")
					fmt.Scan(&serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						tambahTransService(&dataPelanggan, dataService, idxService, idxPelanggan)
					}
				}
			} else {
				fmt.Printf("Maaf tidak terdapat %s. \n", pelangganDicari)
			}

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
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
				listSparePart(dataSparepart, countArrSparepart)
			case "2":
				menuSortingMotorBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					sortMotorMerek(dataMotor, countArrMotor, &dataMotor)
				case "2":
					sortMotorTahun(dataMotor, countArrMotor, &dataMotor)
				case "3":
					sortMotorStok(dataMotor, countArrMotor, &dataMotor)
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
				listMotor(dataMotor, countArrMotor)
			case "3":
				menuSortingPelangganBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					sortPelangganNama(dataPelanggan, countArrPelanggan, &dataPelanggan)
				case "2":
					sortPelangganCount(dataPelanggan, countArrPelanggan, &dataPelanggan)
				}
				listPelanggan(dataPelanggan, countArrPelanggan)
			case "4":
				menuSortingServiceBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					sortServiceNama(dataService, countArrService, &dataService)
				case "2":
					sortServiceHarga(dataService, countArrService, &dataService)
				}
				listService(dataService, countArrService)
			default:
				fmt.Println("Maaf pilihan tidak valid.")
			}
		default:
			fmt.Println("Maaf pilihan tidak valid.")
		}
		fmt.Println("Kembali ke menu ? (Yes / No): ")
		fmt.Scan(&choiceMenu)
		for strings.ToLower(choiceMenu) != "no" && strings.ToLower(choiceMenu) != "yes" {
			fmt.Println("Maaf input anda tidak valid, silahkan masukkan (Yes) atau (No)")
			fmt.Scan(&choiceMenu)
		}
	}
}

func nambahMotor(arr *ArrMotor, nArr *int) {
	var (
		kembali string
	)
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].tahunPabrikan == 0 && arr[i].merek == "" && arr[i].jenisMotor == "" && arr[i].stokTersedia == 0 {
			fmt.Print("Tahun pabrikan : ")
			fmt.Scan(&arr[i].tahunPabrikan)
			fmt.Print("Merek : ")
			fmt.Scanln()
			(*arr)[i].merek, _ = inputreader.ReadString('\n')
			fmt.Print("Jenis motor : ")
			(*arr)[i].jenisMotor, _ = inputreader.ReadString('\n')
			fmt.Scanln()
			fmt.Print("Stok tersedia : ")
			fmt.Scan(&arr[i].stokTersedia)
			fmt.Println("Apakah kembali ke menu? ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahService(arr *ArrService, nArr *int) {
	var kembali string
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Println("==================================================")
			fmt.Println("Silahkan masukan nama Service yang ingin ditambahkan: ")
			fmt.Scanln()
			(*arr)[i].nama, _ = inputreader.ReadString('\n')
			fmt.Println("Silahkan masukan tarif Service: ")
			fmt.Scan(&arr[i].hargaService)
			fmt.Println("Apakah ingin menambahkan lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahSparePart(arr *ArrSparepart, nArr *int) {
	var kembali string
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].nama == "" && arr[i].harga == 0 && arr[i].stokTersedia == 0 {
			fmt.Println("===============================================================")
			fmt.Println("Silahkan masukan spare-part yang ingin ditambahkan: ")
			fmt.Scanln()
			(*arr)[i].nama, _ = inputreader.ReadString('\n')
			fmt.Println("Silahkan masukan harga: ")
			fmt.Scan(&arr[i].harga)
			fmt.Println("Silahkan masukan Stok yang tersedia: ")
			fmt.Scan(&arr[i].stokTersedia)
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}

	}
}

func tambahPelanggan(arr *ArrPelanggan, nArr *int) {
	var kembali string
	kembali = "yes"
	for i := 0; i < IsiArray && kembali == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Println("=================================================")
			fmt.Println("Silahkan masukan nama pelanggan: ")
			fmt.Scanln()
			input, _ := inputreader.ReadString('\n')
			(*arr)[i].nama = input
			fmt.Println("Silahkan masukan Spesifikasi Motor ")
			fmt.Println("Silahkan masukan merek motor: ")
			fmt.Scan(&arr[i].specMotor.merek)
			fmt.Println("Silahkan masukan Tahun Pabrikan Motor: ")
			fmt.Scan(&arr[i].specMotor.tahunPabrikan)
			fmt.Println("Silahkan masukan Jenis motor: ")
			fmt.Scan(&arr[i].specMotor.jenisMotor)
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahTransSpare(arr *ArrPelanggan, arrData *ArrSparepart, indexSpare, indexPelanggan int) {
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].namaTrans = arrData[indexSpare].nama
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].hargaTrans = arrData[indexSpare].harga
	(*arrData)[indexSpare].stokTersedia = (*arrData)[indexSpare].stokTersedia - 1
	(*arr)[indexPelanggan].countTrans = (*arr)[indexPelanggan].countTrans + 1

}

func tambahTransService(arr *ArrPelanggan, arrData ArrService, indexService, indexPelanggan int) {
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].namaTrans = arrData[indexService].nama
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].hargaTrans = arrData[indexService].hargaService
	(*arr)[indexPelanggan].countTrans = (*arr)[indexPelanggan].countTrans + 1
}

func editService(arr ArrService, n int, arrOut *ArrService) {
	fmt.Print("Silahkan masukan nama Service: ")
	fmt.Scan(&arr[n].nama)
	fmt.Print("Silahkan masukan Tarif Service: ")
	fmt.Scan(&arr[n].hargaService)
	(*arrOut)[n] = arr[n]
}

func editMotor(arr ArrMotor, n int, arrOut *ArrMotor) {
	fmt.Print("Silahkan masukan tahun: ")
	fmt.Scan(&arr[n].tahunPabrikan)
	fmt.Print("Silahkan masukan merek motor: ")
	fmt.Scan(&arr[n].merek)
	fmt.Print("Silahkan masukan jenis motor: ")
	fmt.Scan(&arr[n].jenisMotor)
	fmt.Print("Silahkan masukan Stok yang tersedia: ")
	fmt.Scan(&arr[n].stokTersedia)
	(*arrOut)[n] = arr[n]
}

func editSparePart(arr ArrSparepart, n int, arrOut *ArrSparepart) {
	fmt.Print("Silahkan masukan nama spare-part: ")
	fmt.Scan(&arr[n].nama)
	fmt.Print("Silahkan masukan harga: ")
	fmt.Scan(&arr[n].harga)
	fmt.Print("Silahkan masukan Stok yang tersedia: ")
	fmt.Scan(&arr[n].stokTersedia)
	(*arrOut)[n] = arr[n]
}

func editPelanggan(arr ArrPelanggan, n int, arrOut *ArrPelanggan) {
	fmt.Print("Silahkan masukan nama Pelanggan: ")
	fmt.Scanln()
	input, _ := inputreader.ReadString('\n')
	arr[n].nama = input
	fmt.Println("Silahkan masukan Spesifikasi motor")
	fmt.Print("Silahkan masukan merek motor: ")
	fmt.Scanln()
	arr[n].specMotor.merek, _ = inputreader.ReadString('\n')
	fmt.Print("Silahkan Masukan tahun pabrikan motor: ")
	fmt.Scan(&arr[n].specMotor.tahunPabrikan)
	fmt.Println("Silahkan masukan jenis motor: ")
	fmt.Scanln()
	arr[n].specMotor.jenisMotor, _ = inputreader.ReadString('\n')
	(*arrOut)[n] = arr[n]
}

func deleteService(arr ArrService, indeks int, arrOut *ArrService) {
	arr[indeks].nama = ""
	arr[indeks].hargaService = 0
	(*arrOut)[indeks] = arr[indeks]
}

func deletePelanggan(arr ArrPelanggan, indeks int, arrOut *ArrPelanggan) {
	arr[indeks].nama = ""
	arr[indeks].specMotor.jenisMotor = ""
	arr[indeks].specMotor.merek = ""
	arr[indeks].specMotor.stokTersedia = 0
	arr[indeks].specMotor.tahunPabrikan = 0
	(*arrOut)[indeks] = arr[indeks]
}

func deleteMotor(arr ArrMotor, indeks int, arrOut *ArrMotor) {
	arr[indeks].jenisMotor = ""
	arr[indeks].merek = ""
	arr[indeks].stokTersedia = 0
	arr[indeks].tahunPabrikan = 0
}

func deleteSparepart(arr ArrSparepart, indeks int, arrOut *ArrSparepart) {
	arr[indeks].harga = 0
	arr[indeks].nama = ""
	arr[indeks].stokTersedia = 0
	(*arrOut)[indeks] = arr[indeks]
}

// sortMotor sorting dari harga terkecil ke harga terbesar.
func sortMotorMerek(arr ArrMotor, nArr int, arrOut *ArrMotor) {
	for i := 0; i < nArr; i++ {
		for j := i + 1; j > 0 && strings.ToLower(arr[j].merek) < strings.ToLower(arr[j-1].merek); j-- {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
		}
	}
	*arrOut = arr
}

func sortMotorTahun(arr ArrMotor, nArr int, arrOut *ArrMotor) {
	for i := 0; i < nArr; i++ {
		for j := i + 1; j > 0 && arr[j].tahunPabrikan < arr[j-1].tahunPabrikan; j-- {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
		}
	}
	*arrOut = arr
}

func sortMotorStok(arr ArrMotor, nArr int, arrOut *ArrMotor) {
	for i := 0; i < nArr; i++ {
		for j := i + 1; j > 0 && arr[j].stokTersedia < arr[j-1].stokTersedia; j-- {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
		}
	}
	*arrOut = arr
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
	for i := 0; i < nArr-1; i++ {
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

func sortServiceNama(arr ArrService, nArr int, arrOut *ArrService) {
	var j int
	var temp int
	for i := 0; i < nArr-1; i++ {
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

func sortServiceHarga(arr ArrService, nArr int, arrOut *ArrService) {
	var j int
	var temp int
	for i := 0; i < nArr; i++ {
		temp = i
		j = i + 1
		for j < nArr {
			if arr[j].hargaService < arr[temp].hargaService {
				temp = j
			}
			j++
		}
		arr[i], arr[temp] = arr[temp], arr[i]
		(*arrOut)[i] = arr[i]
		(*arrOut)[temp] = arr[temp]
	}
}

func sortPelangganNama(arr ArrPelanggan, nArr int, arrOut *ArrPelanggan) {
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

func sortPelangganCount(arr ArrPelanggan, nArr int, arrOut *ArrPelanggan) {
	var j int
	var temp int
	for i := 0; i < nArr; i++ {
		temp = i
		j = i + 1
		for j < nArr {
			if arr[j].countTrans < arr[temp].countTrans {
				temp = j
			}
			j++
		}
		arr[i], arr[temp] = arr[temp], arr[i]
		(*arrOut)[i] = arr[i]
		(*arrOut)[temp] = arr[temp]
	}
}

// searchPelanggan mencari pelanggan untuk digunakan pada Edit dan Delete, Mengembalikan Indeks jika ketemu, mengembalikan -1 jika tidak.
func searchPelanggan(arr ArrPelanggan, key string, nArr int) int {
	var indeks int
	var i int
	for i := 0; i < nArr && strings.ToLower(arr[i].nama) != strings.ToLower(key); i++ {
	}
	if strings.ToLower(arr[i].nama) != strings.ToLower(key) {
		indeks = -1
	} else if strings.ToLower(arr[i].nama) == strings.ToLower(key) {
		indeks = i
	}
	return indeks
}

func searchService(arr ArrService, key string, nArr int) int {
	var indeks int
	var i int
	for i := 0; i < nArr && strings.ToLower(arr[i].nama) != strings.ToLower(key); i++ {
	}
	if strings.ToLower(arr[i].nama) != strings.ToLower(key) {
		indeks = -1
	} else if strings.ToLower(arr[i].nama) == strings.ToLower(key) {
		indeks = i
	}
	return indeks
}

func searchMotor(arr ArrMotor, key string, nArr int) int {
	var awal, tengah, akhir int
	awal = 0
	akhir = nArr - 1
	tengah = (awal + akhir) / 2
	for awal < akhir && arr[tengah].merek != key {
		if arr[tengah].merek > key {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
	}
	if arr[tengah].merek == key {
		return tengah
	}
	return -1
}

// Misalkan Data sudah di sort
func searchSparePart(arr ArrSparepart, key string, nArr int) int {
	var awal, tengah, akhir int
	awal = 0
	akhir = nArr - 1
	tengah = (awal + akhir) / 2
	for awal < akhir && arr[tengah].nama != key {
		if arr[tengah].nama > key {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
	}
	if arr[tengah].nama == key {
		return tengah
	}
	return -1
}

func listPelanggan(arr ArrPelanggan, nArr int) {
	for i := 0; i < nArr; i++ {
		if arr[i].nama != "" {
			fmt.Println("========================================")
			fmt.Printf("Nama: %s ", arr[i].nama)
			fmt.Printf("Merek Motor: %s \n", arr[i].specMotor.merek)
			fmt.Printf("Jenis Motor: %s \n", arr[i].specMotor.jenisMotor)
			fmt.Printf("Tahun Pabrikan Motor: %d \n", arr[i].specMotor.tahunPabrikan)
		}
	}
}

func listMotor(arr ArrMotor, nArr int) {
	for i := 0; i < nArr; i++ {
		if arr[i].merek != "" && arr[i].jenisMotor != "" && arr[i].tahunPabrikan != 0 && arr[i].stokTersedia != 0 {
			fmt.Println("=============================")
			fmt.Printf("Merek: %s ", arr[i].merek)
			fmt.Printf("Jenis Motor: %s ", arr[i].jenisMotor)
			fmt.Printf("Tahun pabrikan: %d \n", arr[i].tahunPabrikan)
			fmt.Printf("Stok tersedia: %d \n", arr[i].stokTersedia)
		}
	}
}

func listService(arr ArrService, nArr int) {
	for i := 0; i < nArr; i++ {
		if arr[i].nama != "" && arr[i].hargaService != 0 {
			fmt.Println("===============================================")
			fmt.Printf("Nama Service: %s ", arr[i].nama)
			fmt.Printf("Harga Service: %d \n", arr[i].hargaService)
		}
	}
}

func listSparePart(arr ArrSparepart, nArr int) {
	for i := 0; i < nArr; i++ {
		if arr[i].nama != "" && arr[i].harga != 0 && arr[i].stokTersedia != 0 {
			fmt.Println("=================================================")
			fmt.Printf("Spare-part: %s ", arr[i].nama)
			fmt.Printf("Harga: %d \n", arr[i].harga)
			fmt.Printf("Stok: %d \n", arr[i].stokTersedia)
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
	fmt.Println("4. Service")
	fmt.Print("Pilihan: ")
}

func menuTambahEdit() {
	fmt.Println("Silahkan pilih Edit / Tambah data: ")
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Delete")
	fmt.Print("Pilihan: ")
}

func menuSortingSpareBy() {
	fmt.Println("Silahkan pilih Sorting berdasarkan: ")
	fmt.Println("1. Nama")
	fmt.Println("2. Harga")
	fmt.Println("3. Stok")
	fmt.Print("Pilihan: ")

}

func menuSortingMotorBy() {
	fmt.Println("Silahkan pilih Sorting berdasarkan: ")
	fmt.Println("1. Merek")
	fmt.Println("2. Tahun")
	fmt.Println("3. Stok")
	fmt.Print("Pilihan: ")

}

func menuSortingPelangganBy() {
	fmt.Println("Silahkan pilih Sorting berdasarkan: ")
	fmt.Println("1. Nama")
	fmt.Println("2. Jumlah Transaksi yang Dilakukan")
	fmt.Print("Pilihan: ")
}

func menuSortingServiceBy() {
	fmt.Println("Silahkan pilih Sorting berdasarkan: ")
	fmt.Println("1. Nama")
	fmt.Println("2. Harga")
	fmt.Print("Pilihan: ")
}

func menuTransaksi() {
	fmt.Println("Silakan pilih jenis transaksi: ")
	fmt.Println("1. Service motor")
	fmt.Println("2. Beli Sparepart")
	fmt.Print("Pilihan: ")
}
