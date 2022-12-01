package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tlaceby/go-utils/colors"

	"github.com/joho/godotenv"
)

type AdventOfCode struct {
	SessionID string
	Year      int
	Day       int
}

func main() {
	session, err := GetAOCSession()
	if err != nil {
		panic(err)
	}

	ac := AdventOfCode{session, 2022, 1}
	input, err := GetAOCInput(ac)

	if err != nil {
		panic(err)
	}

	folder := fmt.Sprintf("day_%d", ac.Day)
	err = os.Mkdir(folder, os.ModePerm)
	if err != nil {
		fmt.Printf("%s folder already exists %s\n", colors.BoldRed("Error"), folder)
		os.Exit(1)
	}

	file, err := os.Create(fmt.Sprintf("%s/input.txt", folder))
	if err != nil {
		fmt.Printf("%s %s file could not be created.\n", colors.BoldRed("Error"), folder+"/input.txt")
		os.Exit(1)
	}

	written, err := file.Write(input)

	if err != nil {
		panic(err)
	}

	if written != len(input) {
		fmt.Printf("%s writing input file\n", colors.BoldRed("Error"))
		os.Exit(1)
	}

	GenerateMainFile(ac, string(input), folder)

	fmt.Printf("%s: generated dir at %s\n", colors.BoldGreen("Success"), folder)
}

func GenerateMainFile(ac AdventOfCode, input string, dir string) {
	var contents = fmt.Sprintf(`package main
	import (
		"github.com/tlaceby/go-utils/fs"
	)
	
	const INPUT_FILE = "input.txt"

	func main() {
		println("AdventOfCode %d-%d\n")
		input, _ := fs.ReadTextFile(INPUT_FILE)
		println(input[:20])
	}
	`, ac.Year, ac.Day)

	file, err := os.Create(fmt.Sprintf("%s/main.go", dir))
	if err != nil {
		fmt.Printf("%s %s file could not be created.\n", colors.BoldRed("Error"), dir+"/main.go")
		os.Exit(1)
	}

	file.WriteString(contents)
}

func GetAOCSession() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	const SESSION_NAME = "AOC_SESSION_ID"
	session := os.Getenv(SESSION_NAME)

	if len(session) == 0 {
		return "", fmt.Errorf("%s could not be located inside .env file. Please make sure env file has a %s field", SESSION_NAME, SESSION_NAME)
	}

	return session, nil
}

func GetAOCInput(av AdventOfCode) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", av.Year, av.Day)
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	sessionCookie := &http.Cookie{Name: "session", Value: av.SessionID}
	req.AddCookie(sessionCookie)

	fmt.Printf("%s: %s\n", colors.BoldCyan("Request"), url)
	res, err := client.Do(req)

	if err != nil {
		return make([]byte, 0), err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return make([]byte, 0), err
	}

	return body, nil
}
