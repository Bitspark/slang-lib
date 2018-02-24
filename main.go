package main

import (
	"slang"
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

	succsTotal := 0
	failsTotal := 0
	errors := false
	files := 0
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
				errors = true
			}
			succsTotal += succs
			failsTotal += fails
		} else {
			succsTotal += succs
			failsTotal += fails
		}
		files++
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("SUMMARY OVER ALL TEST FILES")
	fmt.Println()
	fmt.Printf("Files: %3d\n", files)
	fmt.Printf("Tests: %3d / %3d\n", succsTotal, succsTotal + failsTotal)
	fmt.Println()

	if failsTotal == 0 && !errors {
		fmt.Println("ALL TESTS PASSED :)")
		os.Exit(0)
	} else if failsTotal != 0 {
		fmt.Printf("%d TEST(S) FAILED :(\n", failsTotal)
		os.Exit(-1)
	} else {
		fmt.Printf("ERRORS OCCURRED :(\n")
		os.Exit(-1)
	}
}

func main() {
	runTests("./slang/")
}
