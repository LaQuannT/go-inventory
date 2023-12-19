package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrExceedsAttempts = errors.New("invalid input data - exceeded number of allowed attempts")

func dataPrompt(reader io.Reader, prompt string) (string, error) {
	r := bufio.NewReader(reader)
	fmt.Printf("%s: ", prompt)

	input, err := r.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("dataPrompt: %w", err)
		return "", err
	}

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input, nil
}

func ValidateInputData(reader io.Reader, prompt string, attempt int) (string, error) {
	attemptsAllowed := 2

	input, err := dataPrompt(reader, prompt)
	if err != nil {
		return "", fmt.Errorf("ValidateInputData: %w", err)
	}

	if input == "" && attempt < attemptsAllowed {
		ValidateInputData(reader, prompt, attempt+1)
	} else if input == "" && attempt >= attemptsAllowed {
		return "", ErrExceedsAttempts
	}
	return input, nil
}

func ValidateInputDataChange(reader io.Reader, prompt, originalData string) (string, error) {
	prompt = fmt.Sprintf("%s (%s)", prompt, originalData)
	data, err := dataPrompt(reader, prompt)
	if err != nil {
		return "", fmt.Errorf("ValidateInputDataChange: %w", err)
	}

	if data == "" {
		return originalData, nil
	} else {
		return data, nil
	}
}
