package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var config Config

func main() {
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config = cfg

	app := fiber.New()

	app.Get("/entities/:entityID", handleEntity)
	app.Get("/processes/:entityID/:processID", handleProcess)
	app.Get("/posts/:entityID/:postIdx", handlePost)
	app.Get("/validation/:entityID/:validationToken", handleValidationToken)
	app.Use(HandleNotFound)

	addr := fmt.Sprintf(":%d", config.HTTP.Port)
	log.Fatal(app.Listen(addr))
}

// handleEntity handles the reference to an entity and redirects to its dynamic links
func handleEntity(ctx *fiber.Ctx) error {
	entityID := ctx.Params("entityID")

	dynamicLink := MakeEntityLink(entityID, config)
	ctx.Set("Location", dynamicLink)

	return ctx.SendStatus(fiber.StatusFound) // 302
}

// handleProcess handles the reference to a voting process and redirects to its dynamic links
func handleProcess(ctx *fiber.Ctx) error {
	entityID := ctx.Params("entityID")
	processID := ctx.Params("processID")

	dynamicLink := MakeProcessLink(entityID, processID, config)
	ctx.Set("Location", dynamicLink)

	return ctx.SendStatus(fiber.StatusFound) // 302
}

// handlePost handles the reference to a news post and redirects to its dynamic links
func handlePost(ctx *fiber.Ctx) error {
	entityID := ctx.Params("entityID")
	postIdx := ctx.Params("postIdx")

	dynamicLink := MakePostLink(entityID, postIdx, config)
	ctx.Set("Location", dynamicLink)

	return ctx.SendStatus(fiber.StatusFound) // 302
}

// handleValidationToken handles the reference to a voting process and redirects to its dynamic links
func handleValidationToken(ctx *fiber.Ctx) error {
	entityID := ctx.Params("entityID")
	validationToken := ctx.Params("validationToken")

	dynamicLink := MakeRegistryValidationLink(entityID, validationToken, config)
	ctx.Set("Location", dynamicLink)

	return ctx.SendStatus(fiber.StatusFound) // 302
}

// HandleNotFound redirects to the fallback URL
func HandleNotFound(ctx *fiber.Ctx) error {
	ctx.Set("Location", config.Link.Fallback)

	return ctx.SendStatus(fiber.StatusFound) // 302
}
