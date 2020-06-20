package question

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func QuestionYesNo(question string) bool {
	choices := []string{"Yes", "No"}
	if QuestionChoices(question, choices) == choices[0] {
		return true
	}
	return false
}

func QuestionChoices(question string, choices []string) string {
	result := ""
	disp_choices := "[" + strings.Join(choices, "/") + "] "

	fmt.Print(question + disp_choices)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input_text := scanner.Text()

		if isArrayContainsString(choices, input_text) {
			result = input_text
			break
		} else {
			fmt.Println(disp_choices + "で答えてください")
			fmt.Print(question + disp_choices)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func isArrayContainsString(array []string, str string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}
	return false
}
