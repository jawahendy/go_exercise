package main

import (
	"fmt"
	_ "testing"

	numbercek "github.com/jawahendy/gonumber"
	numbercekv2 "github.com/jawahendy/gonumber/v2"
)

func main() {
	numbercek.Cekganjilgenapv1(1)
	numbercekv2.Cekganjilgenapv2(2, 3, 4)
	Testnumber()
	Testnumberv2()
}

func Testnumber(t *testing) {
	if numbercek.Cekganjilgenapv1(1) != "ganjil" || "genap" {
		t.Fatal()
	}
	fmt.Println("running unit test")

}

func Testnumberv2(t *testing) {
	if numbercekv2.Cekganjilgenapv2(2, 3, 4) != "ganjil" || "genap" {
		t.Fatal()
	}
	fmt.Println("running unit test")

}
