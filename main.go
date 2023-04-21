package main

import (
	"fmt"
	"github.com/Lordkissa97/is105sem03/mycrypt"
)

func main() {

	kryptertRune := mycrypt.Krypter([]rune("Kjevik;SN39040;18.03.2022 01:50;6"), mycrypt.AlfSem03, 4)
	kryptertString := string(kryptertRune)
	fmt.Println(kryptertString)

	alfLength := len(mycrypt.AlfSem03)

	dekryptertRune := mycrypt.Krypter(kryptertRune, mycrypt.AlfSem03, alfLength-4)
	dekryptertString := string(dekryptertRune)
	fmt.Println(dekryptertString)

}
