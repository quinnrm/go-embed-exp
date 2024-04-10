package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
)

//go:embed hello/hello.txt
//go:embed data
var eFS embed.FS

func main() {
	fmt.Println("Go embed experiments.")

	efsData, err := eFS.ReadDir(".")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) Problem reading embeded filesystem root dir, %v", err)
	}

	fmt.Println("Listing the files in the embedded filesystem.")
	for i, d := range efsData {
		fmt.Printf("index; %d, is dir: %t, name: %s\n", i, d.IsDir(), d.Name())
	}

	fmt.Println()
	fmt.Println("Read \"hello.txt\" from the embeded FS.")

	txtHello, err := eFS.ReadFile("hello/hello.txt")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) Problem with \"hello.txt\", %v", err)
	}

	fmt.Println("Print \"hello.txt\"")
	fmt.Println(string(txtHello))
}
