package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Attendance saves the arrival history of a user
type Attendance struct {
	Day []time.Time
}

// Set sets an instance of a an Attendance
func (Attendance) Set(u bson.ObjectId, day []time.Time) *Attendance {
	return &Attendance{
		Day: day,
	}
}

func (a *Attendance) SetAttendance(id string) {
	user := new(User)
	use := user.GetID(id)
	use.Tick()

	use.Day = append(use.Day, time.Now())
	use.UpdateUser()
}

func (a *Attendance) GetUserAttendance(phone string) Attendance {
	var attend Attendance
	use := new(User)
	u := use.GetPhone(phone)
	attend = u.Attendance
	return attend
}

func (a *Attendance) GetUserAttendanceID(id string) Attendance {
	var attend Attendance
	use := new(User)
	u := use.GetID(id)
	attend = u.Attendance
	return attend
}

// query gets arrival history of a user
func (a *Attendance) query(phone string) []time.Time {
	tim := a.GetUserAttendance(phone)
	return tim.Day
}

//
