package main

import "fmt"

type Dosen struct {
	nama, alamat, ttl, hp string
	id, nip               int
}

type Mahasiswa struct {
	id                         int
	nama, alamat, ttl, hp, nim string
}

type listDosen struct {
	data Dosen
	next *listDosen
}

type listMahasiswa struct {
	data Mahasiswa
	next *listMahasiswa
}

// TAMBAH DATA
func InsertData(list *List, customer Pelanggan) {
	tempList := List{}

	tempList.customer = customer

	temp := list

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
func ReadAll(list *List) {
	temp := list.next

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

func DeleteData(list *List, id int) {
	temp := list
	var counter int = 0
	for temp.next != nil {
		counter++
		if temp.next.customer.id == id {
			if counter == 1 {
				list.next = temp.next.next
				break
			} else if temp.next.next != nil {
				temp.next = temp.next.next
				break
			} else {
				temp.next = nil
				break
			}
		}
		temp = temp.next
	}
}

func inputUser() {
	var nama, alamat, ttl, hp string
	var nip int

	fmt.Println("Masukkan data dosen")
	fmt.Println("Nama\t: ")
	fmt.Scan(&nama)
	fmt.Println("Alamat\t: ")
	fmt.Scan(&alamat)
	fmt.Println("TTL\t: ")
	fmt.Scan(&ttl)
	fmt.Println("No. HP\t: ")
	fmt.Scan(&hp)
	fmt.Println("NIP\t: ")
	fmt.Scan(&nip)

	// data := Buku{judul, penulis, tahun}
	data := Dosen{nama, alamat, ttl, hp, id, nip}
	// atau data:= Buku{penulis : penulis, judul : judul , tahun : tahun}
	addDos(&dosen, data)
}
