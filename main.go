package main

import (
	"fmt"
	"io/ioutil"
	"os"

	sunit "github.com/coreos/fleet/Godeps/_workspace/src/github.com/coreos/go-systemd/unit"
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

	sourceOptions := source.Options
	destOptions := dest.Options

	length := len(sourceOptions)
	if length != len(destOptions) {
		fmt.Println("Different number of options in each unit")
		os.Exit(1)
	}

	valid := true
	var o *sunit.UnitOption

	for i := 0; i < length; i++ {
		if !sourceOptions[i].Match(destOptions[i]) {
			valid = false

			fmt.Printf("Found difference at option %d:\n", i+1)

			o = sourceOptions[i]
			fmt.Printf("<  [%s]\n", o.Section)
			fmt.Printf("<  %s=%s\n", o.Name, o.Value)

			o = destOptions[i]
			fmt.Printf(">  [%s]\n", o.Section)
			fmt.Printf(">  %s=%s\n", o.Name, o.Value)
		}
	}

	if !valid {
		os.Exit(1)
	}

	fmt.Println("Everything looks fine.")
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
