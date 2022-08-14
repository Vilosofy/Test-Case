package main

import "fmt"

// function buat cek apakah dalam array tersebut mengandung bilangan ganjil atau genap
//jika genap tampilkan bilangannya dan pesan (string "merupakan bil genap")
func checking(number[5] int){
	var i int
	var penjumlahan int
	for i = 0; i < len(number); i++{
		penjumlahan = penjumlahan + number[i]
	}

	if penjumlahan % 2 == 0{
		fmt.Printf("%d merupakan bilangan genap", penjumlahan)
	} else {
		fmt.Printf("%d merupakan bilangan ganjil", penjumlahan)
	}
}

func main(){
	var numeric[5] int
	numeric[0] = 1
	numeric[1] = 2
	numeric[2] = 3
	numeric[3] = 4

	var panjangString = len(numeric)
	fmt.Printf("Panjang String = %d\n", panjangString)
	checking(numeric)
}