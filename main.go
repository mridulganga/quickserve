package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := flag.String("port", "80", "port to run on")
	path := flag.String("path", "./app", "folder to serve")
	cert := flag.String("cert", "", "path to certificate")
	key := flag.String("key", "", "path to key")
	flag.Parse()

	app.Static("/", *path)

	var err error
	if *cert != "" && *key != "" {
		err = app.ListenTLS(":443", *cert, *key)
	} else {
		err = app.Listen(fmt.Sprintf(":%v", *port))
	}

	if err != nil {
		fmt.Println(err)
	}
}
