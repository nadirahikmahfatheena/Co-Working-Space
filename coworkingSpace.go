package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const kapasitas = 100

type CoWorkingSpace struct {
	Nama      string
	Fasilitas string
	Lokasi    string
	HargaSewa int
	Rating    float32
	Review    string
}

type daftarSpace [kapasitas]CoWorkingSpace

// sorting menggunakan selection sort berdasarkan hargaSewa (ascending)
func urutkanHarga(t *daftarSpace, n int) {
	for i := 0; i < n-1; i++ {
		minimum := i
		for j := i + 1; j < n; j++ {
			if t[j].HargaSewa < t[minimum].HargaSewa {
				minimum = j
			}
		}
		// Tukar elemen dengan harga sewa terkecil ke posisi ke-i
		t[i], t[minimum] = t[minimum], t[i]
	}
}

// sorting menggunakan insertion sort berdasarkan rating (descending)
func urutkanRating(t *daftarSpace, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i
		for j > 0 && temp.Rating > t[j-1].Rating {
			t[j] = t[j-1]
			j--
		}
		t[j] = temp
	}
}

// sorting menggunakan insertion sort berdasarkan alfabet dari nama
func urutkanNama(t *daftarSpace, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i
		for j > 0 && strings.ToLower(temp.Nama) < strings.ToLower(t[j-1].Nama) {
			t[j] = t[j-1]
			j--
		}
		t[j] = temp
	}
}

// binary search pada nama Co-Working Space (harus sudah disorting dulu)
func cariNama(t daftarSpace, n int, keyword string) int {
	kiri := 0
	kanan := n - 1
	keyword = strings.ToLower(keyword)

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(t[tengah].Nama)

		if namaTengah == keyword {
			return tengah
		} else if keyword < namaTengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

// menampilkan daftar Co-Working Space yang memiliki fasilitas tertentu
func filterFasilitas(t daftarSpace, n int) {
	var fasilitas string
	fmt.Print("Masukkan fasilitas yang ingin dicari: ")
	fmt.Scanln(&fasilitas)

	ketemu := false
	fmt.Printf("\nDaftar Co-Working Space dengan fasilitas \"%s\":\n", fasilitas)
	fmt.Printf("| %-3s | %-18s | %-35s | %-15s | %-10s | %-6s | %-50s |\n", "No", "Nama", "Fasilitas", "Lokasi", "HargaSewa", "Rating", "Review")
	fmt.Println("|-----|--------------------|-------------------------------------|-----------------|------------|--------|----------------------------------------------------|")

	no := 1
	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(t[i].Fasilitas), strings.ToLower(fasilitas)) {
			fmt.Printf("| %-3d | %-18s | %-35s | %-15s | %-10d | %-6.1f | %-50s |\n",
				no,
				t[i].Nama,
				t[i].Fasilitas,
				t[i].Lokasi,
				t[i].HargaSewa,
				t[i].Rating,
				t[i].Review,
			)
			no++
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ditemukan co-working space dengan fasilitas tersebut.")
	}
}

func cariSpace(t daftarSpace, n int, keyword string) int {
	for i := 0; i < n; i++ {
		if strings.EqualFold(t[i].Nama, keyword) || strings.EqualFold(t[i].Lokasi, keyword) {
			return i
		}
	}
	return -1
}

func tambahSpace(t *daftarSpace, n *int) {
	if *n >= kapasitas {
		fmt.Println("Kapasitas penuh, tidak bisa menambah Co-Working Space.")
		return
	}

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama Co-Working Space baru: ")
	input.Scan()
	nama := input.Text()

	fmt.Print("Masukkan fasilitas Co-Working Space: ")
	input.Scan()
	fasilitas := input.Text()

	fmt.Print("Masukkan lokasi Co-Working Space: ")
	input.Scan()
	lokasi := input.Text()

	var hargaInt int
	for {
		fmt.Print("Masukkan harga sewa Co-Working Space: ")
		input.Scan()
		harga := input.Text()
		if _, salah := fmt.Sscanf(harga, "%d", &hargaInt); salah != nil || hargaInt < 0 {
			fmt.Println("Input harga sewa tidak valid, gunakan angka positif.")
			continue
		}
		break
	}

	var ratingFloat float32
	for {
		fmt.Print("Masukkan rating Co-Working Space (1-5): ")
		input.Scan()
		rating := input.Text()
		if _, salah := fmt.Sscanf(rating, "%f", &ratingFloat); salah != nil || ratingFloat < 1 || ratingFloat > 5 {
			fmt.Println("Input rating tidak valid, gunakan angka 1-5.")
			continue
		}
		break
	}

	fmt.Print("Masukkan review Co-Working Space: ")
	input.Scan()
	review := input.Text()

	t[*n] = CoWorkingSpace{
		Nama:      nama,
		Fasilitas: fasilitas,
		Lokasi:    lokasi,
		HargaSewa: hargaInt,
		Rating:    ratingFloat,
		Review:    review,
	}
	*n = *n + 1

	fmt.Println("Data berhasil ditambah.")
}

