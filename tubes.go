package main

import "fmt"

type pengeluaran struct {
	Kategori string
	Nominal  int
}

const maxData int = 1000

type data [maxData]pengeluaran

var saldo int
var anggaranAwal int

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Info Saldo")
	fmt.Println("2. Update Saldo")
	fmt.Println("3. History Transaksi")
	fmt.Println("4. Selisih Anggaran")
	fmt.Println("5. Cari Pengeluaran")
	fmt.Println("6. Exit")
	fmt.Println("-----------------------")
}

func main() {
	fmt.Print("Masukkan saldo awal: Rp ")
	fmt.Scan(&anggaranAwal)
	saldo = anggaranAwal

	var pilih int
	var daftar data
	var jumlahData int
	for pilih != 6 {
		menu()
		fmt.Print("Pilih (1/2/3/4/5/6/7)? ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			infoSaldo(daftar, jumlahData)
		} else if pilih == 2 {
			updateSaldo(&daftar, &jumlahData)
		} else if pilih == 3 {
			historyTransaksi(daftar, jumlahData)
		} else if pilih == 4 {
			selisihAnggaran(daftar, jumlahData)
		} else if pilih == 5 {
			cariPengeluaran(&daftar, &jumlahData)
		}
	}
}

func infoSaldo(daftar data, jumlahData int) {
	total := 0
	for i := 0; i < jumlahData; i++ {
		total += daftar[i].Nominal
	}
	fmt.Printf("Saldo saat ini: Rp. %d\n", saldo)
	fmt.Printf("Total anggaran: Rp. %d\n", anggaranAwal)
	fmt.Printf("Total pengeluaran: Rp. %d\n", total)
}

func updateSaldo(daftar *data, jumlahData *int) {
	var pilih int
	for pilih != 4 {
		menuUpdate()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahPengeluaran(daftar, jumlahData)
		} else if pilih == 2 {
			editPengeluaran(daftar, jumlahData)
		} else if pilih == 3 {
			hapusPengeluaran(daftar, jumlahData)
		}
	}
}

func menuUpdate() {
	fmt.Println("---------------------------")
	fmt.Println("        MENU UPDATE        ")
	fmt.Println("---------------------------")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Data Pengeluaran")
	fmt.Println("3. Hapus Data Pengeluaran")
	fmt.Println("4. Exit")
	fmt.Println("---------------------------")
}

func bacaKategori(label string) string {
	var kategori string
	fmt.Print(label)
	fmt.Scan(&kategori)
	return kategori
}

func bacaJumlah(label string) int {
	fmt.Print(label)
	var jumlah int
	fmt.Scan(&jumlah)
	return jumlah
}

func tambahPengeluaran(daftar *data, jumlahData *int) {
	if *jumlahData >= maxData {
		fmt.Println("Data penuh.")
		return
	}
	kategori := bacaKategori("Masukkan Kategori (Akomodasi/Transportasi/Konsumsi/Hiburan/Belanja): ")
	jumlah := bacaJumlah("Masukkan Jumlah: ")
	daftar[*jumlahData] = pengeluaran{kategori, jumlah}
	*jumlahData++
	saldo -= jumlah
	fmt.Println("Saldo berhasil diupdate")
}

func editPengeluaran(daftar *data, jumlahData *int) {
	if *jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran")
		return
	}
	menuPengeluaran(daftar, *jumlahData)
	nomor := noPengeluaran(*jumlahData)
	if nomor == -1 {
		return
	}
	nominalBaru := bacaJumlah("Masukkan jumlah baru: ")
	nominalLama := daftar[nomor].Nominal
	daftar[nomor].Nominal = nominalBaru
	saldo += nominalLama - nominalBaru
	fmt.Println("Saldo berhasil diupdate")
}

func menuPengeluaran(daftar *data, jumlahData int) {
	fmt.Println("Daftar Pengeluaran")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. Kategori: %s; Jumlah: Rp. %d\n", i+1, daftar[i].Kategori, daftar[i].Nominal)
	}
}

func noPengeluaran(jumlahData int) int {
	var nomor int
	fmt.Print("Masukkan nomor pengeluaran: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > jumlahData {
		fmt.Println("Nomor tidak valid")
		return -1
	}
	return nomor - 1
}

func hapusPengeluaran(daftar *data, jumlahData *int) {
	if *jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran")
		return
	}
	menuPengeluaran(daftar, *jumlahData)
	nomor := noPengeluaran(*jumlahData)
	if nomor == -1 {
		return
	}
	nominalDihapus := daftar[nomor].Nominal
	for i := nomor; i < *jumlahData-1; i++ {
		daftar[i] = daftar[i+1]
	}
	*jumlahData--
	saldo += nominalDihapus
	fmt.Println("Saldo berhasil diupdate")
}

func historyTransaksi(daftar data, jumlahData int) {
	if jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran.")
		return
	}
	var pilih int
	for pilih != 4 {
		menuHistory()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Println("Riwayat berdasarkan waktu input:")
			tampilkanHistory(daftar, jumlahData)
		} else if pilih == 2 {
			fmt.Println("Riwayat berdasarkan nominal terbesar:")
			salin := copyData(daftar, jumlahData)
			selectionSortTurun(&salin, jumlahData)
			tampilkanHistory(salin, jumlahData)
		} else if pilih == 3 {
			fmt.Println("Riwayat berdasarkan nominal terkecil:")
			salin := copyData(daftar, jumlahData)
			insertionSortNaik(&salin, jumlahData)
			tampilkanHistory(salin, jumlahData)
		}
	}
}

