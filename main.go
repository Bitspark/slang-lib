package main

import (
	"fmt"
	"github.com/Bitspark/slang/pkg/api"
	"github.com/Bitspark/slang/pkg/storage"
	"log"
	"os"
)

func runTests(dir string) (int, int, int) {
	stor := storage.NewStorage(nil).AddLoader(storage.NewFileSystem(dir))
	tb := api.NewTestBench(stor)

	succsTotal := 0
	failsTotal := 0
	filesTotal := 0

	opIds, err := stor.List()

	if err != nil {
		log.Fatal(err)
	}

	for _, opId := range opIds {
		opDef, err := stor.Load(opId)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("BLUEPRINT", opId.String(), opDef.Meta.Name)
		if opDef.Meta.ShortDescription != "" {
			fmt.Println("         ", opDef.Meta.ShortDescription)
		}
		fmt.Println()

		if succs, fails, err := tb.Run(opId, os.Stdout, false); err != nil || succs == 0 || fails != 0 {
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

	fmt.Println("SUMMARY OVER ALL TEST FILES")
	fmt.Println()
	fmt.Printf("Files: %3d\n", filesTotal)
	fmt.Printf("Tests: %3d / %3d\n", succsTotal, succsTotal+failsTotal)
	fmt.Println()

	if failsTotal == 0 {
		fmt.Println("ALL TESTS PASSED :)")
		os.Exit(0)
	} else {
		fmt.Printf("%d TEST(S) FAILED :(\n", failsTotal)
		os.Exit(-1)
	}
}
