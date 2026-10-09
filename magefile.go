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

	"github.com/bitfield/script"
	"github.com/magefile/mage/mg"
)

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
	files := Files("lib", site.Directory(problem), site.InputFilename(problem))

	_, err := script.
		Echo(files).
		WithEnv(append(os.Environ(), []string{
			fmt.Sprintf("SITE=%s", site.ID()),
			fmt.Sprintf("YEAR=%d", problem.Year),
			fmt.Sprintf("DAY=%d", problem.Day),
			fmt.Sprintf("PART=%d", problem.Part),
		}...)).
		Exec("entr -c sh -c 'go tool mage run'").
		Stdout()

	return err
}

// Next will create and populate the working directory for the next program to
// work on.
//
//goland:noinspection GoUnusedExportedFunction
func Next() error {
	mg.Deps(ParseEnv)

	next := problem
	switch {
	case problem.Day == site.NumDays(problem.Year) && problem.Part == site.NumParts(problem.Year, problem.Day):
		next.Year++
		next.Day = 1
		next.Part = 1
	case problem.Part == site.NumParts(problem.Year, problem.Day):
		next.Day++
		next.Part = 1
	default:
		next.Part++
	}

	// The new directory we're going to work in.
	dir := site.Directory(next)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// The source of what we're going to copy as our main.go.
	var source string
	if next.Part == 1 {
		source = site.TemplateFilename()
	} else {
		previous := next
		previous.Part--
		source = fmt.Sprintf("%s/main.go", site.Directory(previous))
	}

	filename := fmt.Sprintf("%s/main.go", dir)
	if script.IfExists(filename).Error() == nil {
		return fmt.Errorf("filename %s already exists", filename)
	}
	if _, err := script.File(source).WriteFile(filename); err != nil {
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

	expected, err := script.File(site.SolutionsFilename()).
		FilterScan(func(line string, w io.Writer) {
			// Skip over blank lines
			if strings.TrimSpace(line) == "" {
				return
			}

			var buf bytes.Buffer

			// Parse the year/day/part prefix on each line
			fields := strings.Split(line, " ")
			year, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(fmt.Sprintf("unable to parse year: %s", fields[0]))
			}
			day, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(fmt.Sprintf("unable to parse day: %s", fields[1]))
			}
			part, err := strconv.Atoi(fields[2])
			if err != nil {
				panic(fmt.Sprintf("unable to parse part: %s", fields[2]))
			}

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

	if len(aLines) > 0 && reflect.DeepEqual(aLines, eLines) {
		fmt.Printf("✅ SITE=%s YEAR=%d DAY=%02d PART=%d %s [%dms]\n", site.ID(), problem.Year, problem.Day, problem.Part, aLines[0], duration.Milliseconds())
	} else {
		fmt.Printf("❌ SITE=%s YEAR=%d DAY=%02d PART=%d [%dms]\n", site.ID(), problem.Year, problem.Day, problem.Part, duration.Milliseconds())
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

	startHour, startMinute, startSecond := site.StartTime()

	for {
		now := time.Now()
		start := time.Date(
			now.Year(), now.Month(), now.Day(),
			startHour, startMinute, startSecond, 0,
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

	p := problem
	for p.Part = 1; p.Part <= site.NumParts(p.Year, p.Day); p.Part++ {
		// Check if a main.go file exists
		if script.IfExists(site.Directory(p)+"/main.go").Error() == nil {
			fmt.Printf("%d %d %d\n", p.Year, p.Day, p.Part)
		}
	}
}

// ListYear enumerates all days and parts that are expected for the year.
//
//goland:noinspection GoUnusedExportedFunction
func ListYear() {
	mg.Deps(ParseEnv)

	p := problem
	for p.Day = 1; p.Day <= site.NumDays(p.Year); p.Day++ {
		for p.Part = 1; p.Part <= site.NumParts(p.Year, p.Day); p.Part++ {
			// Check if a main.go file exists
			if script.IfExists(site.Directory(p)+"/main.go").Error() == nil {
				fmt.Printf("%d %d %d\n", p.Year, p.Day, p.Part)
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

	switch name {
	case "advent-of-code", "adventofcode", "aoc":
		site = AdventOfCode{}
	case "everybody-codes", "everybodycodes", "ec":
		site = EverybodyCodes{}
	default:
		panic(fmt.Sprintf("unrecognized site: %s", name))
	}

	year, err := LookupInt("YEAR")
	if err != nil {
		// The year wasn't in the environment, infer it from the filesystem.
		for year = time.Now().Year(); year > 0; year-- {
			dir := site.Directory(Problem{Year: year, Day: 1, Part: 1})
			if script.IfExists(dir).Error() == nil {
				break
			}
		}

		if year == 0 {
			panic("unable to infer year")
		}
	}

	day, err := LookupInt("DAY")
	if err != nil {
		// The day wasn't in the environment, infer it from the filesystem.
		for day = site.NumDays(year); day > 0; day-- {
			if script.IfExists(site.Directory(Problem{Year: year, Day: day, Part: 1})).Error() == nil {
				break
			}
		}

		if day == 0 {
			panic("unable to infer day")
		}
	}

	part, err := LookupInt("PART")
	if err != nil {
		// The part wasn't in the environment, infer it from the filesystem.
		for part = site.NumParts(year, day); part > 0; part-- {
			if script.IfExists(site.Directory(Problem{Year: year, Day: day, Part: part})).Error() == nil {
				break
			}
		}

		if part == 0 {
			panic("unable to infer part")
		}
	}
	problem = Problem{Year: year, Day: day, Part: part}
}

// DownloadInput will ensure that the input file for the year, day and part
// being worked on is present in the filesystem.  If the input file is not
// present then an attempt will be made to download it from the site.
func DownloadInput() error {
	mg.Deps(ParseEnv)
	return site.DownloadInput(problem)
}

//
// Sites
//

type Problem struct {
	Year int
	Day  int
	Part int
}

// Site is a place that publishes puzzles.  Everything that differs from site
// to site lives in an implementation of this interface.
type Site interface {
	ID() string
	StartTime() (int, int, int)
	TemplateFilename() string
	SolutionsFilename() string
	NumDays(year int) int
	NumParts(year, day int) int
	Directory(Problem) string
	InputFilename(Problem) string
	DownloadInput(Problem) error
	Authenticate(*http.Request) error
}

//
// Advent of Code
//

type AdventOfCode struct{}

func (site AdventOfCode) ID() string {
	return "advent-of-code"
}
func (site AdventOfCode) StartTime() (int, int, int) { return 23, 0, 0 }
func (site AdventOfCode) TemplateFilename() string {
	return "cmd/advent-of-code/.template"
}
func (site AdventOfCode) SolutionsFilename() string {
	return "cmd/advent-of-code/.solutions"
}
func (site AdventOfCode) NumDays(year int) int {
	// Starting in 2025 Advent of Code moved to 12 days.
	if year >= 2025 {
		return 12
	}
	return 25
}
func (site AdventOfCode) NumParts(year, day int) int {
	// There is no 2nd part on the last day.
	if day == site.NumDays(year) {
		return 1
	}
	return 2
}
func (site AdventOfCode) Directory(p Problem) string {
	return fmt.Sprintf("cmd/advent-of-code/%d/%02d-%d", p.Year, p.Day, p.Part)
}
func (site AdventOfCode) InputFilename(p Problem) string {
	// Input is stored in the part 1 directory.
	dir := site.Directory(Problem{Year: p.Year, Day: p.Day, Part: 1})
	return fmt.Sprintf("%s/input.txt", dir)
}
func (site AdventOfCode) DownloadInput(p Problem) error {
	filename := site.InputFilename(p)

	// First check if the file is already present.
	if script.IfExists(filename).Error() == nil {
		return nil
	}

	// The file wasn't present, download it.
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", p.Year, p.Day)
	bs, err := Fetch(site, url)
	if err != nil {
		return err
	}

	// Save the input
	_, err = script.Echo(string(bs)).WriteFile(filename)
	return err
}
func (site AdventOfCode) Authenticate(request *http.Request) error {
	session, err := script.File("cmd/advent-of-code/.session").String()
	if err != nil {
		return fmt.Errorf("unable to read session file: %w", err)
	}
	session = strings.TrimSpace(session)
	request.AddCookie(&http.Cookie{Name: "session", Value: session})
	return nil
}

//
// Everybody Codes
//

type EverybodyCodes struct{}

func (site EverybodyCodes) ID() string                 { return "everybody-codes" }
func (site EverybodyCodes) StartTime() (int, int, int) { return 17, 0, 0 }
func (site EverybodyCodes) TemplateFilename() string {
	return "cmd/everybody-codes/.template"
}
func (site EverybodyCodes) SolutionsFilename() string {
	return "cmd/everybody-codes/.solutions"
}
func (site EverybodyCodes) NumDays(year int) int       { return 20 }
func (site EverybodyCodes) NumParts(year, day int) int { return 3 }
func (site EverybodyCodes) Directory(p Problem) string {
	return fmt.Sprintf("cmd/everybody-codes/%d/%02d-%d", p.Year, p.Day, p.Part)
}
func (site EverybodyCodes) InputFilename(p Problem) string {
	return fmt.Sprintf("%s/input.txt", site.Directory(p))
}
func (site EverybodyCodes) DownloadInput(p Problem) error {
	filename := site.InputFilename(p)

	// First check if the file is already present.
	if script.IfExists(filename).Error() == nil {
		return nil
	}

	// Load the seed
	type SeedResponse struct {
		Seed int `json:"seed"`
	}

	url := "https://api.everybody.codes/user/me"
	sr, err := FetchJSON[SeedResponse](site, url)
	if err != nil {
		return err
	}

	// Fetch input notes
	type InputNotesResponse struct {
		Part1 string `json:"1"`
		Part2 string `json:"2"`
		Part3 string `json:"3"`
	}

	url = fmt.Sprintf("https://everybody.codes/assets/%d/%d/input/%d.json", p.Year, p.Day, sr.Seed)
	inr, err := FetchJSON[InputNotesResponse](site, url)
	if err != nil {
		return err
	}

	// Fetch AES keys
	type AESKeysResponse struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
		Key3 string `json:"key3"`
	}

	url = fmt.Sprintf("https://api.everybody.codes/event/%d/quest/%d", p.Year, p.Day)
	kr, err := FetchJSON[AESKeysResponse](site, url)
	if err != nil {
		return err
	}

	// Decrypt the input
	var input string
	switch p.Part {
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
func (site EverybodyCodes) Authenticate(request *http.Request) error {
	session, err := script.File("cmd/everybody-codes/.session").String()
	if err != nil {
		return fmt.Errorf("unable to read session file: %w", err)
	}
	session = strings.TrimSpace(session)
	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: session})
	return nil
}

//
// Helpers
//

func RunHelper() (string, time.Duration, error) {
	dir := site.Directory(problem)
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

	return strconv.Atoi(value)
}

func Files(paths ...string) string {
	var files []string
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			continue
		}

		if !info.IsDir() {
			files = append(files, path)
			continue
		}

		fs, err := script.FindFiles(path).Slice()
		if err != nil {
			continue
		}

		files = append(files, fs...)
	}

	return strings.Join(files, "\n")
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

func Fetch(site Site, url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "automation by bmbeck@gmail.com")

	if err := site.Authenticate(request); err != nil {
		return nil, err
	}

	return script.Do(request).Bytes()
}

func FetchJSON[T any](site Site, url string) (T, error) {
	var t T

	bs, err := Fetch(site, url)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(bs, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
