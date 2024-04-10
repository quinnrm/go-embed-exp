package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/quinnrm/go-embed-exp/internal/embedData"
)

func main() {
	fmt.Println("Go embed experiments.")

	efsData, err := embedData.EFS.ReadDir(".")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) Problem reading embeded filesystem root dir, %v", err)
	}

	fmt.Println("Listing the files in the embedded filesystem.")
	for i, d := range efsData {
		fmt.Printf("index; %d, is dir: %t, name: %s\n", i, d.IsDir(), d.Name())
	}

	fmt.Println()
	fmt.Println("Read \"hello.txt\" from the embeded FS.")

	txtHello, err := embedData.EFS.ReadFile("hello/hello.txt")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) Problem with \"hello.txt\", %v", err)
	}

	fmt.Println("Print \"hello.txt\"")
	fmt.Printf("%s", string(txtHello))
}
