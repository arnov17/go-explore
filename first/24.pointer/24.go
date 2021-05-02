package main

import "fmt"

type Address struct {
	City, Province, Country string
}

//pointer di func
func changeToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Man struct {
	Name string
}

// pointer di method
func (man *Man) calling() {
	man.Name = "Mr " + man.Name
}

func main() {
	address1 := Address{"Subang", "jabar", "indonesia"}
	address2 := address1
	address3 := &address1
	address2.City = "bandung"
	address3.City = "Sumedang"

	address4 := &address1
	// tanpa* addres1 tetap value nya diawal, namun addres4 membuat memory baru sendiri
	address4 = &Address{"malang", "jatim", "indonesia"}
	// dengan * membuat addres1 menjadi seperti memory address4
	// *address4 = Address{"malang", "jatim", "indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	// new pointer kosong
	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println(address5)

	var alamat = Address{
		City:     "semarang",
		Province: "jateng",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	changeToIndonesia(alamatPointer)
	fmt.Println(alamat)

	// pointer di method
	eko := Man{"eko"}
	eko.calling()

	fmt.Println(eko.Name)
}
