package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Pix3lAssassin/go-text-adventure/entities"
	"github.com/Pix3lAssassin/go-text-adventure/state"
)

func start(q *state.QuestionNode) {

	player := entities.NewPlayer()

	gameState := state.GameState{
		Player: player,
		Factions: []state.Faction{
			{},
		},
	}

	// TODO Remove this
	fmt.Print(gameState)

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
