package main

import (
	"fmt"
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
			fmt.Println("-Data Dosen-")
			fmt.Println("-ID\t\t : ", tempMahasiswa.data.id)
			fmt.Println("-Nama\t : ", tempMahasiswa.data.nama)
			fmt.Println("-Alamat\t : ", tempMahasiswa.data.alamat)
			fmt.Println("-NIP\t : ", tempMahasiswa.data.nim)
			fmt.Println("-No. HP\t : ", tempMahasiswa.data.hp)
			fmt.Println("-TTL\t : ", tempMahasiswa.data.ttl)
			tempMahasiswa = tempMahasiswa.next
		}
	}
}

func DeleteData(listDosen *structDosen, listMahasiswa *structMahasiswa, id int, selector int) {
	tempDosen := listDosen.next
	tempMahasiswa := listMahasiswa.next

	var counter int = 0
	if selector == 1 {
		for tempDosen.next != nil {
			counter++
			if tempDosen.next.data.id == id {
				if tempDosen.next.next != nil {
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
			if tempMahasiswa.next.data.id == id {
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
	var nama, alamat, ttl, hp, nip, nim string

	if selector == 1 {
		fmt.Println("Masukkan data dosen")
		fmt.Print("Nama\t: ")
		fmt.Scan(&nama)
		fmt.Print("Alamat\t: ")
		fmt.Scan(&alamat)
		fmt.Print("TTL\t: ")
		fmt.Scan(&ttl)
		fmt.Print("No. HP\t: ")
		fmt.Scan(&hp)
		fmt.Print("NIP\t: ")
		fmt.Scan(&nip)

		dataDosen := Dosen{id, nama, alamat, ttl, hp, nip}
		dataMahasiswa := Mahasiswa{id, nama, alamat, ttl, hp, nim}
		InsertData(&dosen, nil, dataDosen, dataMahasiswa, 1)
	}
}

func main() {
	fmt.Println("Test")

	for {

		inputUser(1)
		var x string
		fmt.Print("Tambah Data Lagi ? ( y / n ) \n-> ")
		fmt.Scan(&x)

		if x != "y" {
			break
		}

		fmt.Print("\n")
	}
	inputUser(1)
	DeleteData(&dosen, &mahasiswa, 2, 1)
	ReadAll(&dosen, &mahasiswa, 1)

}
