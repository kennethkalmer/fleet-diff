package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/fleet/unit"
)

func main() {
	count := len(os.Args)

	if count != 3 {
		fmt.Println("No unit files provided!")
		os.Exit(1)
	}

	first := os.Args[1]
	second := os.Args[2]

	var source *unit.UnitFile
	var dest *unit.UnitFile
	var err error

	if first == "-" {
		source, err = getUnitFromStdin()
	} else {
		source, err = getUnitFromFile(first)
	}

	if err != nil {
		panic(err)
	}

	if second == "-" {
		dest, err = getUnitFromStdin()
	} else {
		dest, err = getUnitFromFile(second)
	}

	if err != nil {
		panic(err)
	}

	sourceHash := source.Hash()
	destHash := dest.Hash()

	if sourceHash != destHash {
		fmt.Println("Units are different!")
		os.Exit(1)
	} else {
		fmt.Println("Everything looks fine.")
		os.Exit(0)
	}
}

// getUnitFromStdin attempts to load a Unit from stdin
func getUnitFromStdin() (*unit.UnitFile, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return unit.NewUnitFile(string(bytes))
}

// getUnitFromFile attempts to load a Unit from a given filename
// It returns the Unit or nil, and any error encountered
func getUnitFromFile(file string) (*unit.UnitFile, error) {
	out, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return unit.NewUnitFile(string(out))
}
