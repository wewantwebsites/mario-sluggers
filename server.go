package main

import (
	"os"

	"sluggers/cmd/handlers"
	"sluggers/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
    // init the server
    e := echo.New()
    e.Use(handlers.LogRequest)
    e.Use(middleware.Logger())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))
    e.GET("/", handlers.Home)
    // connect to the DB
    db := storage.InitDB()
    defer db.Close()
    
    // check to see if the migrate flag has been passed to update the db
    args := os.Args[1:]
    if len(args) > 0 && os.Args[1] == "m" {
        storage.Migrate()
    }

    e.Logger.Fatal(e.Start(":1337"))
}
