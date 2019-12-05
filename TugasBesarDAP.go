package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

// Pelanggan Tipe bentukan pelanggan yang berisi nama, specMotor, jumTransaksi.
type Pelanggan struct {
	nama         string
	specMotor    Motor
	jumTransaksi [100]Transaksi
	countTrans   int
	totalHarga   int
}

// Service Tipe bentukan service yang berisi nama, hargaService.
type Service struct {
	nama         string
	hargaService string
}

// Transaksi Tipe bentukan Transaksi yang berisi namaTrans, hargaTrans.
type Transaksi struct {
	namaTrans  string
	hargaTrans string
	waktu	   History
}

// SparePart Tipe bentukan SparePart yang berisi nama, harga, stokTersedia.
type SparePart struct {
	nama         string
	harga        string
	stokTersedia string
}

// History Tipe bentukan history yang berisi tglTransaksi, blnTransaksi, thnTransaksi.
type History struct {
	tglTransaksi string
	blnTransaksi string
	thnTransaksi string
}

// Motor Tipe bentukan motor yang berisi tahunPabrikan, merek, jenisMotr, stokTersedia.
type Motor struct {
	tahunPabrikan string
	merek         string
	jenisMotor    string
	stokTersedia  string
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
	var pendapatanTahun int
	choiceMenu = "Yes"
	for strings.ToLower(choiceMenu) == "yes" {
		cls()
		menuUtama()
		fmt.Scan(&pilihan)
		for pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
			fmt.Print("Maaf input tidak valid, Silahkan masukkan lagi: ")
			fmt.Scan(&pilihan)
		}
 		switch pilihan {
		case "1":
			cls()
			menuTambahEdit()
			fmt.Scan(&pilihan)
			switch pilihan {
			case "1":
				cls()
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					tambahSparePart(&dataSparepart, &countArrSparepart)
				case "2":
					cls()
					nambahMotor(&dataMotor, &countArrMotor)
				case "3":
					cls()
					tambahPelanggan(&dataPelanggan, &countArrPelanggan)
				case "4":
					cls()
					tambahService(&dataService, &countArrService)
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
			case "2":
				cls()
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					listSparePart(dataSparepart, countArrSparepart)
					fmt.Print("Sparepart yang ingin diedit (berdasarkan nama spare-part): ")
					fmt.Scanln()
					spareDicari, _ = inputreader.ReadString('\n')
					spareDicari = strings.TrimSpace(spareDicari)
					idxSpare = searchSparePart(dataSparepart, spareDicari, countArrSparepart)
					if idxSpare != -1 {
						editSparePart(dataSparepart, idxSpare, &dataSparepart)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", spareDicari)
					}
				case "2":
					cls()
					listMotor(dataMotor, countArrMotor)
					fmt.Print("Motor yang ingin diedit (berdasarkan merek motor): ")
					fmt.Scanln()
					motorDicari, _ = inputreader.ReadString('\n')
					motorDicari = strings.TrimSpace(motorDicari)
					idxMotor = searchMotor(dataMotor, motorDicari, countArrMotor)
					if idxMotor != -1 {
						editMotor(dataMotor, idxMotor, &dataMotor)
					} else {
						fmt.Printf("Maaf tidak terdapat %s.\n", motorDicari)
					}
				case "3":
					cls()
					listPelanggan(dataPelanggan, countArrPelanggan)
					fmt.Print("Pelanggan yang ingin diedit (berdasarkan nama pelanggan): ")
					fmt.Scanln()
					pelangganDicari, _ = inputreader.ReadString('\n')
					pelangganDicari = strings.TrimSpace(pelangganDicari)
					idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
					if idxPelanggan != -1 {
						editPelanggan(dataPelanggan, idxPelanggan, &dataPelanggan)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", pelangganDicari)
					}
				case "4":
					cls()
					listService(dataService, countArrService)
					fmt.Print("Jenis service yang ingin diedit (berdasarkan nama service): ")
					fmt.Scanln()
					serviceDicari, _ = inputreader.ReadString('\n')
					serviceDicari = strings.TrimSpace(serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						editService(dataService, idxService, &dataService)
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", serviceDicari)
					}
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
			case "3":
				cls()
				menuJenisData()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					listSparePart(dataSparepart, countArrSparepart)
					fmt.Print("Sparepart yang ingin dihapus: ")
					fmt.Scanln()
					spareDicari, _ = inputreader.ReadString('\n')
					spareDicari = strings.TrimSpace(spareDicari)
					idxSpare = searchSparePart(dataSparepart, spareDicari, countArrSparepart)
					if idxSpare != -1 {
						deleteSparepart(dataSparepart, idxSpare, &dataSparepart)
						fmt.Println("Data telah dihapus.")
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", motorDicari)
					}
				case "2":
					cls()
					listMotor(dataMotor, countArrPelanggan)
					fmt.Print("Motor yang ingin dihapus: ")
					fmt.Scanln()
					motorDicari, _ = inputreader.ReadString('\n')
					motorDicari = strings.TrimSpace(motorDicari)
					idxMotor = searchMotor(dataMotor, motorDicari, countArrMotor)
					if idxMotor != -1 {
						deleteMotor(dataMotor, idxMotor, &dataMotor)
						fmt.Println("Data telah dihapus.")
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", motorDicari)
					}
				case "3":
					cls()
					listPelanggan(dataPelanggan, countArrPelanggan)
					fmt.Print("Pelanggan yang ingin dihapus: ")
					fmt.Scanln()
					pelangganDicari, _ = inputreader.ReadString('\n')
					pelangganDicari = strings.TrimSpace(pelangganDicari)
					idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
					if idxPelanggan != -1 {
						deletePelanggan(dataPelanggan, idxPelanggan, &dataPelanggan)
						fmt.Println("Data telah dihapus.")
					} else {
						fmt.Printf("Maaf idak terdapat %s.\n", pelangganDicari)
					}
				case "4":
					cls()
					listService(dataService, countArrService)
					fmt.Print("Service yang ingin dihapus: ")
					fmt.Scanln()
					serviceDicari, _ = inputreader.ReadString('\n')
					serviceDicari = strings.TrimSpace(serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						deleteService(dataService, idxService, &dataService)
						fmt.Println("Data telah dihapus.")
					} else {
						fmt.Printf("Maaf tidak terdapat %s. \n", serviceDicari)
					}
				}
			default:
				fmt.Println("Maaf pilihan tidak valid.")
			}
		case "2":
			cls()
			menuTransaksi()
			fmt.Scan(&pilihan)
			for pilihan != "1" && pilihan != "2" {
				fmt.Println("Maaf input tidak valid.")
				fmt.Print("Silahkan coba lagi: ")
				fmt.Scan(&pilihan)
			}
			cls()
			viewPelanggan(dataPelanggan, countArrPelanggan)
			fmt.Println("Silahkan masukan nama pelanggan: ")
			fmt.Scanln()
			pelangganDicari, _ = inputreader.ReadString('\n')
			pelangganDicari = strings.TrimSpace(pelangganDicari)
			idxPelanggan = searchPelanggan(dataPelanggan, pelangganDicari, countArrPelanggan)
			if idxPelanggan != 1 {
				switch pilihan {
				case "1":
					fmt.Println("Silahkan masukkan jenis service yang ingin dilakukan: ")
					sortServiceHarga(dataService, countArrService, &dataService)
					listService(dataService, countArrService)
					fmt.Print("Pilihan: ")
					serviceDicari, _ = inputreader.ReadString('\n')
					serviceDicari = strings.TrimSpace(serviceDicari)
					idxService = searchService(dataService, serviceDicari, countArrService)
					if idxService != -1 {
						tambahTransService(&dataPelanggan, dataService, idxService, idxPelanggan)
					}else {
						fmt.Printf("Maaf tidak terdapat %s. \n",serviceDicari)
					}
				case "2":
					fmt.Println("Silahkan masukkan jenis service yang ingin dilakukan: ")
					sortSparePartHarga(dataSparepart, countArrSparepart, &dataSparepart)
					listSparePart(dataSparepart, countArrSparepart)
					fmt.Print("Pilihan: ")
					spareDicari, _ =inputreader.ReadString('\n')
					spareDicari = strings.TrimSpace(spareDicari)
					idxSpare = searchSparePart(dataSparepart, spareDicari, countArrSparepart)
					if idxSpare != -1 {
						tambahTransSpare(&dataPelanggan, &dataSparepart, idxSpare, idxPelanggan)
					}else {
						fmt.Printf("Maaf tidak terdapat %s. \n",spareDicari)
					}
				}
			} else {
				fmt.Printf("Maaf tidak terdapat %s. \n", pelangganDicari)
			}

		case "3":
			cls()
			menuJenisData()
			fmt.Scan(&pilihan)
			switch pilihan {
			case "1":
				cls()
				menuSortingSpareBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					sortSparePartNama(dataSparepart, countArrSparepart, &dataSparepart)
				case "2":
					cls()
					sortSparePartHarga(dataSparepart, countArrSparepart, &dataSparepart)
				case "3":
					cls()
					sortSparePartStok(dataSparepart, countArrSparepart, &dataSparepart)
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
				listSparePart(dataSparepart, countArrSparepart)
			case "2":
				cls()
				menuSortingMotorBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					sortMotorMerek(dataMotor, countArrMotor, &dataMotor)
				case "2":
					cls()
					sortMotorTahun(dataMotor, countArrMotor, &dataMotor)
				case "3":
					cls()
					sortMotorStok(dataMotor, countArrMotor, &dataMotor)
				default:
					fmt.Println("Maaf pilihan tidak valid.")
				}
				listMotor(dataMotor, countArrMotor)
			case "3":
				cls()
				menuSortingPelangganBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					sortPelangganNama(dataPelanggan, countArrPelanggan, &dataPelanggan)
				case "2":
					cls()
					sortPelangganCount(dataPelanggan, countArrPelanggan, &dataPelanggan)
				}
				listPelanggan(dataPelanggan, countArrPelanggan)
			case "4":
				cls()
				menuSortingServiceBy()
				fmt.Scan(&pilihan)
				switch pilihan {
				case "1":
					cls()
					sortServiceNama(dataService, countArrService, &dataService)
				case "2":
					cls()
					sortServiceHarga(dataService, countArrService, &dataService)
				}
				listService(dataService, countArrService)
			default:
				fmt.Println("Maaf pilihan tidak valid.")
			}
		case "4":
			cls()
			totalPendapatan(dataPelanggan, &pendapatanTahun, countArrPelanggan)
			fmt.Printf("Total Pendapatan: %d\n",pendapatanTahun)
		}
		// cls()
		fmt.Println("Kembali ke menu utama? (Yes / No): ")
		fmt.Scan(&choiceMenu)
		for strings.ToLower(choiceMenu) != "no" && strings.ToLower(choiceMenu) != "yes" {
			fmt.Println("Maaf input anda tidak valid, silahkan masukkan (Yes) atau (No)")
			fmt.Scan(&choiceMenu)
		}
	}
}

