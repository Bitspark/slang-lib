package main

import (
	"slang"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func runTests(dir string) (int, int, int) {
	testFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	succsTotal := 0
	failsTotal := 0
	filesTotal := 0
	for _, file := range testFiles {
		if file.IsDir() {
			succs, fails, files := runTests(dir + file.Name() + "/")
			succsTotal += succs
			failsTotal += fails
			filesTotal += files
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
			succsTotal += succs
			failsTotal += fails
		} else {
			succsTotal += succs
			failsTotal += fails
		}
		filesTotal++
		fmt.Println()
	}

	return succsTotal, failsTotal, filesTotal
}

func main() {
	succsTotal, failsTotal, filesTotal := runTests("./slang/")

	fmt.Println()
	fmt.Println()
	fmt.Println("SUMMARY OVER ALL TEST FILES")
	fmt.Println()
	fmt.Printf("Files: %3d\n", filesTotal)
	fmt.Printf("Tests: %3d / %3d\n", succsTotal, succsTotal + failsTotal)
	fmt.Println()

	if failsTotal == 0 {
		fmt.Println("ALL TESTS PASSED :)")
		os.Exit(0)
	} else {
		fmt.Printf("%d TEST(S) FAILED :(\n", failsTotal)
		os.Exit(-1)
	}
}
