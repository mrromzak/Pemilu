package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX int = 100

type tCalon struct {
	nama   string
	partai string
	wni    string
	suara  int
}

type tPemilih struct {
	id           string
	sudahMemilih bool
	namaCalon 	 string
}

type tPartai struct {
	nama string
}

var votingStartTime, votingEndTime time.Time
var votingThreshold int

type tabPemilih [NMAX]tPemilih
type tabCalon [NMAX]tCalon
type tabPartai [NMAX]tPartai

func menuUtama() int {
	var pilihan int
	for {
		fmt.Println("")
        fmt.Println("Menu Utama:")
        fmt.Println("1. Masuk sebagai KPU")
        fmt.Println("2. Masuk sebagai Pemilih")
        fmt.Println("3. Keluar")
        fmt.Println("")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)
		return pilihan
	}
}

func menuKPU() int {
    var pilihan int
    for {
        fmt.Println("")
        fmt.Println("Menu KPU:")
        fmt.Println("1. Tambah Partai")
        fmt.Println("2. Tambah Calon")
        fmt.Println("3. Tambah Pemilih")
        fmt.Println("4. Edit data")
        fmt.Println("5. Periksa Data")
        fmt.Println("6. Atur Waktu Pemilihan")
        fmt.Println("7. Tentukan Ambang Batas Suara")
        fmt.Println("10. Kembali ke menu utama")
        fmt.Println("")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)
        return pilihan
    }
}

func setVotingThreshold() {
    fmt.Print("Masukkan ambang batas suara: ")
    fmt.Scan(&votingThreshold)
    fmt.Printf("Ambang batas suara telah diatur menjadi %d\n", votingThreshold)
}

func menuPeriksaData() int {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("1. Tampilkan hasil perolehan suara")
		fmt.Println("2. Cari NIK pemilih berdasarkan pilihan calon")
		fmt.Println("3. Cari calon berdasarkan partai")
		fmt.Println("4. Cari data calon berdasarkan nama")
		fmt.Println("5. Periksa NIK pemilih terdaftar")
		fmt.Println("10. Kembali")
		fmt.Println("")
		fmt.Print("Masukkan input: ")
		fmt.Scan(&pilihan)
		return pilihan
	}
}

func menuPemilih() int {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("Menu Pemilih:")
		fmt.Println("1. Pilih Calon")
		fmt.Println("2. Cari calon berdasarkan partai")
		fmt.Println("10. Kembali ke menu utama")
		fmt.Println("")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		return pilihan
	}
}

func menuEdit() int {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("1. Edit data calon")
		fmt.Println("2. Hapus data calon")
		fmt.Println("3. Edit data pemilih")
		fmt.Println("4. Hapus data pemilih")
		fmt.Println("5. Edit data partai")
		fmt.Println("6. Hapus data partai")
		fmt.Println("10. Kembalii")
		fmt.Println("")
		fmt.Print("Masukkan input: ")
		fmt.Scan(&pilihan)
		return pilihan
	}
}

func setVotingTime() {
	var startYear, startMonth, startDay, startHour, startMinute int
	var endYear, endMonth, endDay, endHour, endMinute int

	fmt.Println("Masukkan waktu mulai pemilihan:")
	fmt.Print("Tahun: ")
	fmt.Scan(&startYear)
	fmt.Print("Bulan: ")
	fmt.Scan(&startMonth)
	fmt.Print("Hari: ")
	fmt.Scan(&startDay)
	fmt.Print("Jam: ")
	fmt.Scan(&startHour)
	fmt.Print("Menit: ")
	fmt.Scan(&startMinute)

	votingStartTime = time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinute, 0, 0, time.Local)

	fmt.Println("Masukkan waktu akhir pemilihan:")
	fmt.Print("Tahun: ")
	fmt.Scan(&endYear)
	fmt.Print("Bulan: ")
	fmt.Scan(&endMonth)
	fmt.Print("Hari: ")
	fmt.Scan(&endDay)
	fmt.Print("Jam: ")
	fmt.Scan(&endHour)
	fmt.Print("Menit: ")
	fmt.Scan(&endMinute)

	votingEndTime = time.Date(endYear, time.Month(endMonth), endDay, endHour, endMinute, 0, 0, time.Local)

	fmt.Println("Waktu pemilihan berhasil diatur.")
	fmt.Printf("Waktu mulai: %s\n", votingStartTime)
	fmt.Printf("Waktu akhir: %s\n", votingEndTime)
}

