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

func UpdateData(listDosen *structDosen, listMahasiswa *structMahasiswa, dosen Dosen, mahasiswa Mahasiswa, selector int) {
	tempDosen := listDosen
	tempMahasiswa := listMahasiswa

	if selector == 1 {
		for tempDosen.next != nil {
			if tempDosen.next.data.nip == dosen.nip {
				tempDosen.next.data = dosen
				break
			}
			tempDosen = tempDosen.next
		}
	} else {
		for tempMahasiswa.next != nil {
			if tempMahasiswa.next.data.nim == mahasiswa.nim {
				tempMahasiswa.next.data = mahasiswa
				break
			}
			tempMahasiswa = tempMahasiswa.next
		}
	}
}

func DeleteData(listDosen *structDosen, listMahasiswa *structMahasiswa, selector int, userId string) {
	tempDosen := listDosen
	tempMahasiswa := listMahasiswa

	var counter int = 0
	if selector == 1 {
		for tempDosen.next != nil {
			counter++
			if tempDosen.next.data.nip == userId {
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
			if tempMahasiswa.next.data.nim == userId {
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

// LIHAT DATA
func ReadData(listDosen *structDosen, listMahasiswa *structMahasiswa, selector int, sort string) {
	tempDosen := listDosen.next
	tempMahasiswa := listMahasiswa.next

	fmt.Print("View Data ")
	if selector == 1 {
		fmt.Println("Dosen")
	} else {
		fmt.Println("Mahasiswa")
	}

	if sort == "nil" {
		// View All
		if selector == 1 {
			for tempDosen != nil {
				fmt.Println("-ID\t : ", tempDosen.data.id)
				fmt.Println("-Nama\t : ", tempDosen.data.nama)
				fmt.Println("-Alamat\t : ", tempDosen.data.alamat)
				fmt.Println("-NIP\t : ", tempDosen.data.nip)
				fmt.Println("-No. HP\t : ", tempDosen.data.hp)
				fmt.Println("-TTL\t : ", tempDosen.data.ttl)
				tempDosen = tempDosen.next
			}
		} else {
			for tempMahasiswa != nil {
				fmt.Println("-ID\t : ", tempMahasiswa.data.id)
				fmt.Println("-Nama\t : ", tempMahasiswa.data.nama)
				fmt.Println("-Alamat\t : ", tempMahasiswa.data.alamat)
				fmt.Println("-NIM\t : ", tempMahasiswa.data.nim)
				fmt.Println("-No. HP\t : ", tempMahasiswa.data.hp)
				fmt.Println("-TTL\t : ", tempMahasiswa.data.ttl)
				tempMahasiswa = tempMahasiswa.next
			}
		}
	} else {
		// View By NIP or NIM
		if selector == 1 {
			for tempDosen != nil {
				if tempDosen.data.nip == sort {
					fmt.Println("-ID\t : ", tempDosen.data.id)
					fmt.Println("-Nama\t : ", tempDosen.data.nama)
					fmt.Println("-Alamat\t : ", tempDosen.data.alamat)
					fmt.Println("-NIP\t : ", tempDosen.data.nip)
					fmt.Println("-No. HP\t : ", tempDosen.data.hp)
					fmt.Println("-TTL\t : ", tempDosen.data.ttl)
				}
				tempDosen = tempDosen.next
			}
		} else {
			for tempMahasiswa != nil {
				if tempMahasiswa.data.nim == sort {
					fmt.Println("-ID\t : ", tempMahasiswa.data.id)
					fmt.Println("-Nama\t : ", tempMahasiswa.data.nama)
					fmt.Println("-Alamat\t : ", tempMahasiswa.data.alamat)
					fmt.Println("-NIM\t : ", tempMahasiswa.data.nim)
					fmt.Println("-No. HP\t : ", tempMahasiswa.data.hp)
					fmt.Println("-TTL\t : ", tempMahasiswa.data.ttl)
				}
				tempMahasiswa = tempMahasiswa.next
			}
		}
	}
}

func inputUser(selector int, model int) {

	var x string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if model == 1 {
		for x != "t" {
			fmt.Print("Masukkan Data ")
			if selector == 1 {
				fmt.Println("Dosen")
			} else {
				fmt.Println("Mahasiswa")
			}

			id++

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

			var nip, nim string
			if selector == 1 {
				fmt.Print("NIP\t: ")
				scanner.Scan()
				nip = scanner.Text()
			} else {
				fmt.Print("NIM\t: ")
				scanner.Scan()
				nim = scanner.Text()
			}

			// Data Skeleton
			dataDosen := Dosen{id, nama, alamat, ttl, hp, nip}
			dataMahasiswa := Mahasiswa{id, nama, alamat, ttl, hp, nim}

			// Insert Data
			InsertData(&dosen, &mahasiswa, dataDosen, dataMahasiswa, selector)

			// Break Looping
			fmt.Println("TAMBAH DATA ? [ y / t ]")
			scanner.Scan()
			x = strings.TrimSpace(scanner.Text())
		}
	} else {
		fmt.Print("Masukkan Data Baru ")
		if selector == 1 {
			fmt.Println("Dosen")
		} else {
			fmt.Println("Mahasiswa")
		}

		id++

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

		var nip, nim string
		if selector == 1 {
			fmt.Print("NIP\t: ")
			scanner.Scan()
			nip = scanner.Text()
		} else {
			fmt.Print("NIM\t: ")
			scanner.Scan()
			nim = scanner.Text()
		}

		// Data Skeleton
		dataDosen := Dosen{id, nama, alamat, ttl, hp, nip}
		dataMahasiswa := Mahasiswa{id, nama, alamat, ttl, hp, nim}

		// Update Data
		UpdateData(&dosen, &mahasiswa, dataDosen, dataMahasiswa, selector)
	}

}

func mainMenu() {
	var input int
	for input != 3 {
		fmt.Println("ITATS")
		fmt.Println("1. DOSEN")
		fmt.Println("2. MAHASISWA")
		fmt.Println("3. EXIT")
		fmt.Scan(&input)
		if input < 3 && input > 0 {
			chooseMenu(input)
		} else if input != 3 {
			fmt.Println("Invalid Input")
		} else {
			fmt.Println("Exiting Program")
		}
	}
}

func chooseMenu(selector int) {
	var input int
	var query string

	for input != 6 {
		fmt.Print("PENGOLAHAN DATA ")
		if selector == 1 {
			fmt.Println("DOSEN")
		} else {
			fmt.Println("MAHASISWA")
		}

		fmt.Println("1. Tambah Data")
		fmt.Println("2. Hapus Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Lihat Semua Data")
		fmt.Println("5. Lihat Data Menggunakan NIP / NIM")
		fmt.Println("6. Kembali")

		// Select Menu
		fmt.Scan(&input)

		switch input {
		case 1:
			if selector == 1 {
				inputUser(1, 1)
			} else {
				inputUser(2, 1)
			}
		case 2:
			if selector == 1 {
				fmt.Println("Masukkan NIP yang ingin dihapus")
				fmt.Scan(&query)
				DeleteData(&dosen, &mahasiswa, 1, query)
			} else {
				fmt.Println("Masukkan NIM yang ingin dihapus")
				fmt.Scan(&query)
				DeleteData(&dosen, &mahasiswa, 2, query)
			}
		case 3:
			if selector == 1 {
				fmt.Println("Masukkan NIP yang ingin di update")
				fmt.Scan(&query)
				ReadData(&dosen, &mahasiswa, 1, query)
				inputUser(1, 2)
			} else {
				fmt.Println("Masukkan NIM yang ingin di update")
				fmt.Scan(&query)
				ReadData(&dosen, &mahasiswa, 2, query)
				inputUser(2, 2)
			}
		case 4:
			if selector == 1 {
				ReadData(&dosen, &mahasiswa, 1, "nil")
			} else {
				ReadData(&dosen, &mahasiswa, 2, "nil")
			}
		case 5:
			if selector == 1 {
				fmt.Println("Masukkan NIP yang ingin dilihat")
				fmt.Scan(&query)
				ReadData(&dosen, &mahasiswa, 1, query)
			} else {
				fmt.Println("Masukkan NIM yang ingin dilihat")
				fmt.Scan(&query)
				ReadData(&dosen, &mahasiswa, 2, query)
			}
		}
	}
}

func main() {
	mainMenu()
}
