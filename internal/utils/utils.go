package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func StringPrompt(prompt string) string {
	var res string
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprintf(os.Stdout, "%s: ", prompt)
		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if res != "" {
			break
		}
	}

	return strings.TrimSpace(res)
}
