package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Meter float64
type Feet float64

const FeetsInMeter = 3.28

type Kilogram float64
type Pound float64

const PoundsInKilogram = 0.45

func MToF(m Meter) Feet {
	return Feet(m * FeetsInMeter)
}

func FToM(f Feet) Meter {
	return Meter(f / FeetsInMeter)
}

func KToP(k Kilogram) Pound {
	return Pound(k * PoundsInKilogram)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p / PoundsInKilogram)
}

func convert(s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%gm = %gft, %gft = %gm\n", v, MToF(Meter(v)), v, FToM(Feet(v)))
	fmt.Printf("%gkg = %glb, %glb = %gkg\n", v, KToP(Kilogram(v)), v, PToK(Pound(v)))
}

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			convert(s.Text())
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
