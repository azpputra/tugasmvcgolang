package main

import (
	"fmt"
	"sepatu/model"
)

/*







func SearchSepatu(id int) {
	var tempLL *TableSepatu
	tempLL = DBsepatu.Next
	cek := false
	if DBsepatu.Next == nil {
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

func cariSepatu(id int) *TableSepatu {
	var tempLL *TableSepatu
	tempLL = DBsepatu.Next
	cek := false
	if DBsepatu.Next == nil {
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

func UpdateSepatu(id int, merk string, kondisi string) {
	sepatu := cariSepatu(id)
	if sepatu != nil {
		sepatu.Data.Merk = merk
		sepatu.Data.Kondisi = kondisi
		fmt.Println("update berhasil")
	} else {
		fmt.Println("tidak ada data yang dicari")
	}
}

func DeleteSepatu(id int) {
	var tempLL *TableSepatu
	tempLL = &DBsepatu
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

} */

func main() {
	model.InsertSepatu(1, "Nike", "New", 40)
	model.InsertSepatu(2, "Adidas", "New", 41)

	model.DeleteSepatu(2)
	sepatus := model.ReadAllSepatu()

	/* model.SearchSepatu(1) */
	/* model.UpdateSepatu(1, "Converse", "Second", 44)


	/* fmt.Println(sepatus)
	fmt.Println(sepatus.Next) */
	if sepatus != nil {
		for sepatus.Next != nil {
			fmt.Println(sepatus.Next.Data)
			sepatus = sepatus.Next
		}
	}
	/*
		 	InsertSepatu(1, "Nike", "New", 40)
			InsertSepatu(2, "Adidas", "New", 41)
			InsertSepatu(3, "NewBalance", "New", 39)
			InsertSepatu(4, "Vans", "Second", 42)
			//SearchSepatu(3)
			//UpdateSepatu(3, "Converse", "Second")
			DeleteSepatu(4)
			ReadAllSepatu()
	*/
}