//untuk nge cek inputannya integer semua atau ngga
func checkInt(x int) (bool bool){
	if x > 0{
		bool = true
	}else{
		bool = false
	}
	return bool
}

func nambahMotor(arr *ArrMotor, nArr *int) {
	var (
		kembali string
		bool bool
	)
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].tahunPabrikan == "" && arr[i].merek == "" && arr[i].jenisMotor == "" && arr[i].stokTersedia == "" {
			fmt.Print("Merek: ")
			fmt.Scanln()
			(*arr)[i].merek, _ = inputreader.ReadString('\n')
			(*arr)[i].merek = strings.TrimSpace((*arr)[i].merek)
			fmt.Print("Jenis motor: ")
			(*arr)[i].jenisMotor, _ = inputreader.ReadString('\n')
			(*arr)[i].jenisMotor = strings.TrimSpace((*arr)[i].jenisMotor)
			fmt.Print("Tahun pabrikan: ")
			fmt.Scanln(&arr[i].tahunPabrikan)
			thnpab, _ := strconv.Atoi(arr[i].tahunPabrikan)
			bool = checkInt(thnpab)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Tahun pabrikan: ")
				fmt.Scanln(&arr[i].tahunPabrikan)
				thnpab, _ := strconv.Atoi(arr[i].tahunPabrikan)
				bool = checkInt(thnpab)
			}
			fmt.Print("Stok tersedia: ")
			fmt.Scanln(&arr[i].stokTersedia)
			stock, _ := strconv.Atoi(arr[i].stokTersedia)
			bool = checkInt(stock)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Stok tersedia: ")
				fmt.Scanln(&arr[i].stokTersedia)
				stock, _ := strconv.Atoi(arr[i].stokTersedia)
				bool = checkInt(stock)
			}
			fmt.Println("Apakah ingin menambahkan lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahService(arr *ArrService, nArr *int) {
	var (
		kembali string
		bool bool
	)
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Print("Nama Service: ")
			fmt.Scanln()
			(*arr)[i].nama, _ = inputreader.ReadString('\n')
			(*arr)[i].nama = strings.TrimSpace((*arr)[i].nama)
			fmt.Print("Tarif Service: ")
			fmt.Scanln(&arr[i].hargaService)
			harga, _ := strconv.Atoi(arr[i].hargaService)
			bool = checkInt(harga)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Tarif Service: ")
				fmt.Scanln(&arr[i].hargaService)
				harga, _ := strconv.Atoi(arr[i].hargaService)
				bool = checkInt(harga)
			}
			fmt.Println("Apakah ingin menambahkan lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahSparePart(arr *ArrSparepart, nArr *int) {
	var (
		kembali string
		bool bool
	)
	kembali = "yes"
	for i := 0; i < IsiArray && strings.ToLower(kembali) == "yes"; i++ {
		if arr[i].nama == "" && arr[i].harga == "" && arr[i].stokTersedia == "" {
			fmt.Print("Nama spare-part: ")
			fmt.Scanln()
			(*arr)[i].nama, _ = inputreader.ReadString('\n')
			(*arr)[i].nama = strings.TrimSpace((*arr)[i].nama)
			fmt.Print("Jumlah harga: ")
			fmt.Scan(&arr[i].harga)
			harga, _ := strconv.Atoi(arr[i].harga)
			bool = checkInt(harga)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Jumlah harga: ")
				fmt.Scan(&arr[i].harga)
				harga, _ := strconv.Atoi(arr[i].harga)
				bool = checkInt(harga)
			}
			fmt.Print("Stok tersedia: ")
			fmt.Scan(&arr[i].stokTersedia)
			stock, _ := strconv.Atoi(arr[i].stokTersedia)
			bool = checkInt(stock)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Stok tersedia: ")
				fmt.Scan(&arr[i].stokTersedia)
				stock, _ := strconv.Atoi(arr[i].stokTersedia)
				bool = checkInt(stock)
			}
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahPelanggan(arr *ArrPelanggan, nArr *int) {
	var (
		kembali string
		bool bool
	)
	kembali = "yes"
	for i := 0; i < IsiArray && kembali == "yes"; i++ {
		if arr[i].nama == "" {
			fmt.Print("Nama pelanggan: ")
			fmt.Scanln()
			(*arr)[i].nama, _ = inputreader.ReadString('\n')
			(*arr)[i].nama = strings.TrimSpace((*arr)[i].nama)
			fmt.Println("Silahkan masukan Spesifikasi Motor")
			fmt.Print("Merek motor: ")
			(*arr)[i].specMotor.merek, _ = inputreader.ReadString('\n')
			(*arr)[i].specMotor.merek = strings.TrimSpace((*arr)[i].specMotor.merek)
			fmt.Print("Jenis motor: ")
			(*arr)[i].specMotor.jenisMotor, _ = inputreader.ReadString('\n')
			(*arr)[i].specMotor.jenisMotor = strings.TrimSpace((*arr)[i].specMotor.jenisMotor)
			fmt.Print("Tahun pabrikan Motor: ")
			fmt.Scan(&arr[i].specMotor.tahunPabrikan)
			thnpab2, _ := strconv.Atoi(arr[i].specMotor.tahunPabrikan)
			bool = checkInt(thnpab2)
			for bool == false {
				fmt.Println("Inputan harus integer, silahkan ulangi.")
				fmt.Print("Tahun pabrikan Motor: ")
				fmt.Scan(&arr[i].specMotor.tahunPabrikan)
				thnpab2, _ := strconv.Atoi(arr[i].specMotor.tahunPabrikan)
				bool = checkInt(thnpab2)
			}
			fmt.Println("Apakah ingin memasukan data lagi? (Yes/No): ")
			fmt.Scan(&kembali)
			*nArr = *nArr + 1
		}
	}
}

func tambahTransSpare(arr *ArrPelanggan, arrData *ArrSparepart, indexSpare, indexPelanggan int) {
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].namaTrans = arrData[indexSpare].nama
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].hargaTrans = arrData[indexSpare].harga
	stock, _ := strconv.Atoi((*arrData)[indexSpare].stokTersedia)
	if stock != 0 {
		sisa := stock - 1
		(*arrData)[indexSpare].stokTersedia = strconv.Itoa(sisa)
		(*arr)[indexPelanggan].countTrans = (*arr)[indexPelanggan].countTrans + 1
		fmt.Println("Masukkan tanggal transaksi (Contoh: 20 Maret 2019): ")
		fmt.Scan(&arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.tglTransaksi, &arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.blnTransaksi, &arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.thnTransaksi)
	}else {
		fmt.Println("Maaf stock habis.")
	}
}