func isVotingTime() bool {
	currentTime := time.Now()
	return currentTime.After(votingStartTime) && currentTime.Before(votingEndTime)
}

func editDataPartai(A *tabPartai, nCalon int) bool {
	var i int
	var namaPartai, gantiNama string
	fmt.Print("Masukkan nama partai yang ingin di-edit: ")
	fmt.Scan(&namaPartai)
	namaPartai = strings.ToLower(namaPartai)
	for i = 0; i < nCalon; i++ {
		if strings.ToLower(A[i].nama) == namaPartai {
			fmt.Print("Masukkan nama baru untuk partai tersebut: ")
			fmt.Scan(&gantiNama)
			A[i].nama = gantiNama
			return true
		}
	}
	return false
}


func editDataPemilih (A *tabPemilih, nPemilih int, idPemilih string) bool {
	var i int
	var gantiId string
	fmt.Print("Masukkan NIK pemilih yang ingin di-edit: ")
	fmt.Scan(&idPemilih)
	idPemilih = strings.ToLower(idPemilih)
	for i = 0; i < nPemilih; i++ {
		if strings.ToLower(A[i].id) == idPemilih {
			fmt.Print("Masukkan NIK baru untuk pemilih tersebut: ")
			fmt.Scan(&gantiId)
			A[i].id = gantiId
			return true
		}
	}
	return false
}

func tambahCalon(A *tabCalon, B tabPartai, nPartai int, n *int) {
    var i, j, tambahan int

    fmt.Println("")
    fmt.Print("Masukkan jumlah calon yang akan di-inputkan: ")
    fmt.Scan(&tambahan)
    fmt.Println("")

    if *n+tambahan > NMAX {
        tambahan = NMAX - *n
    }

    for i = 0; i < tambahan; i++ {
        var calon tCalon
        fmt.Printf("Masukkan nama calon ke-%d: ", i+1)
        fmt.Scan(&calon.nama)
        fmt.Print("Masukkan nama partai dari calon tersebut: ")
        fmt.Scan(&calon.partai)
        calon.partai = strings.ToLower(calon.partai)

        partaiDitemukan := false
        for j = 0; j < nPartai; j++ {
            if strings.ToLower(B[j].nama) == calon.partai {
                partaiDitemukan = true
            }
        }

        if partaiDitemukan {
            fmt.Print("Apakah calon tersebut merupakan WNI? (y/n): ")
            fmt.Scan(&calon.wni)
            if calon.wni != "y" {
                fmt.Println("Gagal melakukan input, calon bukan WNI")
            } else {
                A[*n + i] = calon
                fmt.Println("Data telah tersimpan")
                fmt.Println("")
            }
        } else {
            fmt.Println("")
            fmt.Println("Tidak ada nama partai tersebut")
        }
    }

    *n += tambahan

    insertionSortByName(A, *n)

    fmt.Println("")
    fmt.Println("Data calon saat ini:")
    for i = 0; i < *n; i++ {
        fmt.Printf("%d. Nama calon: %s, Nama partai: %s, WNI: %s, Suara: %d\n", i+1, A[i].nama, A[i].partai, A[i].wni, A[i].suara)
    }
}

func hapusDataCalon(A *tabCalon, nCalon *int, namaCalon string) bool {
	var i, j int
	fmt.Print("Masukkan nama calon yang ingin dihapus datanya: ")
	fmt.Scan(&namaCalon)
	namaCalon = strings.ToLower(namaCalon)
	for i = 0; i < *nCalon; i++ {
		if strings.ToLower(A[i].nama) == namaCalon {
			for j = i; j < *nCalon-1; j++ {
				A[j] = A[j+1]
			}
			*nCalon--
			return true
		}
	}
	fmt.Printf("Calon dengan nama %s tidak ditemukan\n", namaCalon)
	return false
}

func hapusDataPemilih(A *tabPemilih, nPemilih *int, idPemilih string) bool {
	var i, j int
	fmt.Print("Masukkan NIK pemilih yang ingin dihapus datanya: ")
	fmt.Scan(&idPemilih)
	for i = 0; i < *nPemilih; i++ {
		if A[i].id == idPemilih {
			for j = i; j < *nPemilih-1; j++ {
				A[j] = A[j+1]
			}
			*nPemilih--
			return true
		}
	}
	fmt.Printf("NIK pemilih dengan ID %s tidak ditemukan\n", idPemilih)
	return false
}

