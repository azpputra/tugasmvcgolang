package model

import (
	"fmt"
	"sepatu/database"
	"sepatu/node"
)

func InsertSepatu(id int, merk string, kondisi string, ukuran float32) {
	newDataSepatu := node.TableSepatu{
		Data: node.Sepatu{id, merk, kondisi, ukuran},
		Next: nil,
	}
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if database.DBsepatu.Next == nil {
		database.DBsepatu.Next = &newDataSepatu
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataSepatu
	}
}
func SearchSepatu(id int) {
	var tempLL *node.TableSepatu
	tempLL = database.DBsepatu.Next
	cek := false
	if database.DBsepatu.Next == nil {
		fmt.Println("DATA SEPATU KOSONG")
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				fmt.Println("------------------------------------")
				fmt.Println("Nomer Sepatu :", tempLL.Data.Id)
				fmt.Println("Merk Sepatu :", tempLL.Data.Merk)
				fmt.Println("Kondisi Sepatu :", tempLL.Data.Kondisi)
				fmt.Println("Ukuran Sepatu :", tempLL.Data.Ukuran)
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			fmt.Println("ID tidak ditemukan")
		}
	}
}

func cariSepatu(id int) *node.TableSepatu {
	var tempLL *node.TableSepatu
	tempLL = database.DBsepatu.Next
	cek := false
	if database.DBsepatu.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateSepatu(id int, merk string, kondisi string, ukuran float32) {
	sepatu := cariSepatu(id)
	if sepatu != nil {
		sepatu.Data.Merk = merk
		sepatu.Data.Kondisi = kondisi
		sepatu.Data.Ukuran = ukuran
		fmt.Println("update berhasil")
	} else {
		fmt.Println("tidak ada data yang dicari")
	}
}

func ReadAllSepatu() *node.TableSepatu {
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if database.DBsepatu.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteSepatu(id int) {
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if tempLL.Next == nil {
		//cek data kosong
		fmt.Println("data kosong")
	} else {
		for tempLL.Next != nil {
			//fmt.Println(tempLL.Next.Data)
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return
			}
			tempLL = tempLL.Next
		}
	}

}
