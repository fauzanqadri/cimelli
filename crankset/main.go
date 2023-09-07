package main

import (
	"crankset/controllers"
)

func main() {
	// defer models.Db.Close()

	// err := models.Db.Ping()

	// if err != nil {
	// 	panic(err)
	// }

	controllers.Router.Run()
}