func hapusDataPartai(A *tabPartai, nPartai *int, namaPartai string) bool {
	var i, j int
	fmt.Print("Masukkan nama partai yang ingin dihapus datanya: ")
	fmt.Scan(&namaPartai)
	namaPartai = strings.ToLower(namaPartai)
	for i = 0; i < *nPartai; i++ {
		if strings.ToLower(A[i].nama) == namaPartai {
			for j = i; j < *nPartai-1; j++ {
				A[j] = A[j+1]
			}
			*nPartai--
			return true
		}
	}
	fmt.Printf("Partai dengan nama %s tidak ditemukan\n", namaPartai)
	return false
}

func tambahPemilih(A *tabPemilih, nPemilih *int) {
	var i, tambahan int
	fmt.Println("")
	fmt.Print("Masukkan jumlah pemilih yang akan di-inputkan:  ")
	fmt.Scan(&tambahan)
	if *nPemilih+tambahan > NMAX {
		tambahan = NMAX - *nPemilih
	}
	for i = 0; i < tambahan; i++ {
		fmt.Print("Masukkan NIK pemilih: ")
		fmt.Scan(&A[*nPemilih+i].id)
		A[*nPemilih+i].sudahMemilih = false
	}
	fmt.Println("Data telah tersimpan")
	fmt.Println("")
	for i = 0; i < *nPemilih; i++ {
		fmt.Printf("NIK: %s\n", A[i].id)
	}
	*nPemilih += tambahan
}