// fungsi untuk mengedit data berdasarkan nama
func editSpace(t *daftarSpace, n int, nama string) {
	indeks := cariSpace(*t, n, nama)
	if indeks == -1 {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama baru: ")
	input.Scan()
	t[indeks].Nama = input.Text()

	fmt.Print("Masukkan fasilitas baru: ")
	input.Scan()
	t[indeks].Fasilitas = input.Text()

	fmt.Print("Masukkan lokasi baru: ")
	input.Scan()
	t[indeks].Lokasi = input.Text()

	fmt.Print("Masukkan harga sewa baru: ")
	input.Scan()
	hargaBaru := 0
	if _, salah := fmt.Sscanf(input.Text(), "%d", &hargaBaru); salah != nil { // err handling untuk memastikan input valid
		fmt.Println("Input harga sewa tidak valid, gunakan angka.")
		return
	}
	t[indeks].HargaSewa = hargaBaru

	fmt.Print("Masukkan rating baru (1-5): ")
	input.Scan()
	var ratingBaru float32 = 0
	if _, salah := fmt.Sscanf(input.Text(), "%f", &ratingBaru); salah != nil || ratingBaru < 1 || ratingBaru > 5 { // Meminta input rating baru dengan validasi rentang 1-10
		fmt.Println("Input rating tidak valid, gunakan angka 1-5.")
		return
	}
	t[indeks].Rating = ratingBaru

	fmt.Print("Masukkan review baru: ")
	input.Scan()
	t[indeks].Review = input.Text()

	fmt.Println("Data berhasil diperbarui.")
}

// fungsi untuk menghapus data berdasarkan nama
func hapusSpace(t *daftarSpace, n *int, nama string) {
	indeks := cariSpace(*t, *n, nama)
	if indeks == -1 {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}
	for i := indeks; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n = *n - 1
	fmt.Println("Co-Working Space berhasil dihapus.")
}

// menampilkan seluruh daftar Co-Working Space
func tampilDaftar(t daftarSpace, n int) {
	if n == 0 {
		fmt.Println("Belum ada Co-Working Space.")
		return
	}

	fmt.Printf("| %-3s | %-18s | %-35s | %-15s | %-10s | %-6s | %-50s |\n", "No", "Nama", "Fasilitas", "Lokasi", "HargaSewa", "Rating", "Review")
	fmt.Println("|-----|--------------------|-------------------------------------|-----------------|------------|--------|----------------------------------------------------|")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-3d | %-18s | %-35s | %-15s | %-10d | %-6.1f | %-50s |\n",
			i+1,
			t[i].Nama,
			t[i].Fasilitas,
			t[i].Lokasi,
			t[i].HargaSewa,
			t[i].Rating,
			t[i].Review,
		)
	}

}

