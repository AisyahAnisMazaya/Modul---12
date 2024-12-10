*NOMOR 1*
Deskripsi Program
Program ini membaca sejumlah n daerah yang memiliki jumlah rumah kerabat tertentu. Setiap daerah memiliki daftar nomor rumah yang diurutkan menggunakan algoritma Selection Sort. Program ini memastikan input valid berdasarkan batasan tertentu dan mencetak daftar nomor rumah yang sudah terurut untuk setiap daerah.
Penjelasan Setiap Bagian dan Variabel
1. Fungsi SelectionSort
Fungsi ini mengimplementasikan algoritma Selection Sort untuk mengurutkan elemen-elemen dalam array secara ascending.
Parameter: arr *[]int adalah pointer ke array yang akan diurutkan.
Variabel lokal:
i: Indeks awal untuk iterasi outer loop.
minIdx: Indeks elemen dengan nilai terkecil saat ini.
j: Indeks untuk iterasi inner loop.
2. Konstanta
maksimalN: Batas maksimum jumlah daerah yang dapat diproses (1 hingga 1000).
maksimalM: Batas maksimum jumlah rumah dalam setiap daerah (1 hingga 1.000.000).
3. Variabel Global
m: Menyimpan jumlah rumah kerabat pada daerah tertentu (diinputkan oleh pengguna).
n: Menyimpan jumlah total daerah (diinputkan oleh pengguna).
4. Variabel Lokal di main
rmhKerabat: Array 2 dimensi untuk menyimpan nomor rumah di setiap daerah.
Tipe: [][]int.
Contoh: Jika ada 3 daerah dengan masing-masing 2 rumah, array-nya mungkin seperti [[12, 5], [3, 9], [8, 7]].
panjangDaerah: Array untuk menyimpan jumlah rumah di setiap daerah.
Tipe: []int.
Contoh: Jika ada 3 daerah dengan jumlah rumah [2, 3, 1], maka panjangDaerah = [2, 3, 1].
Alur Program
Input Jumlah Daerah
Program membaca jumlah daerah n dari pengguna.
Jika n tidak dalam rentang [1, maksimalN], program akan berhenti.
Iterasi Setiap Daerah
Membaca jumlah rumah kerabat (m) untuk daerah ke-i.
Jika m tidak valid (tidak dalam [1, maksimalM]), program akan melewati daerah tersebut.
Membaca m nomor rumah dan menyimpannya dalam rmhKerabat[i].
Mengurutkan nomor rumah dalam rmhKerabat[i] menggunakan Selection Sort.
Output
Program mencetak daftar nomor rumah yang sudah terurut untuk setiap daerah.

*NOMOR 2*
Deskripsi Program
Program ini membaca n daerah yang masing-masing memiliki m rumah kerabat. Setiap nomor rumah dipisahkan menjadi ganjil dan genap, lalu:
Nomor rumah ganjil diurutkan secara menaik.
Nomor rumah genap diurutkan secara menurun.
Hasil kedua daftar digabungkan dan dicetak.
Penjelasan Setiap Bagian dan Variabel
1. Fungsi slcSortAsc
Fungsi ini mengurutkan elemen-elemen array secara menaik menggunakan algoritma Selection Sort.
Parameter: arr []int adalah array yang akan diurutkan.
Variabel Lokal:
i: Indeks iterasi utama.
minIdx: Indeks elemen terkecil yang ditemukan.
j: Indeks iterasi untuk mencari elemen terkecil.
2. Fungsi slcSortDsc
Fungsi ini mengurutkan elemen-elemen array secara menurun menggunakan algoritma Selection Sort.
Parameter: arr []int adalah array yang akan diurutkan.
Variabel Lokal:
i: Indeks iterasi utama.
maxIdx: Indeks elemen terbesar yang ditemukan.
j: Indeks iterasi untuk mencari elemen terbesar.
3. Konstanta
maxN: Batas maksimum jumlah daerah (1 hingga 1000).
maxM: Batas maksimum jumlah rumah dalam setiap daerah (1 hingga 1.000.000).
4. Variabel Lokal di main
n: Jumlah daerah yang akan diproses.
m: Jumlah rumah kerabat pada daerah tertentu.
rmhKerabat: Array 2 dimensi untuk menyimpan daftar nomor rumah setiap daerah.
Tipe: [][]int.
Contoh: Jika ada 2 daerah dengan daftar nomor rumah [1, 2, 3] dan [4, 5], maka rmhKerabat = [[1, 2, 3], [4, 5]].
ganjil: Array sementara untuk menyimpan nomor rumah ganjil pada daerah tertentu.
Tipe: []int.
genap: Array sementara untuk menyimpan nomor rumah genap pada daerah tertentu.
Tipe: []int.
Alur Program
Input Jumlah Daerah
Membaca jumlah daerah n dari pengguna.
Jika n tidak valid (tidak dalam rentang [1, maxN]), program berhenti.
Iterasi Setiap Daerah
Membaca jumlah rumah kerabat (m) untuk daerah ke-i.
Jika m tidak valid (tidak dalam rentang [1, maxM]), program berhenti.
Memisahkan nomor rumah ke dalam array ganjil dan genap berdasarkan paritas.
Mengurutkan:
ganjil secara menaik menggunakan slcSortAsc.
genap secara menurun menggunakan slcSortDsc.
Menggabungkan array ganjil dan genap menggunakan append.
Output
Program mencetak daftar nomor rumah (gabungan ganjil dan genap) untuk setiap daerah dalam satu baris.

*NOMOR 3*
Deskripsi Program
Program ini mengimplementasikan penghitungan median dari sekumpulan bilangan yang diinputkan secara dinamis. Median dihitung setiap kali pengguna memasukkan angka 0. Input akan berhenti jika pengguna memasukkan angka khusus -5313.
Penjelasan Fungsi dan Variabel
Fungsi InsertionSort
Fungsi ini mengurutkan array menggunakan algoritma Insertion Sort.
Parameter: arr []int, array yang akan diurutkan.
Variabel lokal:
i: Indeks elemen yang akan diproses.
key: Elemen yang sedang dibandingkan dengan elemen sebelumnya.
j: Indeks untuk elemen-elemen sebelum i.
Fungsi hitungMedian
Fungsi ini menghitung median dari array yang sudah terurut:
Parameter: arr []int, array yang sudah diurutkan.
Kondisi:
Jika panjang array ganjil, median adalah elemen tengah.
Jika panjang array genap, median adalah rata-rata dua elemen tengah (dibulatkan ke bawah).
Konstanta
maksimalData: Batas maksimum jumlah elemen dalam array (1 juta). Jika melebihi batas, program berhenti untuk menghindari penggunaan memori berlebih.
Variabel Lokal di main
data: Array dinamis yang menyimpan bilangan yang diinputkan oleh pengguna.
Tipe: []int.
angka: Bilangan yang sedang diinputkan oleh pengguna.
Tipe: int.
Alur Program
Input Bilangan Secara Dinamis
Pengguna memasukkan bilangan:
Jika bilangan adalah 0, array data diurutkan, median dihitung, dan dicetak.
Jika bilangan adalah -5313, program berhenti.
Selain itu, bilangan ditambahkan ke array data.
Validasi Jumlah Data
Jika jumlah elemen dalam array data melebihi maksimalData, program mencetak pesan kesalahan dan berhenti.
Proses Median
Jika bilangan 0 dimasukkan:
Array data diurutkan menggunakan InsertionSort.
Median dihitung menggunakan hitungMedian.
Median dicetak ke layar.
Program Berhenti
Program akan berhenti jika pengguna memasukkan -5313.
