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

func zon1(troskel int, vila int) int {
	hog := ((troskel-vila) / 2) + vila
	return hog
}

//Zon2 (tröskel-högsta zon1) / 2 + högsta zon1 = högsta zon2
func zon2(troskel int, vila int) (int, int) {
	hog1 := zon1(troskel, vila)
	hog2 := (troskel-hog1) / 2 + hog1
	return hog1, hog2
}

// högsta zon2 och lägsta zon4
func zon3(troskel int, vila int, hogh int) (int, int){
	_, hogh2 := zon2(hogh,vila)
	luw4, _ := zon4(troskel, vila, hogh)
	return hogh2, luw4
}

//Zon4 tröskelpuls +- (pulsspann * 0.03)
func zon4(troskel int, vila int, hogh int) (int, int) {
	spann := pulsspann(hogh, vila)
	luw4 := troskel - int(float64(spann) * 0.03)
	hogh4 := troskel + int(float64(spann) * 0.03)
	return luw4, hogh4
}

//Zon5 högsta zon4 och luwsta zon6
func zon5(troskel int, vila int, hogh int) (int, int){
	_, hogh4 := zon4(troskel,vila,hogh)
	luw6, _ := zon6(troskel, vila, hogh)
	return hogh4, luw6
}

//Zon6 hoghpuls +- (pulsspann * 0.06)
func zon6(troskel int,vila int, hogh int) (int, int) {
	spann := pulsspann(hogh, vila)
	luw6 := hogh - int(float64(spann) * 0.06)
	hogh6 := hogh + int(float64(spann) * 0.06)
	return luw6, hogh6
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
	z1 := zon1(troskel,vila)
	fmt.Printf("Zon1 slutar vid %v\n", z1)
	z21, z22 := zon2(troskel, vila)
	fmt.Printf("Zon2 %v - %v\n", z21, z22)
	z31, z32 := zon3(troskel, vila, hogh)
	fmt.Printf("Zon3 %v - %v\n", z31, z32)
	z41, z42 := zon4(troskel,vila,hogh)
	fmt.Printf("Zon4 %v - %v\n", z41, z42)
	z51, z52 := zon5(troskel,vila,hogh)
	fmt.Printf("Zon5 %v - %v\n", z51, z52)
	z61, z62 := zon6(troskel,vila,hogh)
	fmt.Printf("Zon6 %v - %v\n", z61, z62)
}
