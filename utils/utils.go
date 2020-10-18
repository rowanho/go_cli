package utilsCli

import (
	"bufio"
	"fmt"
	"os"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Handles a yes/no answer
// Returns true if yes, otherwise false
func YesNoPrompt(question string) bool {
	yesStrings := []string{"y", "Y", "yes"}
	noStrings := []string{"n", "N", "no"}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	for true {
		if StringInSlice(text, yesStrings) {
			return true
		} else if StringInSlice(text, noStrings) {
			return false
		} else {
			fmt.Print("Enter y/n: ")
			text, _ = reader.ReadString('\n')
		}
	}
	return false
}
