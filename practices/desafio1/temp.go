package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273.0
	
	fmt.Printf("A temperatura de ebulicao da água em K = %g e a temperatura de ebulicao da água em C = %g.", tempK, tempC)
}