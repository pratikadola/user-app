package app

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/user-assignment/app/config"
	"github.com/user-assignment/app/db"
	"github.com/user-assignment/app/db/migration"
	"github.com/user-assignment/app/handler"
	"github.com/user-assignment/app/service"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}

func ConfigureAndRunApp(path string) {
	app := new(App)
	app.Config = config.LoadConfig(path)
	userDB := db.NewUserDB(app.Config)
	defer userDB.DB.Close()
	migration.CreateTables(userDB.DB)
	service := service.NewService(userDB)
	controller := handler.NewUserController(service)

	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/user", controller.CreateUser)
	e.PUT("/user", controller.UpdateUser)
	e.GET("/all", controller.GelAllUsers)
	e.DELETE("/user", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8089"))

}
