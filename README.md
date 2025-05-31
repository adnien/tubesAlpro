# -tubesAlpro-
Aplikasi Pengelolaan Budget Traveling 

Aplikasi Pengelolaan Budget Traveling adalah aplikasi berbasis terminal yang dikembangkan menggunakan bahasa pemrograman Go (Golang), dirancang untuk membantu pengguna dalam mencatat dan mengelola keuangan selama perjalanan mereka. Dengan aplikasi ini, pengguna dapat memasukkan berbagai jenis pengeluaran berdasarkan kategori. Kategori yang disediakan meliputi transportasi, akomodasi, konsumsi, hiburan, serta belanja. Aplikasi ini juga memungkinkan pengguna untuk memantau sisa saldo, mengetahui total pengeluaran yang telah dilakukan, mencari pengeluaran berdasarkan kategori atau nominal, serta memperoleh saran penghematan ketika pengeluaran melebihi anggaran awal yang telah ditentukan.

Aplikasi ini cocok digunakan oleh siapa saja yang ingin merancang anggaran perjalanan secara sederhana ataupun yang ingin mengevaluasi pengeluaran setelah melakukan perjalanan.

Fitur-Fitur Program:

1. Info Saldo

   Informasi ini membantu pengguna mengetahui sisa dana dan progres terhadap rencana anggaran.

   Fitur ini menampilkan:

    a. Saldo saat ini.

    b. Anggaran awal (budget awal pengguna).

    c. Total seluruh pengeluaran yang telah dicatat.

  Fungsi terkait: infoSaldo()

  Proses input: Akses dari menu utama, pilih opsi satu (Info Saldo)

2. Update Saldo

  Fungsi Terkait: updateSaldo(),  menuUpdate()

  a. Tambah Pengeluaran:

  Pengguna dapat menambahkan data pengeluaran baru dengan memilih kategori dan memasukkan jumlah nominal. Setiap penambahan akan otomatis mengurangi saldo saat ini.

  Fungsi terkait: bacaKategori(), bacaJumlah(), tambahPengeluaran()
  
  Proses input: 

  1. akses dari menu utama, pilih opsi dua (Update saldo)

  2. Pilih opsi satu (Tambah Pengeluaran)

  3. Input Kategori (Contoh: Akomodasi)

  4. Input Jumlah (Contoh: 500000)

  b. Edit Pengeluaran:

  Memungkinkan pengguna memperbarui jumlah pengeluaran tertentu jika terjadi kesalahan input atau perubahan biaya. Saldo akan disesuaikan berdasarkan perubahan nominal.

  Fungsi terkait: bacaJumlah(), editPengeluaran(), noPengeluaran(), menuPengeluaran()

  Proses input: 
  1. akses dari menu utama, pilih opsi dua (Update saldo)
  
  2. Pilih opsi dua (Edit Pengeluaran)
  
  3. Pilih nomor data yang akan diedit
  
  4. Input nominal baru

  c. Hapus Pengeluaran:

  Pengguna bisa menghapus data pengeluaran tertentu. Nominal yang dihapus akan dikembalikan ke saldo.

  Fungsi terkait: hapusPengeluaran(), menuPengeluaran(), noPengeluaran()

  Proses input: 

  1. akses dari menu utama, pilih opsi dua (Update saldo)
 
  2. Pilih opsi tiga (Hapus Pengeluaran)
  
  3. Pilih nomor data yang akan dihapus

3. History Transaksi

  Fungsi terkait: menuHistory(), historyTransaksi()
  Pengguna dapat melihat daftar seluruh pengeluaran berdasarkan tiga kriteria berikut:
 
  a. Waktu input (default): ditampilkan sesuai urutan data dimasukkan.
    
  Fungsi terkait: tampilkanHistory()

  Proses input: 

  1. Akses dari menu utama, pilih opsi tiga (History Transaksi)

  2. Pilih metode tampilan satu (Berdasarkan Waktu Input)

  b. Nominal terbesar: Menggunakan Selection Sort, pengeluaran diurutkan dari jumlah tertinggi ke terendah.

  Fungsi terkait: copyData(), SelectionSortTurun(), tampilkanHistory().

  Proses input: 

  1. Akses dari menu utama, pilih opsi tiga (History Transaksi)

  2. Pilih metode tampilan dua (Berdasarkan Nominal terbesar)
  
  c. Nominal Terkecil: Menggunakan Insertion Sort, pengeluaran diurutkan dari jumlah terendah ke tertinggi.

  Fungsi terkait: copyData(), SelectionSortNaik(), tampilkanHistory().

  Proses input: 

  1. Akses dari menu utama, pilih opsi tiga (History Transaksi)

  2. Pilih metode tampilan satu (Berdasarkan Nominal Terkecil).

4. Selisih Anggaran

  a. Menghitung selisih anggaran awal dengan total pengeluaran.
  
  b. Menampilkan pesan apakah pengguna masih berada dalam batas anggaran atau sudah melebihi.
  
  c. Jika melebihi, sistem akan memberikan saran kategori mana yang paling besar pengeluarannya dan sebaiknya dikurangi atau dihemat.

  Fungsi terkait: selisihAnggaran(), saranHemat()

  Proses input: Akses dari menu utama, pilih opsi empat (selisih anggaran)

5. Pencarian Pengeluaran

  Fungsi pengeluaran: cariPengeluaran(), menuCariData(), 

  a. Berdasarkan Kategori: Menggunakan Sequential Search, menampilkan semua pengeluaran yang memiliki kategori tertentu (misal: “Transportasi”). Cocok untuk mencari berdasarkan nama tanpa perlu diurutkan.

  Fungsi terkait: sequentialSearch()
  
  Proses input: 
  
  1. Akses dari menu utama, pilih opsi lima (Cari Pengeluaran)
  
  2. Pilih metode pencarian satu (Berdasarkan kategori)
  
  3. Masukkan nama kategori (contoh: akomodasi)

  b. Berdasarkan Nominal: Menggunakan Binary Search, menampilkan pengeluaran berdasarkan jumlah nominal tertentu. Menggunakan Insertion Sort terlebih dahulu untuk mengurutkan data, lalu mencari dengan binary search.

  Fungsi terkait: insertionSort(), binarySearch()

  Proses input: 

  1. Akses dari menu utama, pilih opsi lima (Cari Pengeluaran)

  2. Pilih metode pencarian satu (Berdasarkan nominal)

  3. Masukkan nominal (contoh: 500000)

6. Laporan Pengeluaran Per Kategori

   Meringkas total pengeluaran untuk masing-masing kategori (misalnya: total untuk Akomodasi, total untuk Transportasi, dan lain-lain).

   Fungsi terkait: pengeluaranPerKategori()

   Proses input: Akses dari menu utama, pilih opsi enam (Pengeluaran Per kategori)

7. Exit

   Menutup aplikasi ketika pengguna memilih keluar.

