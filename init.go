package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runInit() error {
	_, e := os.Stat("private")
	if e == nil {
		return nil
	}
	if !os.IsNotExist(e) {
		return fmt.Errorf("error getting information about the private directory: %w", e)
	}
	if e := os.Mkdir("private", 0700); e != nil {
		return fmt.Errorf("error creating private directory: %w", e)
	}
	if e := initCookie(); e != nil {
		return fmt.Errorf("error initializing cookie: %w", e)
	}
	if e := initAnswers(); e != nil {
		return fmt.Errorf("error initializing answers: %w", e)
	}
	return nil
}

func initCookie() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your session cookie: ")
	scanner.Scan()
	cookie := strings.TrimSpace(scanner.Text())
	if e := scanner.Err(); e != nil {
		return fmt.Errorf("error reading cookie from stdin: %w", e)
	}
	cookieFile, e := os.Create("private/cookie.txt")
	if e != nil {
		return fmt.Errorf("error creating cookie file: %w", e)
	}
	defer cookieFile.Close()
	_, e = cookieFile.WriteString(cookie)
	if e != nil {
		return fmt.Errorf("error writing cookie to file: %w", e)
	}
	return nil
}

func initAnswers() error {
	answersFile, e := os.Create("private/answers.json")
	if e != nil {
		return fmt.Errorf("error creating answers file: %w", e)
	}
	defer answersFile.Close()
	_, e = answersFile.WriteString("{}")
	if e != nil {
		return fmt.Errorf("error writing empty answers to file: %w", e)
	}
	return nil
}