func menuHistory() {
	fmt.Println("---------------------------------")
	fmt.Println("          MENU HISTORY           ")
	fmt.Println("---------------------------------")
	fmt.Println("1. Berdasarkan Waktu Input")
	fmt.Println("2. Berdasarkan Nominal Terbesar")
	fmt.Println("3. Berdasarkan Nominal Terkecil")
	fmt.Println("4. Exit")
	fmt.Println("---------------------------------")
}

func tampilkanHistory(daftar data, jumlahData int) {
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s - Rp %d\n", i+1, daftar[i].Kategori, daftar[i].Nominal)
	}
}

func copyData(dataAsli data, jumlahData int) data {
	var dataSalinan data
	for i := 0; i < jumlahData; i++ {
		dataSalinan[i] = dataAsli[i]
	}
	return dataSalinan
}

func insertionSortNaik(daftar *data, jumlahData int) {
	for i := 1; i < jumlahData; i++ {
		temp := daftar[i]
		j := i - 1
		for j >= 0 && daftar[j].Nominal > temp.Nominal {
			daftar[j+1] = daftar[j]
			j--
		}
		daftar[j+1] = temp
	}
}

func selectionSortTurun(daftar *data, jumlahData int) {
	for i := 0; i < jumlahData-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if daftar[j].Nominal > daftar[maxIdx].Nominal {
				maxIdx = j
			}
		}
		daftar[i], daftar[maxIdx] = daftar[maxIdx], daftar[i]
	}
}

func selisihAnggaran(daftar data, jumlahData int) {
	total := 0
	for i := 0; i < jumlahData; i++ {
		total += daftar[i].Nominal
	}
	selisih := anggaranAwal - total
	fmt.Printf("Total Pengeluaran: Rp. %d\n", total)
	fmt.Printf("Selisih Anggaran: Rp. %d\n", selisih)
	if selisih < 0 {
		fmt.Println("Pengeluaran melebihi anggaran!")
		
		saranHemat(daftar, jumlahData)
	} else {
		fmt.Println("Pengeluaran masih dalam batas anggaran.")
	}
}

func saranHemat(daftar data, jumlahData int) {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data pengeluaran.")
		return
	}
	maks := daftar[0].Nominal
	kat := daftar[0].Kategori
	for i := 1; i < jumlahData; i++ {
		if daftar[i].Nominal > maks {
			maks = daftar[i].Nominal
			kat = daftar[i].Kategori
		}
	}
	fmt.Printf("- Kategori: %s, Jumlah: Rp. %d\n", kat, maks)
}

func cariPengeluaran(daftar *data, jumlahData *int){
	var pilih int
	for pilih != 3{
		menuCariData()
		fmt.Print("Pilih (1/2/3)? ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			var kategori string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kategori)
			sequentialSearch(*daftar, *jumlahData, kategori)
		} else if pilih == 2 {
			var nominal int
			fmt.Print("Masukkan nominal: Rp ")
			fmt.Scan(&nominal)
			binarySearch(*daftar, *jumlahData, nominal)
		}
	}
}

func menuCariData(){
	fmt.Println("-------------------------")
	fmt.Println("       MENU UPDATE       ")
	fmt.Println("-------------------------")
	fmt.Println("1. Berdasarkan Kategori")
	fmt.Println("2. Berdasarkan Nominal")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
}

func sequentialSearch(daftar data, jumlahData int, kategori string) {
	ditemukan := false
	for i := 0; i < jumlahData; i++ {
		if daftar[i].Kategori == kategori {
			fmt.Printf("Ditemukan: %s - Rp %d\n", daftar[i].Kategori, daftar[i].Nominal)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func insertionSort(daftar *data, jumlahData int) {
	for i := 1; i < jumlahData; i++ {
		idx := daftar[i]
		j := i - 1
		for j >= 0 && daftar[j].Nominal > idx.Nominal {
			daftar[j+1] = daftar[j]
			j--
		}
		daftar[j+1] = idx
	}
}

func binarySearch(daftar data, jumlahData int, nominal int) {
	insertionSort(&daftar, jumlahData)
	left := 0
	right := jumlahData - 1
	ditemukan := false
	for left <= right && !ditemukan {
		mid := (left + right) / 2
		if daftar[mid].Nominal == nominal {
			fmt.Printf("Ditemukan: %s - Rp %d\n", daftar[mid].Kategori, daftar[mid].Nominal)
			ditemukan = true
			for i := mid - 1; i >= 0 && daftar[i].Nominal == nominal; i-- {
				fmt.Printf("Ditemukan: %s - Rp %d\n", daftar[i].Kategori, daftar[i].Nominal)
			}
			for i := mid + 1; i < jumlahData && daftar[i].Nominal == nominal; i++ {
				fmt.Printf("Ditemukan: %s - Rp %d\n", daftar[i].Kategori, daftar[i].Nominal)
			}
		} else if nominal < daftar[mid].Nominal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !ditemukan {
		fmt.Println("Nominal tidak ditemukan.")
	}
}
