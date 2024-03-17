package main

import (
	"sluggers/cmd/handlers"
	"sluggers/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "sluggers/docs"
	// _ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	// init the server
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(handlers.LogRequest)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.File("/favicon.ico", "public/favicon.ico")
	e.GET("/", handlers.Home)
	e.GET("/:id", handlers.GetCharacter)
	e.GET("/all", handlers.GetAllCharacters)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// connect to the DB
	db := storage.InitDB()
	defer db.Close()

	e.Logger.Fatal(e.Start(":1337"))
}
