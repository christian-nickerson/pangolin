package main

import (
	"log"
	"strconv"

	"github.com/goccy/go-json"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load settings
	settings, err := configs.LoadSettings("settings")
	if err != nil {
		log.Fatal("Loading settings failed:", err)
	}

	// set up fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// run app
	app.Listen(":" + strconv.Itoa(settings.Server.API.Port))
}
