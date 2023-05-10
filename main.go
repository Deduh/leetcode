package main

import (
	"strings"
)

func main() {
	s := "MCMXCIV"
	romanToInt(s)
}

func romanToInt(s string) int {
	old := strings.Split(s, "")
	new := make([]string, 0, len(old))
	sum := 0

	toRomans(&old, &new)

	for i := 0; i < len(new); i++ {
		sum += int(toNumber(new[i]))
	}
	return sum
}

func toRomans(old *[]string, new *[]string) {
	for i := 0; i < len(*old); i++ {
		if i+1 < len(*old) && (*old)[i] == "I" && (*old)[i+1] == "V" {
			*new = append(*new, "IV")
			i++
		} else if i+1 < len(*old) && (*old)[i] == "I" && (*old)[i+1] == "X" {
			*new = append(*new, "IX")
			i++
		} else if i+1 < len(*old) && (*old)[i] == "X" && (*old)[i+1] == "L" {
			*new = append(*new, "XL")
			i++
		} else if i+1 < len(*old) && (*old)[i] == "X" && (*old)[i+1] == "C" {
			*new = append(*new, "XC")
			i++
		} else if i+1 < len(*old) && (*old)[i] == "C" && (*old)[i+1] == "D" {
			*new = append(*new, "CD")
			i++
		} else if i+1 < len(*old) && (*old)[i] == "C" && (*old)[i+1] == "M" {
			*new = append(*new, "CM")
			i++
		} else {
			*new = append(*new, (*old)[i])
		}
	}
}

func toNumber(s string) int {
	switch s {
	case "I":
		return 1
	case "IV":
		return 4
	case "V":
		return 5
	case "IX":
		return 9
	case "X":
		return 10
	case "XL":
		return 40
	case "L":
		return 50
	case "XC":
		return 90
	case "C":
		return 100
	case "CD":
		return 400
	case "D":
		return 500
	case "CM":
		return 900
	case "M":
		return 1000
	default:
		return 0
	}
}
