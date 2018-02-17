package main

import (
	"github.com/Bitspark/slang"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func runTests(dir string) {
	testFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	tests := 0
	for _, file := range testFiles {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), "_test.yaml") {
			continue
		}
		path := dir + file.Name()
		fmt.Println()
		fmt.Printf("RUNNING TEST FILE %s\n", path)
		fmt.Println()
		if succs, fails, err := slang.TestOperator(path, os.Stdout, false); err != nil || succs == 0 || fails != 0 {
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(-1)
		} else {
			tests += succs
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("SUMMARY OVER ALL TEST FILES")
	fmt.Println()
	fmt.Printf("Files: %3d\n", len(testFiles))
	fmt.Printf("Tests: %3d\n", tests)
	fmt.Println()
}

func main() {
	runTests("./slib/")
}
