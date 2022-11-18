package state

import "strings"

type QuestionNode struct {
	Question string
	Default  *string
	Next     map[string]*QuestionNode
	Callback func(string)
}

func (q *QuestionNode) AddNext(key string, next *QuestionNode) *QuestionNode {
	if q.Next == nil {
		q.Next = make(map[string]*QuestionNode)
	}
	q.Next[strings.ToLower(key)] = next
	return next
}
