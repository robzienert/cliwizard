package cliwizard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MustReadIn will attempt to read stdin 3 times before fataling aborting.
func MustReadIn() string {
	for i := 0; i < 3; i++ {
		text, err := tryReadStdIn()
		if err == nil {
			return strings.TrimSpace(text)
		}
		// TODO Theme
		fmt.Printf("Error while reading stdin:: %s. Failed %d/3\n", err, i)
	}
	fmt.Println("Could not read stdin. Aborting...")
	os.Exit(1)
	return ""
}

// ReadIn will try to read stdin. If an error occurs, it will be ignored.
func ReadIn() string {
	text, err := tryReadStdIn()
	if err == nil {
		return strings.TrimSpace(text)
	}
	fmt.Printf("Error while reading stdin: %s", err)
	return ""
}

func tryReadStdIn() (text string, err error) {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	return
}
