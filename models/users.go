package models

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var f = "db/test.db"

// User is a structure of users
type User struct {
	Name      string
	Email     string
	PhoneNo   string
	Role      string
	Age       int
	Gender    byte
	Qcode     string
	CheckedIn bool
	Attendance
}

// Set sets an instance of a user
func (u *User) Set(Name, Email, PhoneNo, Role string, Age int, Gender string) *User {
	var gen byte
	if Gender == "m" {
		gen = 'm'
	} else {
		gen = 'f'
	}
	var days []time.Time
	days = append(days, time.Now())

	att := Attendance{Day: days}

	return &User{
		Name:       Name,
		Email:      Email,
		PhoneNo:    PhoneNo,
		Role:       Role,
		Age:        Age,
		Gender:     gen,
		CheckedIn:  true,
		Attendance: att,
	}
}

// Set sets an instance of a user
func (u *User) SetApi(Name, Email, PhoneNo, Role string, Age int, Gender byte) *User {

	var days []time.Time
	days = append(days, time.Now())

	att := Attendance{Day: days}
	return &User{
		Name:       Name,
		Email:      Email,
		PhoneNo:    PhoneNo,
		Role:       Role,
		Age:        Age,
		Gender:     Gender,
		CheckedIn:  false,
		Attendance: att,
	}
}

func (user *User) GenerateQR() string {
	if user.Name == "Name" || user.Name == "" {
		return ""
	}
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("Error :%s", js)
	}

	qrCode, _ := qr.Encode(string(js), qr.H, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	name := user.Email + ".png"

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../public/images", name)
	dst, _ := os.Create(path)

	defer dst.Close()

	// encode the barcode as png
	png.Encode(dst, qrCode)

	return path

}

func (u *User) GetQrImage(email string) string {
	users, _ := u.Get()
	for _, user := range users {
		if user.Email == email {
			return u.Qcode
		}
	}
	return ""
}

func (u *User) GetQrImages(email string) []string {
	users, _ := u.Get()
	var qrs []string
	for _, user := range users {
		qrs = append(qrs, user.Qcode)
	}
	return qrs
}

// Save saves user to the database
func (u *User) Save() {
	// Generate QRcode and save to image
	// Save the user to the database
	session, err := Session()
	defer session.Close()
	fmt.Println(u.Name)
	u.Qcode = u.GenerateQR()
	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")
	err = c.Insert(u)
	if err != nil {
		log.Fatalf("Error Inserting user %s", err)
	}
}

// Tick notify the user has checked in
func (u *User) Tick() {
	u.CheckedIn = true
	u.UpdateUser()
}

// Get gets the users from a database
func (*User) Get() ([]User, int) {
	var users []User
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{}).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	return users, len(users)
}

// GetName gets the users from a database
func (u *User) GetPhone(phoneno string) User {
	var user User
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Find(bson.M{"phoneno": phoneno}).One(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

// GetID gets the users from a database
func (u *User) GetID(id string) User {
	var user User
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")
	if u.Exists(id) {
		err = c.Find(bson.M{"id": id}).One(&user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user
}

// GetName gets the users from a database
func (u *User) GetName(name string) []User {
	var user []User
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")
	err = c.Find(bson.M{"name": name}).All(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

// GetName gets the users from a database
func (u *User) GetEmail(email string) User {
	var user User
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")
	err = c.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (u *User) Exists(id string) bool {
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	num, err := c.Find(bson.M{"_id": id}).Count()
	if err != nil {
		log.Fatal(err)
	}
	if num == 0 {
		return false
	}

	return true
}

func (u *User) UpdateUser() {
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	err = c.Update(bson.M{"phoneno": u.PhoneNo}, u)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *User) ValidateName(name string) error {
	if name == "" || len(name) <= 3 {
		return fmt.Errorf("Error UserName must be a greater than 3 charactors long %s", "")
	}
	// Check if username exists in the database
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("qcode").C("users")

	_, err = c.Find(bson.M{"Name ": name}).Count()
	if err != nil {
		log.Fatal(err)
	}
	// if n != 0 {
	// 	return fmt.Errorf("Error This name already exists %s", n)
	// }
	return nil
}

func (u *User) ValidateEmail(mail string) error {
	if mail == "" { //check email syntax
		return fmt.Errorf("Error Email is not Right %s", "")
	}
	// Check if Email exists in the database
	session, err := Session()
	defer session.Close()

	if err != nil {
		log.Fatalf("Error Sessionizing %s", err)
	}
	c := session.DB("event_management").C("users")

	n, err := c.Find(bson.M{"Email ": mail}).Count()
	if err != nil {
		log.Fatal(err)
	}
	if n > 0 {
		return fmt.Errorf("Error This Email already exists %s", n)
	}
	return nil
}

func (u *User) ValidateAge(age int) error {

	if age < 10 && age > 71 {
		return fmt.Errorf("Error: %s Age  must be between 11 and 70 %s", age)
	}
	return nil
}

func (u *User) ValidateGender(sex string) error {

	if sex == "m" || sex == "f" {
		return nil
	}
	return fmt.Errorf("Error Gender:%s is not Right", string(sex))
}

// Session starts the session for the XORM
func Session() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return session, nil
}
