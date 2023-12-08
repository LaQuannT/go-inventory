package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func StringPrompt(prompt string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)

	res, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	res = strings.TrimSpace(res)
	return strings.ToLower(res)
}
