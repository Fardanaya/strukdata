package main

import "fmt"

type Dosen struct {
	nama, alamat, ttl, hp string
	id, nip               int
}

type Mhs struct {
	id                         int
	nama, alamat, ttl, hp, nim string
}

type listDos struct {
	data Dosen
	next *listDos
}
type listMhs struct {
	data Mhs
	next *listMhs
}

// TAMBAH DATA
func addDos(chain *listDos, data Dosen) {
	var tempList listDos
	tempList.data = data

	temp := chain
	if temp.next == nil {
		temp.next = &tempList
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &tempList
	}
}

// LIHAT DATA SEMUA DOSEN
func viewAllDos(chain *listDos) {
	temp := chain.next

	for temp != nil {
		fmt.Println("-Data Dosen-")
		fmt.Println("-ID\t\t : ", temp.data.id)
		fmt.Println("-Nama\t : ", temp.data.nama)
		fmt.Println("-Alamat\t : ", temp.data.alamat)
		fmt.Println("-NIP\t : ", temp.data.nip)
		fmt.Println("-No. HP\t : ", temp.data.hp)
		fmt.Println("-TTL\t : ", temp.data.ttl)
		temp = temp.next
	}
}

// LIHAT DATA SEMUA DOSEN
func viewAllMhs(chain *listMhs) {
	temp := chain.next

	for temp != nil {
		fmt.Println("-Data Mahasiswa-")
		fmt.Println("-ID\t\t : ", temp.data.id)
		fmt.Println("-Nama\t : ", temp.data.nama)
		fmt.Println("-Alamat\t : ", temp.data.alamat)
		fmt.Println("-NIM\t : ", temp.data.nim)
		fmt.Println("-No. HP\t : ", temp.data.hp)
		fmt.Println("-TTL\t : ", temp.data.ttl)
		temp = temp.next
	}
}


func utama(){
	var in int
	for in !=3{
		fmt.Println("PENGISIAN DATA")
		fmt.Println("1. DOSEN")
		fmt.Println("2. MAHASISWA")
		fmt.Println("3. EXIT")
		fmt.Scan(&in)
		if in !=3 {
			menu(in)
		}
	}
}

func menu(x int){
	var in int
	for in != 6{
		switch x {
			case 1:
				fmt.Println("PENGISIAN DATA DOSEN")
			case 2:
				fmt.Println("PENGISIAN DATA MAHASISWA")
			}
			fmt.Println("1. Tambah data")
			fmt.Println("2. Hapus data")
			fmt.Println("3. Update data")
			fmt.Println("4. Lihat semua data")
			fmt.Println("5. Lihat dari ID")
			fmt.Println("6. Kembali ")
			if x==1{
				// punya dosen
				fmt.Scan(&in)
				switch in {
				case 1:
					// add()
					fmt.Println("add DOS")
				case 2:
					// del()
					fmt.Println("del DOS")
				case 3:
					// update()
					fmt.Println("up DOS")
				case 4:
					// viewAll()
					fmt.Println("viewall DOS")
				case 5:
					// viewID()
					fmt.Println("view DOS")
				}
			} else if x==2{
				// punya mahasiswa
				fmt.Scan(&in)
				switch in {
					case 1:
						// add()
						fmt.Println("add MHS")
					case 2:
						// del()
						fmt.Println("del MHS")
					case 3:
						// update()
						fmt.Println("up MHS")
					case 4:
						// viewAll()
						fmt.Println("viewall MHS")
					case 5:
						// viewID()
						fmt.Println("view MHS")
				}
			}
		}
	}

func main() {
	utama()
}
