package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	message.Set(language.English, "%d items to do",
		plural.Selectf(1, "%d", "=0", "no items to do",
			plural.One, "one item to do", "<100", "%[1]d items to do",
			plural.Other, "lot of item to do"))

	message.Set(language.English, "The average is %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "The average is zero",
			"=1", "The average is one",
			plural.Other, "The average is %[1]f",
		))

	prt := message.NewPrinter(language.English)
	prt.Printf("%d items to do", 0)
	fmt.Println()
	prt.Printf("%d items to do", 1)
	fmt.Println()
	prt.Printf("%d items to do", 10)
	fmt.Println()
	prt.Printf("%d items to do", 1000)
	fmt.Println()

	prt.Printf("The average is %.2f", 0.8)
	fmt.Println()
	prt.Printf("The average is %.2f", 1.0)
	fmt.Println()
	prt.Printf("The average is %.2f", 10.0)
	fmt.Println()
}
