package day16

import (
	"log"
	"math"
	"strconv"
)

func Solve2() {
	bits := ReadLines()
	v, remaining := eval(bits)
	log.Println(v, remaining)
}

func eval(bits string) (int, string) {
	_, id := verAndType(bits)
	if id == 4 {
		v, b := evalType4(bits)
		return v, bits[b:]
	}

	subPackages, remaining := splitByOp(bits)

	switch id {
	case 0:
		total := 0
		log.Println("summing...")
		for subPackages != "" {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			subPackages = rem
			total += val
			log.Printf("adding %d, total is %d", val, total)
		}
		log.Println("done summing")
		return total, remaining
	case 1:
		total := 1
		log.Println("multiplying...")
		for subPackages != "" {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			total *= val
			log.Printf("multiplying %d, product is %d", val, total)
		}
		log.Println("done multiplying")
		return total, remaining
	case 2:
		log.Println("minning...")
		min := math.MaxInt32
		for subPackages != "" {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			if min > val {
				min = val
			}
		}
		log.Println("done minning")
		return min, remaining
	case 3:
		log.Println("maxing...")
		max := 0
		for subPackages != "" {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			if max < val {
				max = val
			}
		}
		log.Println("done maxing")
		return max, remaining
	case 5:
		log.Println("GT")
		l, r := 0, 0
		for i := 0; i < 2; i++ {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			if i == 0 {
				l = val
			} else {
				r = val
			}
		}
		log.Println("/GT")
		if l > r {
			return 1, remaining
		}
		return 0, remaining
	case 6:
		log.Println("LT")
		l, r := 0, 0
		for i := 0; i < 2; i++ {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			if i == 0 {
				l = val
			} else {
				r = val
			}
		}
		log.Println("/LT")
		if l < r {
			return 1, remaining
		}
		return 0, remaining
	case 7:
		log.Println("EQ")
		l, r := 0, 0
		for i := 0; i < 2; i++ {
			log.Println(subPackages[3:6])
			val, rem := eval(subPackages)
			log.Println(val)
			subPackages = rem
			if i == 0 {
				l = val
			} else {
				r = val
			}
		}
		log.Println("/EQ")
		if l == r {
			return 1, remaining
		}
		return 0, remaining
	}
	panic(id)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func splitByOp(bits string) (subpackages string, remaining string) {
	opType := bits[6]
	if opType == '0' {
		lenNum, err := strconv.ParseInt(bits[7:22], 2, 32)
		panicIf(err)
		subpackages = bits[22 : 22+lenNum]
		remaining = bits[22+lenNum:]
	} else {
		lenNum, err := strconv.ParseInt(bits[7:18], 2, 32)
		panicIf(err)
		b2 := bits[18:]
		for i := 0; i < int(lenNum); i++ {
			_, p, rem := GetPackageAndSubStr(b2)
			_, id := verAndType(p)
			subpackages += p
			if id != 4 {
				sp, rem2 := splitByOp(p + rem)
				subpackages += sp
				rem = rem2
			}
			b2 = rem
		}
		remaining = b2
	}
	return

}

func evalType4(bits string) (int, int) {
	breakPoint := 6
	var num string
	for breakPoint+4 < len(bits) {
		num += bits[breakPoint+1 : breakPoint+5]

		if bits[breakPoint] == '0' {
			break
		}
		breakPoint += 5
	}
	v, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(v), breakPoint + 5
}
