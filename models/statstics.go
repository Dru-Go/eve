package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Stats holds the statstical in of the users
type Stats struct {
	Users []User
	Count int
}

// Set sets the statstical information
func (*Stats) Set(u []User, count int) *Stats {
	return &Stats{
		Users: u,
		Count: count,
	}
}

func (s *Stats) CrawlStats() *Stats {
	var us []User
	users := s.Users
	for _, user := range users {
		if user.Name != "" {
			us = append(us, user)
		}
	}
	return &Stats{Users: us, Count: len(us)}
}

// GetChecked gets all the checked in users
func (s *Stats) GetChecked() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"checkedin": true}).All(&users)
	if err != nil {
		log.Fatal(err)
	}

	return &Stats{Users: users, Count: len(users)}, nil
}
func (s *Stats) GetAll() *Stats {
	u := new(User)
	users, len := u.Get()
	stats := s.Set(users, len)
	return stats
}

// GetVoluntiers gets the statstical information of voluntiers
func (s *Stats) GetVoluntiers() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"role": "Organizing Commitee<"}).All(&users)
	if err != nil {
		log.Fatal(err)
	}

	return &Stats{Users: users, Count: len(users)}, nil
}

// getguests gets the statstical information of guests
func (s *Stats) Getguests() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"role": "Guest"}).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil
}

// getSpeakers gets the statstical information of speakers
func (s *Stats) GetSpeakers() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"role ": "Speakers"}).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil
}

// getMaleguests gets the statstical information of male guests
func (s *Stats) GetMaleguests() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"gender ": 'm'}).All(users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil
}

// getFemaleguests gets the statstical information of female guests
func (s *Stats) GetFemaleguests() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"gender ": 'f'}).All(users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil
}

// getFemaleVolunteres gets the statstical information of female voluntiers
func (s *Stats) GetFemaleVolunteres() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"Gender": 'f'}).Select(bson.M{"role": "Organizing commite"}).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil

}

// getMaleVoluntiers gets the statstical information of male voluntiers
func (s *Stats) GetMaleVoluntiers() (*Stats, error) {
	session, err := Session()
	defer session.Close()

	var users []User
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"Gender": 'm'}).Select(bson.M{"role": "Organizing commite"}).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	ss := s.Set(users, len(users))

	return ss, nil
}

// getAgeStats gets the statstical information of Age
func (s *Stats) GetAgeStats() (*Stats, *Stats, *Stats, error) {

	var user_gt15_ls25 []User
	var user_gt25_ls35 []User
	var user_gt35 []User
	var use *User
	users, _ := use.Get()
	for _, u := range users {
		if u.Age > 15 && u.Age < 25 {
			user_gt15_ls25 = append(user_gt15_ls25, u)
		} else if u.Age > 25 && u.Age < 35 {
			user_gt25_ls35 = append(user_gt25_ls35, u)
		} else {
			user_gt35 = append(user_gt35, u)
		}
	}
	a15_25 := s.Set(user_gt15_ls25, len(user_gt15_ls25))
	a25_35 := s.Set(user_gt25_ls35, len(user_gt25_ls35))
	agt35 := s.Set(user_gt35, len(user_gt35))

	return a15_25, a25_35, agt35, nil
}

// Get Unique visitors

// Get Bounce Rate

//
