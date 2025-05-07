package main
import "fmt"

type pengeluaran struct{
	Kategori string
	nominal int
}
var saldo int
var anggaranAwal int

func menu(){
	fmt.Println("\n-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Info Saldo")
	fmt.Println("2. Update Saldo")
	fmt.Println("3. History Transaksi")
	fmt.Println("4. Selisih Anggaran")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------")
}

func main(){
	fmt.Print("Masukkan saldo awal: Rp ")
	fmt.Scan(&anggaranAwal)
	saldo = anggaranAwal
	var pilih int
	var pengeluaran []pengeluaran
	for pilih != 5 {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)
			if pilih == 1{
				infoSaldo(pengeluaran)
			} else if pilih == 2{
				updateSaldo(&pengeluaran)
			// } else if pilih == 3{
			// 	historyTransaksi()
			// } else if pilih == 4{
			// 	selisihAnggaran()
			// } else if pilih == 5{
			// 	return
			}
	}
}

func infoSaldo(daftarPengeluaran []pengeluaran) {
	totalPengeluaran := 0
	for _, pengeluaran := range daftarPengeluaran{
		totalPengeluaran += pengeluaran.nominal
	}
	fmt.Printf("Saldo saat ini: Rp. %d\n", saldo)
	fmt.Printf("Total anggaran: Rp. %d\n", anggaranAwal)
	fmt.Printf("Total pengeluaran: Rp. %d\n", totalPengeluaran)
}

func bacaKetegori(kategori string) string{
	fmt.Print(kategori)
	fmt.Scan(&kategori)
	return kategori
}

func bacaJumlah() int{
	fmt.Print("Masukkan Jumlah:")
	var jumlah int
	fmt.Scan(&jumlah)
	return jumlah
}

func tambahPengeluaran(daftarPengeluaran *[]pengeluaran){
	kategori := bacaKetegori("Masukkan Kategori (Akomodasi/Trasnportasi/Makanan dan Minuman/Hiburan/Belanja):")
	jumlah := bacaJumlah()
	*daftarPengeluaran = append(*daftarPengeluaran, pengeluaran{Kategori: kategori, nominal: jumlah})
	saldo -= jumlah
	fmt.Println("Saldo berhasil diupdate")
}

func editPengeluaran(daftarPengeluaran *[]pengeluaran) {
	if len(*daftarPengeluaran) == 0{
		fmt.Print("Belum ada pengeluaran\n")
		return
	}
	menuPengeluaran(*daftarPengeluaran)
	nomor := noPengeluaran(len(*daftarPengeluaran))
	if nomor == -1{
		return
	}
	nominalBaru := bacaJumlah()
	nominalLama := (*daftarPengeluaran)[nomor-1].nominal
	(*daftarPengeluaran)[nomor-1].nominal = nominalBaru
	saldo += nominalLama
	saldo -= nominalBaru
	fmt.Print("Saldo berhasil diupdate")
}

func menuPengeluaran(daftarPengeluaran []pengeluaran){
	fmt.Println("Daftar Pengeluaran")
	for i, pengeluran := range daftarPengeluaran {
		fmt.Printf("%d. Kategori: %s; Jumlah: Rp. %d\n", i+1, pengeluran.Kategori, pengeluran.nominal)
	}
}

func noPengeluaran(panjangDaftar int) int{
	fmt.Print("Masukkan nomor pengeluaran:")
	var nomorPengeluaran int
	fmt.Scan(&nomorPengeluaran)
	if nomorPengeluaran < 1 || nomorPengeluaran > panjangDaftar{
		fmt.Print("Nomor tidak valid\n")
		return -1
	}
	return nomorPengeluaran
}

func hapusPengeluaran(daftarPengeluaran *[]pengeluaran){
	if len(*daftarPengeluaran) == 0{
		fmt.Print("Belum ada pengeluaran")
		return
	}
	menuPengeluaran(*daftarPengeluaran)
	nomor := noPengeluaran(len(*daftarPengeluaran))
	if nomor == -1{
		return
	}
	nominalDihapus := (*daftarPengeluaran)[nomor-1].nominal
	*daftarPengeluaran = append((*daftarPengeluaran) [:nomor-1], (*daftarPengeluaran) [nomor:]...)
	saldo += nominalDihapus
	fmt.Print("Saldo berhasil diupdate")
}

func updateSaldo(daftarPengeluaran *[]pengeluaran){
	var pilih int
	for pilih != 4 {
		menuUpdate()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		if pilih == 1{
			tambahPengeluaran(daftarPengeluaran)
		} else if pilih == 2{
			editPengeluaran(daftarPengeluaran)
		} else if pilih == 3{
			hapusPengeluaran(daftarPengeluaran)
		} else if pilih == 4{
			return
		}
	}
}
			
func menuUpdate(){
	fmt.Println("\n-----------------------")
	fmt.Println("      MENU UPDATE      ")
	fmt.Println("-----------------------")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Data Pengeluaran")
	fmt.Println("3. Hapus Data pengeluaran")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}
