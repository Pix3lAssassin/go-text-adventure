package state

func getStory() *QuestionNode {
	startNode := &QuestionNode{
		Question: "You wake up in a field of wheat with the sun shining down on you. Off to the right is a farmhouse and to your left is a hill that you could climb to get an outlook of the surrounding area: ",
	}

	// gameOver := &QuestionNode{
	// 	Question: "Game Over!",
	// }

	barn := &QuestionNode{
		Question: "You enter the farmhouse and find a ",
	}

	startNode.AddNext("farmhouse", barn)

	outlook := &QuestionNode{
		Question: "",
	}

	startNode.AddNext("outlook", outlook)

	return startNode
}
