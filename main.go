package main

import (
	"fmt"

	gotask "github.com/ebyantoo/module-task"
)

func main() {

	cek1 := gotask.CekGanjilGenap1(19)
	fmt.Println(cek1)

	cek2 := gotask.CekGanjilGenap2(5, 8, 9)
	fmt.Println(cek2)

}
