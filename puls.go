package main

import (
	"fmt"
)
// hoghpuls-vilopuls
func pulsspann(hogh int, vila int) int {
	return hogh - vila
}

//Zon1 (tröskel-vilopuls) / 2 + vilopuls = högsta puls i zon1
//Aktiv vila

func zon0(troskel int, vila int) int {
	hog := ((troskel-vila) / 2) + vila
	return hog
}

//Zon2 (tröskel-högsta zon1) / 2 + högsta zon1 = högsta zon2
func zon1(troskel int, vila int) (int, int) {
	hog0 := zon0(troskel, vila)
	hog1 := (troskel-hog0) / 2 + hog0
	return hog0, hog1
}

// högsta zon2 och lägsta zon4
func zon2(troskel int, vila int, hogh int) (int, int){
	_, hogh1 := zon1(troskel,vila)
	luw3, _ := zon3(troskel, vila, hogh)
	return hogh1, luw3
}

//Zon4 tröskelpuls +- (pulsspann * 0.03)
func zon3(troskel int, vila int, hogh int) (int, int) {
	spann := pulsspann(hogh, vila)
	luw3 := troskel - int(float64(spann) * 0.03)
	hogh3 := troskel + int(float64(spann) * 0.03)
	return luw3, hogh3
}

//Zon5 högsta zon4 och luwsta zon6
func zon4(troskel int, vila int, hogh int) (int, int){
	_, hogh3 := zon3(troskel,vila,hogh)
	luw5, _ := zon5(troskel, vila, hogh)
	return hogh3, luw5
}

//Zon6 hoghpuls +- (pulsspann * 0.06)
func zon5(troskel int,vila int, hogh int) (int, int) {
	spann := pulsspann(hogh, vila)
	luw5 := hogh - int(float64(spann) * 0.06)
	hogh5 := hogh + int(float64(spann) * 0.06)
	return luw5, hogh5
}

func main() {
	fmt.Println("Vad är din maxpuls?")
	var hogh int
	fmt.Scan(&hogh)
	fmt.Println("Hur hög är din mjölksyratröskel?")
	var troskel int
	fmt.Scan(&troskel)
	fmt.Println("Hur låg är vilopulsen")
	var vila int
	fmt.Scan(&vila)
	z1 := zon0(troskel,vila)
	fmt.Printf("Zon0 slutar vid %v\n", z1)
	z21, z22 := zon1(troskel, vila)
	fmt.Printf("Zon1 %v - %v\n", z21, z22)
	z31, z32 := zon2(troskel, vila, hogh)
	fmt.Printf("Zon2 %v - %v\n", z31, z32)
	z41, z42 := zon3(troskel,vila,hogh)
	fmt.Printf("Zon3 %v - %v\n", z41, z42)
	z51, z52 := zon4(troskel,vila,hogh)
	fmt.Printf("Zon4 %v - %v\n", z51, z52)
	z61, z62 := zon5(troskel,vila,hogh)
	fmt.Printf("Zon5 %v - %v\n", z61, z62)
}
