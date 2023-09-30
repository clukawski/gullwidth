package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clukawski/gullwidth"
)

func main() {
	input := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	fullwidth, err := gullwidth.Fullwidth(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullwidth)
}
