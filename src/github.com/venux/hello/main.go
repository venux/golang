package main

import (
	"fmt"
	"github.com/venux/stringutil"
)

func main() {
	fmt.Printf("Hello,Go")

	fmt.Print(stringutil.Reverse("Hello,Nice to meet you,golang!!!"))
}