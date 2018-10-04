package models

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Question hold the info about the Questions asked duringg the event
type Question struct {
	Lock int // lock helps to lock the current session till allowed other wise
	User
	Que  string
	Time time.Time
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// Get gets all the Questions asked
func (q *Question) Get() ([]Question, error) {
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return nil, nil
}

// Set asks the questions
func (q *Question) Set(lock int, img int, ques string, time time.Time) error {
	que := Question{Lock: lock, Que: ques, Time: time}
	err := client.Set("key", que, 0).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

// GetRescent gets the most resient question
func (q *Question) GetRescent(time time.Time) []Question {
	return nil
}
