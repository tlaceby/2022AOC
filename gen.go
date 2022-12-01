package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/tlaceby/go-utils/colors"

	"github.com/joho/godotenv"
)

type AdventOfCode struct {
	SessionID string
	Year      int
	Day       int
}

func GenerateYearData(ac AdventOfCode) {

	input, err := GetAOCInput(ac)

	if err != nil {
		return
	}

	folder := fmt.Sprintf("day_%d", ac.Day)
	err = os.Mkdir(folder, os.ModePerm)

	if err != nil {
		return
	}

	file, _ := os.Create(fmt.Sprintf("%s/input.txt", folder))
	file.Write(input)
	GenerateMainFile(ac, string(input), folder)

	fmt.Printf("+ %s\n", colors.FCyan(folder))
}

func main() {
	session, err := GetAOCSession()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup

	for i := 1; i <= 25; i++ {
		wg.Add(1)
		go func(indx int) {
			defer wg.Done()
			ac := AdventOfCode{session, 2022, indx}
			GenerateYearData(ac)
		}(i)
	}

	wg.Wait()

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
	res, err := client.Do(req)

	if err != nil {
		return make([]byte, 0), err
	}

	body, err := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return make([]byte, 0), fmt.Errorf("days data is not ready to be retrieved")
	}

	if err != nil {
		return make([]byte, 0), err
	}

	return body, nil
}
