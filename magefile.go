//go:build mage

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bbeck/puzzles/lib"
	"github.com/bitfield/script"
	"github.com/magefile/mage/mg"
)

type Problem struct {
	Year int
	Day  int
	Part int
}

type TimeOfDay struct {
	Hour, Minute, Second int
}

type Site struct {
	ID        string
	Session   []byte
	Directory string
	NumDays   int
	NumParts  int
	StartTime TimeOfDay
}

var Sites = map[string]Site{
	"advent-of-code": {
		ID:        "advent-of-code",
		Directory: "cmd/advent-of-code",
		NumDays:   12,
		NumParts:  2,
		StartTime: TimeOfDay{Hour: 23, Minute: 0, Second: 0},
	},
	"everybody-codes": {
		ID:        "everybody-codes",
		Directory: "cmd/everybody-codes",
		NumDays:   20,
		NumParts:  3,
		StartTime: TimeOfDay{Hour: 17, Minute: 0, Second: 0},
	},
}

var SiteAliases = map[string]string{
	"adventofcode":    "advent-of-code",
	"advent-of-code":  "advent-of-code",
	"aoc":             "advent-of-code",
	"everybodycodes":  "everybody-codes",
	"everybody-codes": "everybody-codes",
	"ec":              "everybody-codes",
}

var problem Problem
var site Site

//
// Targets
//

// Run will run the program for the site, year, day and part that is being
// worked on.  The output of the program will be printed to standard output.
//
//goland:noinspection GoUnusedExportedFunction
func Run() error {
	mg.Deps(ParseEnv, DownloadInput)

	output, duration, err := RunHelper()
	if err == nil {
		// Show the output plus the duration.  Put the duration on a line of its
		// own if the output has multiple lines in it.
		output = strings.TrimSuffix(output, "\n")
		if strings.Contains(output, "\n") {
			fmt.Println(output)
			fmt.Printf("[%dms]\n", duration.Milliseconds())
		} else {
			fmt.Print(output)
			fmt.Printf(" [%dms]\n", duration.Milliseconds())
		}
	} else {
		fmt.Print(output)
	}
	return err
}

// Watch will run the program being worked on whenever a source file changes.
//
//goland:noinspection GoUnusedExportedFunction
func Watch() error {
	mg.Deps(ParseEnv)

	// Always watch the shared library directory, the directory of the part that's
	// being solved, and the input file.
	var files string
	switch site.ID {
	case "advent-of-code":
		// For Advent of Code each day shares a single input file that's located in
		// the part 1 directory.
		files = Files(
			"lib",
			fmt.Sprintf("%s/%d/%02d-%d", site.Directory, problem.Year, problem.Day, problem.Part),
			fmt.Sprintf("%s/%d/%02d-1/input.txt", site.Directory, problem.Year, problem.Day),
		)

	case "everybody-codes":
		// For Everybody Codes each part has a separate input file that's located in
		// the part's directory.
		files = Files(
			"lib",
			fmt.Sprintf("%s/%d/%02d-%d", site.Directory, problem.Year, problem.Day, problem.Part),
		)
	}

	_, err := script.
		Echo(files).
		WithEnv(append(os.Environ(), []string{
			fmt.Sprintf("SITE=%s", site.ID),
			fmt.Sprintf("YEAR=%d", problem.Year),
			fmt.Sprintf("DAY=%d", problem.Day),
			fmt.Sprintf("PART=%d", problem.Part),
		}...)).
		Exec("entr -c sh -c 'go run mage.go run'").
		Stdout()

	return err
}

