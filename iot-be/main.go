package main

import (
	"web-ivsr-be/database"
	"web-ivsr-be/handlers"
	inforrepo "web-ivsr-be/repository/infor-repo"
	"web-ivsr-be/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	sql := &database.Sql{
		Host:     "localhost",
		User:     "postgres",
		Port:     5432,
		Password: "postgres",
		Dbname:   "Web",
	}
	sql.Connect()
	defer sql.Close()

	db := handlers.SiteHandler{
		Repo: inforrepo.NewRepo(sql),
	}

	api := router.Api{
		Echo:    e,
		Handler: db,
	}
	api.SetUpRouter()

	e.Start("127.0.0.1:8080")
}
