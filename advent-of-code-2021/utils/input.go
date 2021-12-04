package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchInput(day int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day), nil)

	check(err)

	req.Header.Add("cookie", fmt.Sprintf("session=%s", os.Getenv("SESSION_ID")))

	resp, err := http.DefaultClient.Do(req)

	check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	check(err)

	bodyStr := strings.Trim(string(body), "\n")

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("received HTTP error %d: %s", resp.StatusCode, bodyStr)
	}

	return bodyStr, nil
}

func readInput(day int) (string, error) {
	path := filepath.Join("input", fmt.Sprintf("day%d.txt", day))

	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		check(err)
		return string(data), nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("cannot find input file %s", path)
	} else {
		panic(err)
	}
}

func LoadInput(day int) string {
	fileData, err := readInput(day)

	if err == nil {
		fmt.Println("Input loaded from file")
		return strings.Trim(fileData, "\n")
	}

	fmt.Println(err.Error())
	fmt.Println("Trying to fethc input from AOC website...")

	httpData, err := fetchInput(day)

	check(err)

	fmt.Println("Input loaded from HTTP request")
	return strings.Trim(httpData, "\n")
}
