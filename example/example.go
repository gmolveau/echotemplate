package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gmolveau/echotemplate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed views/*
var Views embed.FS

//go:embed static/*
var Static embed.FS

func GetStaticFS() http.FileSystem {
	fsys, _ := fs.Sub(Static, "static")
	return http.FS(fsys)
}

var TplConfig = echotemplate.TemplateConfig{
	Root:         "views",
	Extension:    ".html",
	Master:       "layouts/master",
	Partials:     []string{},
	DisableCache: false,
	Delims:       echotemplate.Delims{Left: "{{", Right: "}}"},
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = echotemplate.NewWithConfigEmbed(Views, TplConfig)

	e.GET("/", Index)

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(GetStaticFS()))))

	e.Logger.Fatal(e.Start(":1323"))
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK,
		"index",
		echo.Map{},
	)
}
