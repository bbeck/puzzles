package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bitfield/script"
	"github.com/magefile/mage/mg"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

//goland:noinspection GoUnusedExportedType
type AdventOfCode mg.Namespace

var Year int
var Day int
var Part int
var Session []byte

func (aoc AdventOfCode) Run() error {
	mg.Deps(aoc.parse, aoc.download)

	output, err := aoc.run()
	fmt.Print(output)
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

	filename := fmt.Sprintf("%s/main.go", dir)
	_, err := script.File(source).WriteFile(filename)
	if err != nil {
		return err
	}

	// Open the new file in the editor
	editor := `"/Applications/IntelliJ IDEA.app/Contents/MacOS/idea"`
	return script.Exec(fmt.Sprintf("%s %s", editor, filename)).Wait()
}

func (aoc AdventOfCode) Verify() error {
	mg.Deps(aoc.parse)

	expected, err := script.File("cmd/advent-of-code/.solutions").
		FilterScan(func(line string, w io.Writer) {
			var buf bytes.Buffer

			// Parse the year/day/part prefix on each line
			fields := strings.Split(line, " ")
			year := lib.ParseInt(fields[0])
			day := lib.ParseInt(fields[1])
			part := lib.ParseInt(fields[2])

			if year == Year && day == Day && part == Part {
				_, _ = buf.WriteString(strings.Join(fields[3:], " "))
				buf.WriteRune('\n')
			}

			_, _ = w.Write(buf.Bytes())
		}).
		String()
	if err != nil {
		return err
	}

	actual, err := aoc.run()
	if err != nil {
		return err
	}

	// The output for some problems is multiple lines and sometimes those lines
	// have leading or trailing spaces.  The solution file doesn't always capture
	// trailing spaces properly, so let's convert the expected and actual strings
	// into slices of stripped lines for comparison.
	convert := func(s string) []string {
		var lines []string
		for _, line := range strings.Split(s, "\n") {
			trimmed := strings.Trim(line, " \n")
			if trimmed != "" {
				lines = append(lines, trimmed)
			}
		}
		return lines
	}

	eLines := convert(expected)
	aLines := convert(actual)

	if reflect.DeepEqual(aLines, eLines) {
		fmt.Printf("✅ YEAR=%d DAY=%02d PART=%d %s\n", Year, Day, Part, aLines[0])
	} else {
		fmt.Printf("❌ YEAR=%d DAY=%02d PART=%d\n", Year, Day, Part)
		fmt.Println(actual)
		fmt.Println()
		fmt.Println(expected)
	}

	return nil
}

func (aoc AdventOfCode) ListYear() error {
	mg.Deps(aoc.parse)

	for day := 1; day <= 25; day++ {
		for part := 1; part <= 2; part++ {
			if day == 25 && part == 2 {
				continue
			}
			fmt.Printf("%d %d %d\n", Year, day, part)
		}
	}

	return nil
}

func (aoc AdventOfCode) parse() error {
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

	request.Header.Set("User-Agent", "automation by bmbeck@gmail.com")
	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(Session),
	})

	_, err = script.Do(request).WriteFile(filename)
	return err
}

func (aoc AdventOfCode) run() (string, error) {
	dir := fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", Year, Day, Part)
	err := script.IfExists(dir).Error()
	if err != nil {
		return "", fmt.Errorf("%s does not exist", dir)
	}

	// Change to the new directory, but be sure to return to the current
	// directory on exit in case we are running multiple programs.
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}
	defer func() { _ = os.Chdir(pwd) }()

	// Run the script
	return script.Exec("go run .").String()
}
