package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dosen struct {
	id                         int
	nama, alamat, ttl, hp, nip string
}

type Mahasiswa struct {
	id                         int
	nama, alamat, ttl, hp, nim string
}

type structDosen struct {
	data Dosen
	next *structDosen
}

type structMahasiswa struct {
	data Mahasiswa
	next *structMahasiswa
}

// id counter
var id int = 0
var dosen structDosen
var mahasiswa structMahasiswa

// TAMBAH DATA
func InsertData(listDosen *structDosen, listMahasiswa *structMahasiswa, dosen Dosen, mahasiswa Mahasiswa, selector int) {
	var tempDosen structDosen
	var tempMahasiswa structMahasiswa

	if selector == 1 {
		tempDosen.data = dosen
		temp := listDosen

		if temp.next == nil {
			temp.next = &tempDosen
		} else {
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = &tempDosen
		}
	} else {
		tempMahasiswa.data = mahasiswa
		temp := listMahasiswa

		if temp.next == nil {
			temp.next = &tempMahasiswa
		} else {
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = &tempMahasiswa
		}
	}
}

// LIHAT DATA SEMUA DOSEN
func ReadAll(listDosen *structDosen, listMahasiswa *structMahasiswa, selector int) {
	tempDosen := listDosen.next
	tempMahasiswa := listMahasiswa.next
	if selector == 1 {
		for tempDosen != nil {
			fmt.Println("-Data Dosen-")
			fmt.Println("-ID\t\t : ", tempDosen.data.id)
			fmt.Println("-Nama\t : ", tempDosen.data.nama)
			fmt.Println("-Alamat\t : ", tempDosen.data.alamat)
			fmt.Println("-NIP\t : ", tempDosen.data.nip)
			fmt.Println("-No. HP\t : ", tempDosen.data.hp)
			fmt.Println("-TTL\t : ", tempDosen.data.ttl)
			tempDosen = tempDosen.next
		}
	} else {
		for tempMahasiswa != nil {
			fmt.Println("-Data Mahasiswa-")
			fmt.Println("-ID\t\t : ", tempMahasiswa.data.id)
			fmt.Println("-Nama\t : ", tempMahasiswa.data.nama)
			fmt.Println("-Alamat\t : ", tempMahasiswa.data.alamat)
			fmt.Println("-NIM\t : ", tempMahasiswa.data.nim)
			fmt.Println("-No. HP\t : ", tempMahasiswa.data.hp)
			fmt.Println("-TTL\t : ", tempMahasiswa.data.ttl)
			tempMahasiswa = tempMahasiswa.next
		}
	}
}

func DeleteData(listDosen *structDosen, listMahasiswa *structMahasiswa, selector int) {
	tempDosen := listDosen.next
	tempMahasiswa := listMahasiswa.next

	var counter int = 0
	if selector == 1 {
		for tempDosen.next != nil {
			counter++
			if tempDosen.next.data.id == selector {
				if counter == 1 {
					tempDosen.next = tempDosen.next.next
					break
				} else if tempDosen.next.next != nil {
					tempDosen.next = tempDosen.next.next
					break
				} else {
					tempDosen.next = nil
					break
				}
			}
			tempDosen = tempDosen.next
		}
	} else {
		for tempMahasiswa.next != nil {
			counter++
			if tempMahasiswa.next.data.id == selector {
				if counter == 1 {
					tempMahasiswa.next = tempMahasiswa.next.next
					break
				} else if tempMahasiswa.next.next != nil {
					tempMahasiswa.next = tempMahasiswa.next.next
					break
				} else {
					tempMahasiswa.next = nil
					break
				}
			}
			tempMahasiswa = tempMahasiswa.next
		}
	}
}

func inputUser(selector int) {
	//var nama, alamat, ttl, hp, nip, nim string
	var x string
	scanner := bufio.NewScanner(os.Stdin)
	for x != "t" {
		if selector == 1 {
			fmt.Println("Masukkan data dosen")
		} else {
			fmt.Println("Masukkan data mahasiswa")
		}
		fmt.Print("Nama\t: ")
		scanner.Scan()
		nama := scanner.Text()
		fmt.Print("Alamat\t: ")
		scanner.Scan()
		alamat := scanner.Text()
		fmt.Print("TTL\t: ")
		scanner.Scan()
		ttl := scanner.Text()
		fmt.Print("No. HP\t: ")
		scanner.Scan()
		hp := scanner.Text()
		fmt.Print("NIP\t: ")
		scanner.Scan()
		ni := scanner.Text()
		dataDosen := Dosen{id, nama, alamat, ttl, hp, ni}
		dataMahasiswa := Mahasiswa{id, nama, alamat, ttl, hp, ni}
		InsertData(&dosen, &mahasiswa, dataDosen, dataMahasiswa, selector)

		fmt.Println("TAMBAH DATA ?(y/t)")
		scanner.Scan()
		x = strings.TrimSpace(scanner.Text())
		// if x == "t" {
		// 	break
		// }

	}

}

func utama() {
	var in int
	for in != 3 {
		fmt.Println("PENGISIAN DATA")
		fmt.Println("1. DOSEN")
		fmt.Println("2. MAHASISWA")
		fmt.Println("3. EXIT")
		fmt.Scan(&in)
		if in != 3 {
			menu(in)
		}
	}
}

func menu(x int) {
	var in int
	for in != 6 {
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
		if x == 1 {
			// punya dosen
			fmt.Scan(&in)
			switch in {
			case 1:
				// fmt.Println("add DOS")
				inputUser(1)
			case 2:
				// del()
				fmt.Println("del DOS")
			case 3:
				// update()
				fmt.Println("up DOS")
			case 4:
				// fmt.Println("viewall DOS")
				ReadAll(&dosen, &mahasiswa, 1)
			case 5:
				// viewID()
				fmt.Println("view DOS")
			}
		} else if x == 2 {
			// punya mahasiswa
			fmt.Scan(&in)
			switch in {
			case 1:
				// fmt.Println("add MHS")
				inputUser(2)
			case 2:
				// del()
				fmt.Println("del MHS")
			case 3:
				// update()
				fmt.Println("up MHS")
			case 4:
				// fmt.Println("viewall MHS")
				ReadAll(&dosen, &mahasiswa, 2)
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
