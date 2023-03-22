package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

func (biodata *Biodata) alasanMemilihGolang() {
	fmt.Println("Budi suka golang")
}

func main() {
	args := os.Args
	fmt.Println(args)

	bio := Biodata{}

	budi := &Biodata{
		Nama:      "Budi",
		Alamat:    "Jakarta",
		Pekerjaan: "Programmer",
	}

	bio.alasanMemilihGolang()

	fmt.Println(budi)

}
