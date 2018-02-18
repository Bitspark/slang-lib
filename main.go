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
				log.Println(err)
				failsTotal++
			}
		} else {
			succsTotal += succs
			failsTotal += fails
			files++
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("SUMMARY OVER ALL TEST FILES")
	fmt.Println()
	fmt.Printf("Files: %3d\n", files)
	fmt.Printf("Tests: %3d / %3d\n", succsTotal, succsTotal + failsTotal)
	fmt.Println()

	if failsTotal == 0 {
		fmt.Println("ALL TESTS PASSED :)")
	} else {
		fmt.Printf("%d TEST(S) FAILED :(\n", failsTotal)
	}
}

func main() {
	runTests("./slib/")
}
