package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// e.AutoTLSManager.Cache = autocert.DirCache(".cache")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/comm/messages", commMessages)
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
		`)
	})

	// Start server
	e.Logger.Fatal(e.StartTLS("127.0.0.1:443", "server.crt", "server.key"))

	// e.Logger.Fatal(e.Start("127.0.0.1:80"))
}

// Handler
func commMessages(c echo.Context) error {
	fmt.Println("On new message.")
	f, err := os.OpenFile("comm.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	text, _ := ioutil.ReadAll(c.Request().Body)
	if _, err = f.Write(text); err != nil {
		panic(err)
	}
	f.WriteString("\n")
	return c.String(http.StatusOK, "Hello, World!")
}
