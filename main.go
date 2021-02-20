package main

import (
	"book-api/infrastructure/repository"
	"book-api/infrastructure/router"
	"book-api/registry"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	db := repository.NewDB("./books")
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing DB")
		}
	}()

	r := registry.NewRegistry(db)

	app := fiber.New()

	app = router.NewRouter(app, r.NewBookController())

	log.Fatal(app.Listen(":3000"))

}
