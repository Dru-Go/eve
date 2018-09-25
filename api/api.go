package api

import (
	"github.com/dru/EventManager/models"
	"github.com/kataras/iris"
)

func GetusersApi(ctx iris.Context) {
	// Get users from the database
	var s = new(models.Stats)
	stats := s.GetAll()
	// display the ouput
	ctx.JSON(&stats.Users)
}
func PresentApi(ctx iris.Context) {
	id := ctx.Params().Get("id")
	use := new(models.User)
	if use.Exists(id) {
		if use.CheckedIn == true {
			ctx.Write([]byte("true"))
			return
		}
		ctx.Write([]byte("false"))
	}
	ctx.Write([]byte("Not Found"))
}

func Checklistapi(ctx iris.Context) {
	// Get users from the database
	var s = new(models.Stats)
	stats, err := s.GetChecked()
	if err != nil {
		ctx.ViewData("errors", "Error Getting Checked users")
		return
	}
	// display the ouput
	ctx.JSON(stats)
}