func periksaNikPemilih(A tabPemilih, nPemilih int) {
	var i, j, minIdx int
	for i = 0; i < nPemilih; i++ {
		minIdx = i
		for j = i + 1; j < nPemilih; j++ {
			if A[j].id < A[minIdx].id {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
	fmt.Println("")
	fmt.Println("NIK terdaftar: ")
	for i = 0; i < nPemilih; i++ {
		fmt.Printf("%d. NIK: %s, Status memilih: %v\n", i+1, A[i].id, A[i].sudahMemilih)
	}
}

func tambahPartai(A *tabPartai, nPartai *int) {
	var i, tambahan int
	fmt.Println("")
	fmt.Print("Masukkan jumlah partai yang akan di-inputkan: ")
	fmt.Scan(&tambahan)
	if *nPartai+tambahan > NMAX {
		tambahan = NMAX - *nPartai
	}
	for i = 0; i < tambahan; i++ {
		fmt.Printf("Masukkan nama partai ke-%d: ", *nPartai+i+1)
		fmt.Scan(&A[*nPartai+i].nama)
		A[*nPartai+i].nama = strings.ToLower(A[*nPartai+i].nama)
	}
	fmt.Println("")
	fmt.Println("Data telah tersimpan")
	for i = 0; i < *nPartai+tambahan; i++ {
		fmt.Printf("Nama partai ke-%d: %s\n", i+1, A[i].nama)
	}
	*nPartai += tambahan
	return
}

func insertionSortByName(A *tabCalon, n int) {
	var i, j int
	var key tCalon
	for i = 1; i < n; i++ {
		key = A[i]
		j = i - 1
		for j >= 0 && strings.ToLower(A[j].nama) > strings.ToLower(key.nama) {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func cariCalonBerdasarkanPartai(A tabCalon, nCalon int, partai string) {
	var ditemukan bool
	var i, x int

	x = 0
	fmt.Println("")
	fmt.Print("Masukkan nama partai yang ingin dicari calonnya: ")
	fmt.Scan(&partai)
	partai = strings.ToLower(partai)
	ditemukan = false

	insertionSortByName(&A, nCalon)

	for i = 0; i < nCalon; i++ {
		if strings.ToLower(A[i].partai) == partai {
			x++
		}
	}
	fmt.Println("")
	fmt.Printf("Terdapat %d calon dari partai %s\n", x, partai)
	for i = 0; i < nCalon; i++ {
		if strings.ToLower(A[i].partai) == partai {
			fmt.Printf("%d. Nama Calon: %s, Partai: %s, WNI: %s, Suara: %d\n", i+1, A[i].nama, A[i].partai, A[i].wni, A[i].suara)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Calon dari partai tersebut tidak ditemukan.")
	}
	return
}

func editDataCalon(A *tabCalon, n int, namaCalon string) bool {
	var i int
	var gantiNama string
	fmt.Print("Masukkan nama calon yang ingin di-edit: ")
	fmt.Scan(&namaCalon)
	namaCalon = strings.ToLower(namaCalon)
	for i = 0; i < n; i++ {
		if strings.ToLower(A[i].nama) == namaCalon {
			fmt.Print("Masukkan nama baru untuk calon tersebut: ")
			fmt.Scan(&gantiNama)
			A[i].nama = gantiNama
			return true
		}
	}
	return false
}

func tampilkanHasil(A tabCalon, nCalon int) {
    var i, j, minIdx int
    for i = 0; i < nCalon; i++ {
        minIdx = i
        for j = i + 1; j < nCalon; j++ {
            if A[j].suara > A[minIdx].suara {
                minIdx = j
            }
        }
        A[i], A[minIdx] = A[minIdx], A[i]
    }
    fmt.Println("")
    fmt.Println("Hasil perolehan suara:")
    for i = 0; i < nCalon; i++ {
        if A[i].suara >= votingThreshold {
            fmt.Printf("%d. Nama calon: %s, Partai: %s, Suara: %d (Memenuhi Ambang Batas)\n", i+1, A[i].nama, A[i].partai, A[i].suara)
        } else {
            fmt.Printf("%d. Nama calon: %s, Partai: %s, Suara: %d\n", i+1, A[i].nama, A[i].partai, A[i].suara)
        }
    }
}


func pilihCalon(A *tabPemilih, B *tabCalon, nPemilih int, nCalon int) {
	var NIK, pilih, pilih1 string
	var pemilihIndex int
	var foundPemilih, foundCalon bool

	for {
		fmt.Println("")
		fmt.Print("Masukkan NIK: ")
		fmt.Scan(&NIK)

		foundPemilih = false
		for i := 0; i < nPemilih; i++ {
			if A[i].id == NIK {
				pemilihIndex = i
				foundPemilih = true
				if A[i].sudahMemilih {
					fmt.Println("Hanya dapat memilih satu kali untuk setiap NIK")
					return
				}
			}
		}

		if !foundPemilih {
			fmt.Println("NIK tidak terdaftar")
			return
		}

		fmt.Println("")
		fmt.Print("Masukkan nama calon yang ingin dipilih: ")
		fmt.Scan(&pilih)
		pilih = strings.ToLower(pilih)

		foundCalon = false
		for i := 0; i < nCalon; i++ {
			if strings.ToLower(B[i].nama) == pilih {
				B[i].suara++
				foundCalon = true
				A[pemilihIndex].sudahMemilih = true
				A[pemilihIndex].namaCalon = B[i].nama
			}
		}

		if foundCalon {
			fmt.Println("Pilihan Anda telah tersimpan")
		} else {
			fmt.Println("Calon dengan nama tersebut tidak ditemukan")
		}

		fmt.Print("Lakukan pemilihan kembali? (y/n): ")
		fmt.Scan(&pilih1)
		if pilih1 != "y" {
			return
		}
	}
}

func cariCalon(A tabCalon, nCalon int, namaCalon string) {
	var i int
	fmt.Println("")
	fmt.Print("Masukkan nama calon yang ingin dicari: ")
	fmt.Scan(&namaCalon)
	namaCalon = strings.ToLower(namaCalon)
	for i = 0; i < nCalon; i++ {
		if strings.ToLower(A[i].nama) == namaCalon {
			if A[i].suara >= votingThreshold {
				fmt.Println("Calon ditemukan!")
				fmt.Printf("Nama calon: %s, Partai: %s, Suara: %d (Memenuhi ambang batas)\n", A[i].nama, A[i].partai, A[i].suara)
			} else {
				fmt.Println("Calon ditemukan!")
				fmt.Printf("Nama calon: %s, Partai: %s, Suara: %d\n", A[i].nama, A[i].partai, A[i].suara)
			}
		} else {
			fmt.Println("Calon tidak ditemukan")
		}
	}
}


func pencarianPemilihBerdasarkanCalon(A tabPemilih, B tabCalon, nCalon int, namaCalon string) {
	var i, j, minIdx int
	for i = 0; i < nCalon; i++ {
		minIdx = i
		for j = i + 1; j < nCalon; j++ {
			if A[j].id < A[minIdx].id {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}

    fmt.Println("")
    fmt.Print("Masukkan nama calon yang ingin dicari pemilihnya: ")
    fmt.Scan(&namaCalon)
    namaCalon = strings.ToLower(namaCalon)

    fmt.Println("")
    fmt.Printf("Pemilih yang memilih calon %s:\n", namaCalon)
    var foundPemilih bool = false
    for i := 0; i < NMAX; i++ {
        if strings.ToLower(A[i].namaCalon) == namaCalon {
            fmt.Printf("%d. NIK: %s\n", i+1, A[i].id)
            foundPemilih = true
        }
    }
    if !foundPemilih {
        fmt.Println("Tidak ada pemilih yang memilih calon ini.")
    }
}


func main() {
    var calonS tabCalon
    var pemilihS tabPemilih
    var partaiS tabPartai
    var nCalon, nPemilih, nPartai int
    var pilihan, pilihan1, pilihan2, pilihan3, pilihan4 int
    var x string
	var exit bool

	for !exit {
        pilihan = menuUtama()
        switch pilihan {
		//menukpu
        case 1:
			var kembali bool
			for !kembali {
				pilihan1 = menuKPU()
				switch pilihan1 {
				case 1:
					tambahPartai(&partaiS, &nPartai)
				case 2:
					tambahCalon(&calonS, partaiS, nPartai, &nCalon)
				case 3:
					tambahPemilih(&pemilihS, &nPemilih)
				//menueditdata
				case 4:
					pilihan3 = menuEdit()
					switch pilihan3 {
					case 1:
						if editDataCalon(&calonS, nCalon, x) {
							fmt.Println("")
							fmt.Println("Nama calon berhasil di-edit!")
						} else {
							fmt.Println("Edit data gagal, nama calon tidak ditemukan")
						}
					case 2:
						if hapusDataCalon(&calonS, &nCalon, x) {
							fmt.Println("Data calon telah berhasil dihapus!")
						} else {
							fmt.Println("Penghapusan data calon gagal, nama calon tidak ditemukan.")
						}
					case 3:
						if editDataPemilih(&pemilihS, nPemilih, x) {
							fmt.Println("")
							fmt.Println("ID pemilih berhasil diganti")
						} else {
							fmt.Println("Edit ID gagal, ID pemilih tidak ditemukan")
						}
					case 4:
						if hapusDataPemilih(&pemilihS, &nPemilih, x) {
							fmt.Println("ID pemilih telah berhasil dihapus!")
						} else {
							fmt.Println("Penghapusan ID pemilih gagal, ID tidak ditemukan.")
						}
					case 5:
						if editDataPartai(&partaiS, nPartai) {
							fmt.Println("")
							fmt.Println("Nama partai berhasil diganti")
						} else {
							fmt.Println("Edit data gagal, nama partai tidak ditemukan")
						}
					case 6:
						if hapusDataPartai(&partaiS, &nPartai, x) {
							fmt.Println("Data partai telah berhasil dihapus!")
						} else {
							fmt.Println("Penghapusan data partai gagal, nama partai tidak ditemukan.")
						}
					case 10:
						kembali = true
					default:
						fmt.Println("Pilihan tidak valid")
					}
				//menuperiksadata
				case 5:
					pilihan4 = menuPeriksaData()
					switch pilihan4 {
					case 1:
						tampilkanHasil(calonS, nCalon)
					case 2:
						pencarianPemilihBerdasarkanCalon(pemilihS, calonS, nCalon, x)
					case 3:
						cariCalonBerdasarkanPartai(calonS, nCalon, x)
					case 4:
						cariCalon(calonS, nCalon, x)
					case 5:
						periksaNikPemilih(pemilihS, nPemilih)
					case 10:
						kembali = true
					default:
						fmt.Println("Pilihan tidak valid")
					}
				//menukpu
				case 6:
					setVotingTime()
				case 7:
					setVotingThreshold()
				case 10:
					kembali = true
				default:
					fmt.Println("Pilihan tidak valid")
				}
			}
		//menupemilih
		case 2:
			var kembali bool
			for !kembali {
				pilihan2 = menuPemilih()
				switch pilihan2 {
				case 1:
					if isVotingTime() {
						pilihCalon(&pemilihS, &calonS, nPemilih, nCalon)
					} else {
						fmt.Println("Saat ini bukan waktu pemilihan. Anda hanya bisa melihat daftar calon.")
					}
				case 2:
					cariCalonBerdasarkanPartai(calonS, nCalon, x)
				case 10:
					kembali = true
				default:
					fmt.Println("Pilihan tidak valid")
				}
			}
		//menukpu
		case 3:
			exit = true
		default:
			fmt.Println("Pilihan tidak valid")
        }	
	}
}
