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

	req.Header.Add("cookie", os.Getenv("COOKIE"))

	resp, err := http.DefaultClient.Do(req)

	check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	check(err)

	bodyStr := strings.Trim(string(body), "\n")

	if resp.StatusCode >= 400 {
		return "", errors.New(fmt.Sprintf("Received HTTP error %d: %s", resp.StatusCode, bodyStr))
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
		return "", errors.New(fmt.Sprintf("Cannot find input file %s", path))
	} else {
		panic(err)
	}
}

func LoadInput(day int) string {
	httpData, err := fetchInput(day)

	if err == nil {
		fmt.Println("Input loaded from HTTP request")
		return httpData
	}

	fmt.Println(err.Error())
	fmt.Println("Trying to parse input file...")

	fileData, err := readInput(day)

	check(err)

	fmt.Println("Input loaded from file")
	return strings.Trim(fileData, "\n")
}
