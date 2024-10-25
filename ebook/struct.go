package ebook

import "fmt"

type Buku struct {
	Judul string
	ISBN  int
	Harga int
}

func (b Buku) DisplayBook() {
	fmt.Println(b)
}

func (b *Buku) UbahHarga() {
	*b = Buku{
		Harga: 83000,
	}
	fmt.Println(b)
}

// package main

// import (
// 	"soundwave/ebook"
// )

// func main() {
// 	book := ebook.Buku{
// 		Judul: "The Ladybug",
// 		ISBN:  0001454700,
// 		Harga: 79000,
// 	}
// 	book.DisplayBook()
// 	book.UbahHarga()

// }
