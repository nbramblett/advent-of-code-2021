package day16

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var hexToBin = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func Solve1() {
	bits := ReadLines()
	total := 0
	for !allZeros(bits) {
		v, _, rem := GetPackageAndSubStr(bits)
		bits = rem
		total += v
	}
	log.Println(total)
}

func allZeros(s string) bool {
	for _, r := range s {
		if r != '0' {
			return false
		}
	}
	return true
}

func verAndType(bits string) (int, int) {
	ver := bits[0:3]
	verNum, err := strconv.ParseInt(ver, 2, 8)
	if err != nil {
		panic(err)
	}
	typeId := bits[3:6]
	id, err := strconv.ParseInt(typeId, 2, 8)
	if err != nil {
		panic(err)
	}
	return int(verNum), int(id)
}

func GetPackageAndSubStr(bits string) (int, string, string) {
	verNum, id := verAndType(bits)
	if id == 4 {
		breakPoint := 6
		for breakPoint < len(bits) {
			if bits[breakPoint] == '0' {
				return int(verNum), bits[0:(breakPoint + 5)], bits[(breakPoint + 5):]
			}
			breakPoint += 5
		}
		panic("package seems to exceed total string")
	}

	opType := bits[6]
	if opType == '0' {
		return int(verNum), bits[0:22], bits[22:]
	} else {
		return int(verNum), bits[0:18], bits[18:]
	}
}

func ReadLines() string {
	file, err := os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	bits := ""
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			bits = bits + hexToBin[r]
		}
	}
	return bits
}
