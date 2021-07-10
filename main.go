package main

import (
	"github.com/labstack/echo/v4"
	"logic_king_backend/cores"
	"logic_king_backend/routes"
)

func main() {
	e := echo.New()
	cfg := cores.NewConfig()

	mongodb := cores.NewMongo(&cores.MongoOption{
		Username: cfg.Username(),
		Password: cfg.Password(),
		DBName:   cfg.DBName(),
		URL:      cfg.DBHost(),
		Port:     cfg.DBPort(),
	})

	app := cores.NewApp(e, mongodb, cfg)

	routes.NewProjectRoute(app)

	app.Echo.Logger.Fatal(app.Echo.Start(":1234"))
}
