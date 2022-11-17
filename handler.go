package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start(q *QuestionNode) {
	currentQuestion := q

	scanner := bufio.NewScanner(os.Stdin)
	for currentQuestion != nil {
		fmt.Print(q.Question)

		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			if currentQuestion.Default == nil {
				continue
			}

			input = *currentQuestion.Default
		}

		if currentQuestion.Callback != nil {
			currentQuestion.Callback(input)
		}
		currentQuestion = currentQuestion.Next[strings.ToLower(input)]
	}
}
