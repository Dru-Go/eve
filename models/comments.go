package models

import "time"

type Comment struct {
	UserIMG  int
	UserName string
	comment  string
	time     time.Time
}

func (c *Comment) Get() ([]Comment, error) {
	return nil, nil
}
func (c *Comment) Set(img int, name, Comment string, time time.Time) error {
	// here name is the KEY used
	return nil
}