// Next will create and populate the working directory for the next program to
// work on.
//
//goland:noinspection GoUnusedExportedFunction
func Next() error {
	mg.Deps(ParseEnv)

	if site.ID == "advent-of-code" && problem.Day == site.NumDays && problem.Part == 1 {
		// Special case for Advent of Code where there is no 2nd part on Christmas.
		problem.Year++
		problem.Day = 1
		problem.Part = 1
	} else if problem.Day == site.NumDays && problem.Part == site.NumParts {
		problem.Year++
		problem.Day = 1
		problem.Part = 1
	} else if problem.Part == site.NumParts {
		problem.Day++
		problem.Part = 1
	} else {
		problem.Part++
	}

	// The new directory we're going to work in.
	dir := fmt.Sprintf("%s/%d/%02d-%d", site.Directory, problem.Year, problem.Day, problem.Part)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// The source of what we're going to copy as our main.go.
	var source string
	if problem.Part == 1 {
		source = fmt.Sprintf("%s/.template", site.Directory)
	} else {
		source = fmt.Sprintf("%s/%d/%02d-%d/main.go", site.Directory, problem.Year, problem.Day, problem.Part-1)
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

// Verify will run the program being worked on and report whether the output
// matches the expected solution.
//
//goland:noinspection GoUnusedExportedFunction
func Verify() error {
	mg.Deps(ParseEnv, DownloadInput)

	path := fmt.Sprintf("%s/.solutions", site.Directory)
	expected, err := script.File(path).
		FilterScan(func(line string, w io.Writer) {
			var buf bytes.Buffer

			// Parse the year/day/part prefix on each line
			fields := strings.Split(line, " ")
			year := lib.ParseInt(fields[0])
			day := lib.ParseInt(fields[1])
			part := lib.ParseInt(fields[2])

			if year == problem.Year && day == problem.Day && part == problem.Part {
				_, _ = buf.WriteString(strings.Join(fields[3:], " "))
				buf.WriteRune('\n')
			}

			_, _ = w.Write(buf.Bytes())
		}).
		String()
	if err != nil {
		return err
	}

	actual, duration, err := RunHelper()
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
		fmt.Printf("✅ SITE=%s YEAR=%d DAY=%02d PART=%d %s [%dms]\n", site.ID, problem.Year, problem.Day, problem.Part, aLines[0], duration.Milliseconds())
	} else {
		fmt.Printf("❌ SITE=%s YEAR=%d DAY=%02d PART=%d [%dms]\n", site.ID, problem.Year, problem.Day, problem.Part, duration.Milliseconds())
		fmt.Println("EXPECT:", strings.TrimRight(expected, "\n"))
		fmt.Println("ACTUAL:", strings.TrimRight(actual, "\n"))
		fmt.Println()
	}

	return nil
}

// WaitUntilStartTime will block until it is the start time for puzzles to be
// released today.  If it is already after the start time then this method will
// return immediately.
//
//goland:noinspection GoUnusedExportedFunction
func WaitUntilStartTime() error {
	mg.Deps(ParseEnv)

	startTOD := Sites[site.ID].StartTime

	for {
		now := time.Now()
		start := time.Date(
			now.Year(), now.Month(), now.Day(),
			startTOD.Hour, startTOD.Minute, startTOD.Second, 0,
			now.Location(),
		)

		if now.After(start) {
			break
		}

		time.Sleep(max(start.Sub(now)/2, time.Second))
	}

	return nil
}

// ListDay enumerates all parts that exist for a day.
//
//goland:noinspection GoUnusedExportedFunction
func ListDay() {
	mg.Deps(ParseEnv)

	for part := 1; part <= site.NumParts; part++ {
		// Check if a main.go file exists
		path := fmt.Sprintf("%s/%d/%02d-%d/main.go", site.Directory, problem.Year, problem.Day, part)
		if script.IfExists(path).Error() == nil {
			fmt.Printf("%d %d %d\n", problem.Year, problem.Day, part)
		}
	}
}

// ListYear enumerates all days and parts that are expected for the year.
//
//goland:noinspection GoUnusedExportedFunction
func ListYear() {
	mg.Deps(ParseEnv)

	for day := 1; day <= site.NumDays; day++ {
		for part := 1; part <= site.NumParts; part++ {
			// Check if a main.go file exists
			path := fmt.Sprintf("%s/%d/%02d-%d/main.go", site.Directory, problem.Year, day, part)
			if script.IfExists(path).Error() == nil {
				fmt.Printf("%d %d %d\n", problem.Year, day, part)
			}
		}
	}
}

// ParseEnv will read environment variables to determine which site, year, day
// and part is being worked on.  If a variable is not present in the environment
// then an attempt will be made to infer the most recent problem is being worked
// on.
func ParseEnv() {
	name, err := Lookup("SITE")
	if err != nil {
		panic("unable to infer site")
	}
	site = Sites[SiteAliases[strings.ToLower(name)]]

	year, err := LookupInt("YEAR")
	if err != nil {
		// The year wasn't in the environment, infer it from the filesystem.
		for year = time.Now().Year(); year > 0; year-- {
			dir := fmt.Sprintf("%s/%d", site.Directory, year)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if year == 0 {
			panic("unable to infer year")
		}
	}
	problem.Year = year

	day, err := LookupInt("DAY")
	if err != nil {
		// The day wasn't in the environment, infer it from the filesystem.
		for day = site.NumDays; day > 0; day-- {
			dir := fmt.Sprintf("%s/%d/%02d-1", site.Directory, year, day)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if day == 0 {
			panic("unable to infer day")
		}
	}
	problem.Day = day

	part, err := LookupInt("PART")
	if err != nil {
		// The part wasn't in the environment, infer it from the filesystem.
		for part = site.NumParts; part > 0; part-- {
			dir := fmt.Sprintf("%s/%d/%02d-%d", site.Directory, year, day, part)
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if part == 0 {
			panic("unable to infer part")
		}
	}
	problem.Part = part

	path := fmt.Sprintf("%s/.session", site.Directory)
	session, err := script.File(path).Bytes()
	if err != nil {
		panic(fmt.Errorf("unable to read session file: %w", err))
	}
	site.Session = session
}

// DownloadInput will ensure that the input file for the year, day and part
// being worked on is present in the filesystem.  If the input file is not
// present then an attempt will be made to download it from the site.
func DownloadInput() error {
	mg.Deps(ParseEnv)

	switch site.ID {
	case "advent-of-code":
		return DownloadAdventOfCodeInput()
	case "everybody-codes":
		return DownloadEverybodyCodesInput()
	default:
		return nil
	}
}

func DownloadAdventOfCodeInput() error {
	filename := fmt.Sprintf("%s/%d/%02d-1/input.txt", site.Directory, problem.Year, problem.Day)

	// First check if the file is already present.
	if script.IfExists(filename).Error() == nil {
		return nil
	}

	// The file wasn't present, download it.
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", problem.Year, problem.Day)
	bs, err := Fetch(url)
	if err != nil {
		return err
	}

	// Save the input
	_, err = script.Echo(string(bs)).WriteFile(filename)
	return err
}

func DownloadEverybodyCodesInput() error {
	filename := fmt.Sprintf("%s/%d/%02d-%d/input.txt", site.Directory, problem.Year, problem.Day, problem.Part)

	// First check if the file is already present.
	if script.IfExists(filename).Error() == nil {
		return nil
	}

	// Load the seed
	type SeedResponse struct {
		Seed int `json:"seed"`
	}

	url := "https://everybody.codes/api/user/me"
	sr, err := FetchJSON[SeedResponse](url)
	if err != nil {
		return err
	}

	// Fetch input notes
	type InputNotesResponse struct {
		Part1 string `json:"1"`
		Part2 string `json:"2"`
		Part3 string `json:"3"`
	}

	url = fmt.Sprintf("https://everybody-codes.b-cdn.net/assets/%d/%d/input/%d.json", problem.Year, problem.Day, sr.Seed)
	inr, err := FetchJSON[InputNotesResponse](url)
	if err != nil {
		return err
	}

	// Fetch AES keys
	type AESKeysResponse struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
		Key3 string `json:"key3"`
	}

	url = fmt.Sprintf("https://everybody.codes/api/event/%d/quest/%d", problem.Year, problem.Day)
	kr, err := FetchJSON[AESKeysResponse](url)
	if err != nil {
		return err
	}

	// Decrypt the input
	var input string
	switch problem.Part {
	case 1:
		input, err = DecryptAES(inr.Part1, kr.Key1)
	case 2:
		input, err = DecryptAES(inr.Part2, kr.Key2)
	case 3:
		input, err = DecryptAES(inr.Part3, kr.Key3)
	}
	if err != nil {
		return err
	}

	// Save the input
	_, err = script.Echo(input).WriteFile(filename)
	return err
}

//
// Helpers
//

func RunHelper() (string, time.Duration, error) {
	dir := fmt.Sprintf("%s/%d/%02d-%d", site.Directory, problem.Year, problem.Day, problem.Part)
	err := script.IfExists(dir).Error()
	if err != nil {
		return "", 0, fmt.Errorf("%s does not exist", dir)
	}

	// Change to the new directory, but be sure to return to the current
	// directory on exit in case we are running multiple programs.
	pwd, err := os.Getwd()
	if err != nil {
		return "", 0, err
	}

	err = os.Chdir(dir)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = os.Chdir(pwd) }()

	var out bytes.Buffer

	// Run the script
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = &out
	tm := time.Now()
	err = cmd.Start()
	if err != nil {
		return out.String(), 0, err
	}
	err = cmd.Wait()
	return out.String(), time.Since(tm), err
}

func Lookup(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return value, fmt.Errorf("missing %s environment variable", key)
	}

	return value, nil
}

func LookupInt(key string) (int, error) {
	value, err := Lookup(key)
	if err != nil {
		return -1, err
	}

	return ParseInt(value)
}

func ParseInt(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return -1, err
	}
	return int(n), nil
}

func Files(paths ...string) string {
	var files []string
	for _, path := range paths {
		fs, err := script.FindFiles(path).Slice()
		if err != nil {
			continue
		}

		files = append(files, fs...)
	}

	return strings.Join(files, "\n")
}

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "automation by bmbeck@gmail.com")

	switch site.ID {
	case "advent-of-code":
		request.AddCookie(&http.Cookie{
			Name:  "session",
			Value: string(site.Session),
		})

	case "everybody-codes":
		request.AddCookie(&http.Cookie{
			Name:  "everybody-codes",
			Value: string(site.Session),
		})
	}

	return script.Do(request).Bytes()
}

func FetchJSON[T any](url string) (T, error) {
	var t T

	bs, err := Fetch(url)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(bs, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func DecryptAES(s string, key string) (string, error) {
	cs, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	ks := []byte(key)
	iv := ks[:16]

	block, err := aes.NewCipher(ks)
	if err != nil {
		return "", err
	}

	bs := make([]byte, len(cs))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(bs, cs)

	// Remove padding
	n := int(bs[len(bs)-1])
	bs = bs[:len(bs)-n]

	return string(bs), nil
}