func tambahTransService(arr *ArrPelanggan, arrData ArrService, indexService, indexPelanggan int) {
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].namaTrans = arrData[indexService].nama
	(*arr)[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].hargaTrans = arrData[indexService].hargaService
	(*arr)[indexPelanggan].countTrans = (*arr)[indexPelanggan].countTrans + 1
	fmt.Println("Masukkan tanggal transaksi (Contoh: 20 Maret 2019): ")
	fmt.Scan(&arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.tglTransaksi, &arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.blnTransaksi, &arr[indexPelanggan].jumTransaksi[(*arr)[indexPelanggan].countTrans].waktu.thnTransaksi)
}

func editService(arr ArrService, n int, arrOut *ArrService) {
	var bool bool
	fmt.Print("Nama service: ")
	(arr)[n].nama, _ = inputreader.ReadString('\n')
	(arr)[n].nama = strings.TrimSpace((arr)[n].nama)
	fmt.Print("Tarif Service: ")
	fmt.Scanln(&arr[n].hargaService)
	harga, _ := strconv.Atoi(arr[n].hargaService)
	bool = checkInt(harga)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Tarif Service: ")
		fmt.Scanln(&arr[n].hargaService)
		harga, _ := strconv.Atoi(arr[n].hargaService)
		bool = checkInt(harga)
	}
	(*arrOut)[n] = arr[n]
}