func main() {
	var daftar daftarSpace
	jumlah := 20

	daftar[0] = CoWorkingSpace{"Milestone", "WiFi, AC, Ruang Meeting", "Jl.Sudagaran", 75000, 4.7, "Nyaman dan tenang"}
	daftar[1] = CoWorkingSpace{"Calf", "Kopi, Ruang Private", "Jl.Sudagaran", 60000, 4.5, "Suasana cozy"}
	daftar[2] = CoWorkingSpace{"StartUp Hub", "WiFi, Proyektor, Ruang Meeting", "Jl.Pahlawan", 85000, 4.8, "Tempat inspiratif untuk start-up"}
	daftar[3] = CoWorkingSpace{"TechNest", "WiFi, Ruang Kolaborasi, Parkir", "Jl.Tekno", 80000, 4.6, "Cocok untuk developer dan tech enthusiast"}
	daftar[4] = CoWorkingSpace{"CreativSpace", "WiFi, Whiteboard, Ruang Lounge", "Jl.Kreatif", 70000, 4.4, "Tempat untuk berkolaborasi dengan kreatif"}
	daftar[5] = CoWorkingSpace{"WorkHub", "WiFi, Dapur, Area Santai", "Jl.Semangat", 65000, 4.5, "Sesuai untuk kerja santai dan bertemu klien"}
	daftar[6] = CoWorkingSpace{"OfficeVibe", "AC, WiFi, Meja Kerja", "Jl.Pusat Bisnis", 90000, 4.3, "Menyediakan kenyamanan dan privasi"}
	daftar[7] = CoWorkingSpace{"GreenSpace", "WiFi, Ruang Meeting, Taman", "Jl.Ekologis", 75000, 4.7, "Lingkungan hijau dan tenang untuk bekerja"}
	daftar[8] = CoWorkingSpace{"SocialSpace", "WiFi, Event, Ruang Pertemuan", "Jl.Humanis", 80000, 4.6, "Fasilitas yang mendukung kolaborasi dan sosial"}
	daftar[9] = CoWorkingSpace{"CollabCorner", "WiFi, Printer, Dapur", "Jl.Agreement", 70000, 4.5, "Tempat untuk kolaborasi yang efektif"}
	daftar[10] = CoWorkingSpace{"InnovateLab", "WiFi, Ruang Rapat, Dapur", "Jl.Inovasi", 85000, 4.8, "Tempat untuk inovasi dan eksperimen"}
	daftar[11] = CoWorkingSpace{"BizHub", "WiFi, Meja Kerja, Ruang Presentasi", "Jl.Bisnis", 95000, 4.6, "Solusi untuk para pebisnis dan profesional"}
	daftar[12] = CoWorkingSpace{"CodeSpace", "WiFi, Ruang Coding, Server", "Jl.Programmer", 85000, 4.9, "Didesain untuk developer dan IT"}
	daftar[13] = CoWorkingSpace{"CreativityLab", "WiFi, Whiteboard, Ruang Galeri", "Jl.Art", 70000, 4.4, "Tempat untuk melahirkan ide-ide kreatif"}
	daftar[14] = CoWorkingSpace{"VisionSpace", "WiFi, Ruang Meeting, Projector", "Jl.Vision", 78000, 4.5, "Ideal untuk brainstorming dan diskusi kelompok"}
	daftar[15] = CoWorkingSpace{"FlexSpace", "WiFi, Meja Kerja, Ruang Santai", "Jl.Flexible", 68000, 4.3, "Ruang kerja fleksibel untuk berbagai kebutuhan"}
	daftar[16] = CoWorkingSpace{"StartupSpace", "WiFi, Meja Kerja, Ruang Kolaborasi", "Jl.Startup", 80000, 4.6, "Dikhususkan untuk para startup dan wirausahawan"}
	daftar[17] = CoWorkingSpace{"TechHub", "WiFi, Ruang Coding, Proyektor", "Jl.Teknologi", 88000, 4.7, "Tempat berkembangnya teknologi dan inovasi"}
	daftar[18] = CoWorkingSpace{"WorkSpaceX", "WiFi, Ruang Meeting, Dapur", "Jl.SpaceX", 77000, 4.5, "Ruang kerja nyaman untuk tim dan individu"}
	daftar[19] = CoWorkingSpace{"Collaborative", "WiFi, Printer, Ruang Acara", "Jl.Colab", 76000, 4.4, "Mendukung kolaborasi kreatif dalam berbagai proyek"}

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan Daftar Co-Working Space")
		fmt.Println("2. Tambah Co-Working Space")
		fmt.Println("3. Edit Co-Working Space")
		fmt.Println("4. Cari Co-Working Space")
		fmt.Println("5. Hapus Co-Working Space")
		fmt.Println("6. Urutkan berdasarkan Harga Sewa")
		fmt.Println("7. Urutkan berdasarkan Rating")
		fmt.Println("8. Filter Berdasarkan Fasilitas")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu (1-9): ")

		if !input.Scan() {
			break
		}
		pilihan := input.Text()

		switch pilihan {
		case "1":
			tampilDaftar(daftar, jumlah)
		case "2":
			tambahSpace(&daftar, &jumlah)
		case "3":
			//Sequential Search
			fmt.Print("Masukkan nama Co-Working Space yang ingin diedit: ")
			if input.Scan() {
				nama := input.Text()
				editSpace(&daftar, jumlah, nama)
			}
		case "4":
			//Binary Search Nama
			urutkanNama(&daftar, jumlah)
			fmt.Print("Masukkan nama Co-Working Space: ")
			if input.Scan() {
				nama := input.Text()
				indeks := cariNama(daftar, jumlah, nama)
				if indeks != -1 {
					fmt.Println("Data ditemukan:")
					fmt.Printf("Nama: %s\nLokasi: %s\nFasilitas: %s\nHarga: %d\nRating: %.1f\nReview: %s\n",
						daftar[indeks].Nama,
						daftar[indeks].Lokasi,
						daftar[indeks].Fasilitas,
						daftar[indeks].HargaSewa,
						daftar[indeks].Rating,
						daftar[indeks].Review)
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			}
		case "5":
			//Sequential Search
			fmt.Print("Masukkan nama Co-Working Space yang ingin dihapus: ")
			if input.Scan() {
				nama := input.Text()
				hapusSpace(&daftar, &jumlah, nama)
			}
		case "6":
			//Selection Sort
			urutkanHarga(&daftar, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan harga sewa.")
			tampilDaftar(daftar, jumlah)
		case "7":
			//Insertion Sort Descending
			urutkanRating(&daftar, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan rating.")
			tampilDaftar(daftar, jumlah)
		case "8":
			filterFasilitas(daftar, jumlah)
		case "9":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
