package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func stringPrompt(prompt string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)

	res, err := r.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("stringPrompt: %w", err)
		log.Fatal(err)
	}
	res = strings.TrimSpace(res)
	return strings.ToLower(res)
}

func ValidateInput(prompt string, trys int) string {
	attemptsAllowed := 2
	input := stringPrompt(prompt)

	if input == "" && trys < attemptsAllowed {
		ValidateInput(prompt, trys+1)
	} else if input == "" && trys >= attemptsAllowed {
		fmt.Printf("'%s' contains no vaild data", prompt)
		os.Exit(0)
	}

	return input
}

func ValidateInputChange(prompt, originalData string) string {
	prompt = fmt.Sprintf("%s (%s)", prompt, originalData)
	data := stringPrompt(prompt)
	if data == "" {
		return originalData
	} else {
		return data
	}
}