func editMotor(arr ArrMotor, n int, arrOut *ArrMotor) {
	var bool bool
	fmt.Print("Merek: ")
	(arr)[n].merek, _ = inputreader.ReadString('\n')
	(arr)[n].merek = strings.TrimSpace((arr)[n].merek)
	fmt.Print("Jenis motor: ")
	(arr)[n].jenisMotor, _ = inputreader.ReadString('\n')
	(arr)[n].jenisMotor = strings.TrimSpace((arr)[n].jenisMotor)
	fmt.Print("Tahun pabrikan: ")
	fmt.Scanln(&arr[n].tahunPabrikan)
	thnpab, _ := strconv.Atoi(arr[n].tahunPabrikan)
	bool = checkInt(thnpab)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Tahun pabrikan: ")
		fmt.Scanln(&arr[n].tahunPabrikan)
		thnpab, _ := strconv.Atoi(arr[n].tahunPabrikan)
		bool = checkInt(thnpab)
	}
	fmt.Print("Stok tersedia: ")
	fmt.Scanln(&arr[n].stokTersedia)
	stock, _ := strconv.Atoi(arr[n].stokTersedia)
	bool = checkInt(stock)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Stok tersedia: ")
		fmt.Scanln(&arr[n].stokTersedia)
		stock, _ := strconv.Atoi(arr[n].stokTersedia)
		bool = checkInt(stock)
	}
	(*arrOut)[n] = arr[n]
}

