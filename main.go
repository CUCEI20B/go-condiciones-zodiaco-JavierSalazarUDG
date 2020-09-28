package main

import (
	"fmt"
)

func main() {
	var month int
	var day int
	fmt.Scan(&day)
	fmt.Scan(&month)

	switch month {
	case 1:
		if day >= 21 {
			fmt.Println("acuario")
		} else if day > 1 && day < 21 {
			fmt.Println("capricornio")
		}
	case 2:
		if day >= 19 {
			fmt.Println("piscis")
		} else if day > 1 && day < 19 {
			fmt.Println("acuario")
		}
	case 3:
		if day >= 21 {
			fmt.Println("aries")
		} else if day > 1 && day < 21 {
			fmt.Println("piscis")
		}
	case 4:
		if day >= 21 {
			fmt.Println("tauro")
		} else if day > 1 && day < 21 {
			fmt.Println("aries")
		}
	case 5:
		if day >= 21 {
			fmt.Println("geminis")
		} else if day > 1 && day < 21 {
			fmt.Println("tauro")
		}
	case 6:
		if day >= 22 {
			fmt.Println("cancer")
		} else if day > 1 && day < 22 {
			fmt.Println("geminis")
		}
	case 7:
		if day >= 23 {
			fmt.Println("leo")
		} else if day > 1 && day < 22 {
			fmt.Println("cancer")
		}
	case 8:
		if day >= 23 {
			fmt.Println("virgo")
		} else if day > 1 && day < 22 {
			fmt.Println("leo")
		}
	case 9:
		if day >= 23 {
			fmt.Println("libra")
		} else if day > 1 && day < 23 {
			fmt.Println("virgo")
		}
	case 10:
		if day > 23 {
			fmt.Println("escorpio")
		} else if day > 1 && day < 22 {
			fmt.Println("libra")
		}
	case 11:
		if day > 23 {
			fmt.Println("sagitario")
		} else if day > 1 && day < 22 {
			fmt.Println("escorpio")
		}
	case 12:
		if day > 23 {
			fmt.Println("capricornio")
		} else if day > 1 && day < 22 {
			fmt.Println("sagitario")
		}
	}

}
