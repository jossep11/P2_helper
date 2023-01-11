package main

import (
	"github.com/jossep11/config"
	"github.com/jossep11/routesj"
)

func main() {

	// app := fiber.New()
	config.Connect()
	routesj.Rutas()

	// app.Use("")

	// app.Get("/data", handlers.GetDatas)
	// app.Get("/data/:id", handlers.GetData)
	// app.Post("/data", handlers.AddData)
	// app.Put("/data/:id", handlers.UpdateData)
	// app.Delete("/data/:id", handlers.RemoveData)

	// data := app.Group("/data")
	// data.Get("/", handlers.GetDatas)
	// data.Get("/:id", handlers.GetData)
	// data.Post("/", handlers.AddData)
	// data.Put("/:id", handlers.UpdateData)
	// data.Delete("/:id", handlers.RemoveData)

	// app.Route("/data", func(router fiber.Router) {
	// 	router.Get("/", handlers.GetDatas)
	// 	router.Delete("/:id", handlers.RemoveData)
	// })

	// log.Fatal(app.Listen(":3000"))
}
