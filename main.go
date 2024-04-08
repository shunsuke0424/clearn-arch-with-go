package main


// func main() {
// 	e := echo.New()
// 	db, _ := model.DB.DB()
// 	defer db.Close()

// 	e.GET("/users", controller.GetUsers)
// 	e.GET("/users/:id", controller.GetUser)
// 	e.POST("/users", controller.CreateUser)
// 	e.PUT("/users/:id", controller.UpdateUser)
// 	e.DELETE("/users/:id", controller.DeleteUser)
// 	e.Logger.Fatal(e.Start(":8080"))
// }


import (
	"github.com/so-heee/golang-clean-architecture-example/api/infrastructure"
)

func main() {
	infrastructure.Init()
}
