package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"

	"github.com/quinnrm/go-embed-exp/internal/embedData"
)

func main() {
	fmt.Println("Go embed experiments.")

	fmt.Println()
	fmt.Println("Read \"hello.txt\" from the embeded FS.")

	txtHello, err := embedData.EFS.ReadFile("hello/hello.txt")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) Problem with \"hello.txt\", %v", err)
	}

	fmt.Printf("%s", string(txtHello))
	fmt.Println()

	fmt.Println("Walking the embedded filesystem.")
	efsFileList, err := walkEFS(embedData.EFS, ".")
	if !errors.Is(err, nil) {
		log.Fatalf("ERROR: (main) -> %v", err)
	}

	for _, p := range efsFileList {
		readEFSFile(embedData.EFS, p)
	}
}

// walkEFS steps through the embeded FS creating a list of files.
func walkEFS(efs embed.FS, path string) ([]string, error) {
	var files []string

	// An anonymous function is being used here because "files" can't be passed
	// into the function, and the function needs access to the current context.
	err := fs.WalkDir(
		efs,
		path,
		func(p string, d fs.DirEntry, err error) error {
			if !errors.Is(err, nil) {
				fmt.Println(err)
				return err
			}

			if !d.IsDir() {
				files = append(files, p)
			} else {
				fmt.Printf("%s\n", p)
			}

			return nil
		},
	)
	// err should always be nil since it's an embedded FS, but "should always".
	if !errors.Is(err, nil) {
		return files, fmt.Errorf(
			"(main.walkEFS) Problem walking the embedded filesystem, %w",
			err,
		)
	}

	return files, nil
}

// readEFSFile reads and prints the contents of a file from an embedded FS.
func readEFSFile(efs embed.FS, path string) error {
	data, err := efs.ReadFile(path)
	// err should always be nil since it's an embedded FS, but just in case.
	if !errors.Is(err, nil) {
		return fmt.Errorf(
			"(main.readEFSFile) Problem reading embeded filesystem root dir, %w",
			err,
		)
	}

	fmt.Println(path)
	fmt.Printf("%s\n", string(data))

	return nil
}