func editSparePart(arr ArrSparepart, n int, arrOut *ArrSparepart) {
	var bool bool
	fmt.Print("Nama spare-part: ")
	(arr)[n].nama, _ = inputreader.ReadString('\n')
	(arr)[n].nama = strings.TrimSpace((arr)[n].nama)
	fmt.Print("Jumlah harga: ")
	fmt.Scan(&arr[n].harga)
	harga, _ := strconv.Atoi(arr[n].harga)
	bool = checkInt(harga)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Jumlah harga: ")
		fmt.Scan(&arr[n].harga)
		harga, _ := strconv.Atoi(arr[n].harga)
		bool = checkInt(harga)
	}
	fmt.Print("Stok tersedia: ")
	fmt.Scan(&arr[n].stokTersedia)
	stock, _ := strconv.Atoi(arr[n].stokTersedia)
	bool = checkInt(stock)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Stok tersedia: ")
		fmt.Scan(&arr[n].stokTersedia)
		stock, _ := strconv.Atoi(arr[n].stokTersedia)
		bool = checkInt(stock)
	}
	(*arrOut)[n] = arr[n]
}

func editPelanggan(arr ArrPelanggan, n int, arrOut *ArrPelanggan) {
	var bool bool
	fmt.Print("Nama pelanggan: ")
	(arr)[n].nama, _ = inputreader.ReadString('\n')
	(arr)[n].nama = strings.TrimSpace((arr)[n].nama)
	fmt.Println("Silahkan masukan Spesifikasi Motor")
	fmt.Print("Merek motor: ")
	(arr)[n].specMotor.merek, _ = inputreader.ReadString('\n')
	(arr)[n].specMotor.merek = strings.TrimSpace((arr)[n].specMotor.merek)
	fmt.Print("Jenis motor: ")
	(arr)[n].specMotor.jenisMotor, _ = inputreader.ReadString('\n')
	(arr)[n].specMotor.jenisMotor = strings.TrimSpace((arr)[n].specMotor.jenisMotor)
	fmt.Print("Tahun pabrikan Motor: ")
	fmt.Scan(&arr[n].specMotor.tahunPabrikan)
	thnpab2, _ := strconv.Atoi(arr[n].specMotor.tahunPabrikan)
	bool = checkInt(thnpab2)
	for bool == false {
		fmt.Println("Inputan harus integer, silahkan ulangi.")
		fmt.Print("Tahun pabrikan Motor: ")
		fmt.Scan(&arr[n].specMotor.tahunPabrikan)
		thnpab2, _ := strconv.Atoi(arr[n].specMotor.tahunPabrikan)
		bool = checkInt(thnpab2)
	}
	(*arrOut)[n] = arr[n]
}

