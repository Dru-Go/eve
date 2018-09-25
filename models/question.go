package models

import "time"

// Question hold the info about the Questions asked duringg the event
type Question struct {
	Lock int // lock helps to lock the current session till allowed other wise
	Que  string
	time time.Time
}

// Get gets all the Questions asked
func (q *Question) Get() ([]Question, error) {
	return nil, nil
}

// Set asks the questions
func (q *Question) Set(img int, que string, time time.Time) error {
	return nil
}

// GetRescent gets the most resient question
func (q *Question) GetRescent(time time.Time) []Question {
	return nil
}
