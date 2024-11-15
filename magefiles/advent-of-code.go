package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/magefile/mage/mg"
	"net/http"
	"os"
	"time"
)

type AdventOfCode mg.Namespace

var Year int
var Day int
var Part int
var Session []byte

func (aoc AdventOfCode) Run() error {
	mg.Deps(aoc.parse, aoc.download)

	dir := fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", Year, Day, Part)
	err := script.IfExists(dir).Error()
	if err != nil {
		return fmt.Errorf("%s does not exist", dir)
	}

	// Change to the new directory, but be sure to return to the current
	// directory on return in case we are running multiple programs.
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	defer func() { _ = os.Chdir(pwd) }()

	// Run the script
	_, err = script.Exec("go run .").Stdout()
	return err
}

func (aoc AdventOfCode) Watch() error {
	mg.Deps(aoc.parse)

	// Always watch the shared library directory, the directory of the part that's
	// being solved, and the input file.
	_, err := script.
		Echo(
			Files(
				"puz",
				fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", Year, Day, Part),
				fmt.Sprintf("cmd/advent-of-code/%d/%02d-1/input.txt", Year, Day),
			),
		).
		WithEnv(append(os.Environ(), []string{
			fmt.Sprintf("YEAR=%d", Year),
			fmt.Sprintf("DAY=%d", Day),
			fmt.Sprintf("PART=%d", Part),
		}...)).
		Exec("entr -c sh -c 'go run mage.go AdventOfCode:run'").
		Stdout()

	return err
}

func (aoc AdventOfCode) Next() error {
	mg.Deps(aoc.parse)

	if Day == 25 && Part == 1 {
		Year++
		Day = 1
		Part = 1
	} else if Part == 1 {
		Part++
	} else {
		Day++
		Part = 1
	}

	// The new directory we're going to work in.
	dir := fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", Year, Day, Part)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// The source of what we're going to copy as our main.go.
	var source string
	if Part == 1 {
		source = fmt.Sprintf("cmd/advent-of-code/.template")
	} else {
		source = fmt.Sprintf("cmd/advent-of-code/%d/%02d-1/main.go", Year, Day)
	}

	_, err := script.File(source).WriteFile(fmt.Sprintf("%s/main.go", dir))
	return err
}

func (AdventOfCode) parse() error {
	var err error

	if Year, err = LookupInt("YEAR"); err != nil {
		// The year wasn't in the environment, infer it from the filesystem.
		for Year = time.Now().Year(); Year > 0; Year-- {
			dir := fmt.Sprintf("cmd/advent-of-code/%d", Year)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if Year == 0 {
			return fmt.Errorf("unable to infer year")
		}
	}

	if Day, err = LookupInt("DAY"); err != nil {
		// The day wasn't in the environment, infer it from the filesystem.
		for Day = 25; Day > 0; Day-- {
			dir := fmt.Sprintf("cmd/advent-of-code/%d/%02d-1", Year, Day)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if Day == 0 {
			return fmt.Errorf("unable to infer day")
		}
	}

	if Part, err = LookupInt("PART"); err != nil {
		// The part wasn't in the environment, infer it from the filesystem.
		for Part = 2; Part > 0; Part-- {
			dir := fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", Year, Day, Part)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if Part == 0 {
			return fmt.Errorf("unable to infer part")
		}
	}

	Session, err = script.File("cmd/advent-of-code/.session").Bytes()
	if err != nil {
		return fmt.Errorf("unable to read session file: %w", err)
	}

	return nil
}

func (aoc AdventOfCode) download() error {
	mg.Deps(aoc.parse)

	filename := fmt.Sprintf("cmd/advent-of-code/%d/%02d-1/input.txt", Year, Day)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", Year, Day)

	// First check if the file is already present.
	if script.IfExists(filename).Error() == nil {
		return nil
	}

	// The file wasn't present, download it.
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(Session),
	})

	_, err = script.Do(request).WriteFile(filename)
	return err
}