func deleteService(arr ArrService, indeks int, arrOut *ArrService) {
	arr[indeks].nama = ""
	arr[indeks].hargaService = ""
	(*arrOut)[indeks] = arr[indeks]
}

func deletePelanggan(arr ArrPelanggan, indeks int, arrOut *ArrPelanggan) {
	arr[indeks].nama = ""
	arr[indeks].specMotor.jenisMotor = ""
	arr[indeks].specMotor.merek = ""
	arr[indeks].specMotor.stokTersedia = ""
	arr[indeks].specMotor.tahunPabrikan = ""
	(*arrOut)[indeks] = arr[indeks]
}

func deleteMotor(arr ArrMotor, indeks int, arrOut *ArrMotor) {
	arr[indeks].jenisMotor = ""
	arr[indeks].merek = ""
	arr[indeks].stokTersedia = ""
	arr[indeks].tahunPabrikan = ""
}

func deleteSparepart(arr ArrSparepart, indeks int, arrOut *ArrSparepart) {
	arr[indeks].harga = ""
	arr[indeks].nama = ""
	arr[indeks].stokTersedia = ""
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
	var(
		j int
		temp string
	) 
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
	for awal < akhir && strings.ToLower(arr[tengah].jenisMotor) != strings.ToLower(key) {
		if strings.ToLower(arr[tengah].jenisMotor) > strings.ToLower(key) {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
		fmt.Println(tengah)
	}
	if strings.ToLower(arr[tengah].jenisMotor) == strings.ToLower(key) {
		return tengah
	}else{
		return -1
	}
}

// Misalkan Data sudah di sort
func searchSparePart(arr ArrSparepart, key string, nArr int) int {
	var awal, tengah, akhir int
	awal = 0
	akhir = nArr - 1
	tengah = (awal + akhir) / 2
	for awal < akhir && strings.ToLower(arr[tengah].nama) != strings.ToLower(key) {
		if arr[tengah].nama > key {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
	}
	if strings.ToLower(arr[tengah].nama) == strings.ToLower(key) {
		return tengah
	}else{
		return -1
	}
}

func listPelanggan(arr ArrPelanggan, nArr int) {
	for i := 0; i <= nArr; i++ {
		if arr[i].nama != "" {
			fmt.Println("========================================")
			fmt.Println("Nama: ", arr[i].nama)
			fmt.Println("Merek Motor: ", arr[i].specMotor.merek)
			fmt.Println("Jenis Motor: ", arr[i].specMotor.jenisMotor)
			fmt.Printf("Tahun Pabrikan Motor: %s \n", arr[i].specMotor.tahunPabrikan)
			for j := 0; j<arr[i].countTrans; j++{
				fmt.Printf("Transaksi ke-%d : %s \n", j+1, arr[i].jumTransaksi[j].namaTrans)
				fmt.Printf("Tanggal Transaksi: %s %s %s \n",arr[i].jumTransaksi[j].waktu.tglTransaksi, arr[i].jumTransaksi[j].waktu.blnTransaksi, arr[i].jumTransaksi[j].waktu.thnTransaksi )
				harga,_ := strconv.Atoi(arr[i].jumTransaksi[j].hargaTrans)
				arr[i].totalHarga = arr[i].totalHarga + harga
			}
			fmt.Println("Total harga: ", arr[i].totalHarga)
			fmt.Println("========================================")
		}
	}
}

func listMotor(arr ArrMotor, nArr int) {
	for i := 0; i <= nArr; i++ {
		if arr[i].merek != "" && arr[i].jenisMotor != "" && arr[i].tahunPabrikan != "" && arr[i].stokTersedia != "" {
			fmt.Println("=============================")
			fmt.Println("Merek: ", arr[i].merek)
			fmt.Println("Jenis Motor: ", arr[i].jenisMotor)
			fmt.Println("Tahun Pabrikan: ", arr[i].tahunPabrikan)
			fmt.Println("Stok tersedia: ", arr[i].stokTersedia)
			fmt.Println("=============================")
		}
	}
}

func listService(arr ArrService, nArr int) {
	for i := 0; i <= nArr; i++ {
		if arr[i].nama != "" && arr[i].hargaService != "" {
			fmt.Println("===============================================")
			fmt.Println("Nama Service: ", arr[i].nama)
			fmt.Println("Harga Service: ", arr[i].hargaService)
			fmt.Println("===============================================")
		}
	}
}

func listSparePart(arr ArrSparepart, nArr int) {
	for i := 0; i <= nArr; i++ {
		if arr[i].nama != "" && arr[i].harga != "" && arr[i].stokTersedia != "" {
			fmt.Println("=================================================")
			fmt.Println("Spare-part: ", arr[i].nama)
			fmt.Println("Harga: ", arr[i].harga)
			fmt.Println("Stok: ", arr[i].stokTersedia)
			fmt.Println("=================================================")
		}
	}
}

func totalPendapatan(arr ArrPelanggan, totPendapatan *int, nPelanggan int) {
	*totPendapatan = 0
	for i := 0; i < nPelanggan; i++ {
		for j := 0; j < arr[i].countTrans; j++ {
			harga,_ := strconv.Atoi(arr[i].jumTransaksi[j].hargaTrans)
			*totPendapatan = *totPendapatan + harga
		}
	}
}

func viewPelanggan(arr ArrPelanggan, nPelanggan int) {
	for i:=0; i < nPelanggan; i++ {
		fmt.Printf("Pelanggan ke-%d: %s\n",i+1,arr[i].nama)
	}
}

func menuUtama() {
	fmt.Println("Selamat datang di bengkel onlen")
	fmt.Println("Silahkan pilih menu: ")
	fmt.Println("1. Edit Data")
	fmt.Println("2. Transaksi")
	fmt.Println("3. Lihat Data")
	fmt.Println("4. Lihat Pendapatan")
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

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
