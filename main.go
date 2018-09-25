package main

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/dru/EventManager/api"
	"github.com/dru/EventManager/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

var tpl *template.Template

func main() {
	port := ":8080"
	app := iris.New()

	app.StaticWeb("/public", "./public")

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(logger.New())
	app.RegisterView(iris.HTML("./templates/html/", ".html").Reload(true))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("another.html")
	})
	app.Get("/form", func(ctx iris.Context) {
		ctx.View("home.html")
	})
	app.Post("/add", AddUsers)
	app.Get("/users", Getusers)
	app.Get("/guests", GetGuests)

	app.Get("/checked", GetChecked)
	app.Get("/voluntiers", GetVoluntiers)
	app.Get("/speakers", GetSpeaker)

	app.Get("/api/users", api.GetusersApi)
	app.Post("/api/adduser", AddUsers)
	app.Put("/api/present/:id", api.PresentApi)
	app.Get("/api/checklist", home)
	app.Get("/api/count", home)
	app.Get("/api/getuser/:name", home)
	app.Get("/api/user/:mail", home)

	// listen and serve on http://0.0.0.0:1988.
	app.Run(iris.Addr(port))
}
func home(ctx iris.Context) {
	ctx.View("home.html")
}

func Getusers(ctx iris.Context) {
	// Get users from the database
	var s = new(models.Stats)
	stats := s.GetAll()
	// display the ouput
	ctx.ViewData("stats", stats.CrawlStats())

	// ctx.ViewData("users", stats.Users)
	// ctx.JSON(users)
	ctx.View("report.html")
}

func AddUsers(ctx iris.Context) {
	// Get Context info

	name := ctx.FormValue("icon_prefix")
	email := ctx.FormValue("email")
	phone := ctx.FormValue("phone-no")
	sex := ctx.FormValue("sex")
	valage := ctx.FormValue("age")
	role := ctx.FormValue("role")
	// parse and validate the result
	u := new(models.User)
	if name == "" || email == "" || phone == "" || role == "" {
		ctx.Write([]byte("Error can not parse Empty data"))
		ctx.ViewData("errors", "Error can not parse Empty data")
		return
	}
	if err := u.ValidateName(name); err != nil {
		ctx.Write([]byte(name))
		ctx.Write([]byte("Error: " + err.Error()))
		ctx.ViewData("errors", "Error: Name"+name+"already taken")
		return
	}
	if err := u.ValidateEmail(email); err != nil {
		ctx.Write([]byte("Error Email" + email + "already taken"))
		ctx.ViewData("errors", "Error Email"+email+"already taken")
		return
	}

	ages, _ := strconv.Atoi(valage)
	if err := u.ValidateAge(ages); err != nil {
		ctx.Write([]byte(err.Error()))

		ctx.ViewData("errors", err)
		return
	}

	if err := u.ValidateGender(sex); err != nil {
		ctx.Write([]byte(err.Error()))
		ctx.ViewData("errors", err)
		return
	}

	// insert the data
	user := u.Set(name, email, phone, role, ages, sex)
	user.Save()

	ctx.View("another.html")
}

func profileByUsername(ctx iris.Context) {
	name := ctx.Params().Get("name")

	// .Params are used to get dynamic path parameters.
	s := new(models.Stats)
	u := new(models.User)

	users := u.GetName(name)
	stst := s.Set(users, len(users))
	ctx.ViewData("usernames_stats", stst)
	// renders "./views/user/profile.html"
	// with {{ .Username }} equals to the username dynamic path parameter.
	ctx.View("user/profile.html")
}

func profileByEmail(ctx iris.Context) {
	name := ctx.Params().Get("mail")

	// .Params are used to get dynamic path parameters.
	u := new(models.User)

	user := u.GetEmail(name)
	ctx.ViewData("useremail", user)
	// renders "./views/user/profile.html"
	// with {{ .Username }} equals to the username dynamic path parameter.
	ctx.View("user/profile.html")
}

func add(ctx iris.Context) {
	// extract user from the user form
	// adds the user to database
}

func GetGuests(ctx iris.Context) {
	// get all the statstical informations
	stat := new(models.Stats)
	guests, err := stat.Getguests()
	if err != nil {
		ctx.ViewData("guests_error", err)
	}
	ctx.ViewData("stats", guests)
	fmt.Printf("Count: %v \n", guests.Count)
	ctx.View("Guest.html")
}

func GetChecked(ctx iris.Context) {
	// get all the statstical informations
	stat := new(models.Stats)
	checkedstats, err := stat.GetChecked()
	if err != nil {
		ctx.ViewData("checkedstats_error", err)
	}
	ctx.ViewData("stats", checkedstats)
	fmt.Printf("Count: %v \n", checkedstats.Count)

	ctx.View("CheckedIn.html")
}

func GetSpeaker(ctx iris.Context) {
	// get all the statstical informations
	stat := new(models.Stats)

	speakerStats, err := stat.GetSpeakers()
	if err != nil {
		ctx.ViewData("speakerStats_error", err)
	}
	ctx.ViewData("stats", speakerStats)
	fmt.Printf("Count: %v \n", speakerStats.Count)

	ctx.View("Speaker.html")
}

func GetVoluntiers(ctx iris.Context) {
	// get all the statstical informations
	stat := new(models.Stats)

	voluntirestats, err := stat.GetVoluntiers()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.ViewData("voluntirestats_error", err)
		return
	}
	ctx.ViewData("stats", voluntirestats)

	// if err := ctx.View("Voluntiers.html"); err != nil {
	// 	ctx.StatusCode(iris.StatusInternalServerError)
	// 	ctx.Writef(err.Error())
	// }
	fmt.Printf("Count: %v \n", voluntirestats.Count)
	// ctx.Writef("count:",voluntirestats.Count)
	ctx.View("Voluntiers.html")
}

func report(ctx iris.Context) {
	// get all the statstical informations
	stat := new(models.Stats)

	gueststats, err := stat.Getguests()
	if err != nil {
		ctx.ViewData("gueststats_error", err)
	}
	femaleStats, err := stat.GetFemaleguests()
	if err != nil {
		ctx.ViewData("femaleStats_error", err)
	}
	MaleGuestStats, err := stat.GetMaleguests()
	if err != nil {
		ctx.ViewData("MaleGuestStats_error", err)
	}
	// FemaleVoluntiersStats, err := stat.GetFemaleVolunteres()
	// if err != nil {
	// 	ctx.ViewData("FemaleVoluntiersStats_error", err)
	// }
	// MaleVoluntiersStats, err := stat.GetMaleVoluntiers()
	// if err != nil {
	// 	ctx.ViewData("MaleVoluntiersStats_error", err)
	// }
	AgStat1, AgStat2, AgStat3, err := stat.GetAgeStats()
	if err != nil {
		ctx.ViewData("Age_error", err)
	}
	ctx.ViewData("Guest_Stats", gueststats.Count)
	ctx.ViewData("FemaleGuests_Stats", femaleStats.Count)
	ctx.ViewData("MaleGuest_Stats", MaleGuestStats.Count)
	ctx.ViewData("AgStat1_Stats", AgStat1.Count)
	ctx.ViewData("AgStat2_Stats", AgStat2.Count)
	ctx.ViewData("AgStat3_Stats", AgStat3.Count)

	// show stats
	ctx.View("user/reports.html")
}

func attendance(ctx iris.Context) {
	// get attendance info
	// show attendance
}
