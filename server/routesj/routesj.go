package routesj

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"

	// "github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/jossep11/handlers"
)

func Rutas() {
	app := fiber.New()
	// limit := Limiter.New(limiter.Config{Max: 1, Expiration: time.Second})
	// app.Use(Test)

	// Data
	app.Use(cors.New())

	data := app.Group("/data")
	data.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	data.Get("/", handlers.GetDatas)
	data.Get("/:id", handlers.GetData)
	data.Post("/", handlers.AddData)
	data.Put("/:id", handlers.UpdateData)
	data.Delete("/:id", handlers.RemoveData, handlers.GetDatas)

	// Users
	users := app.Group("/users")
	users.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Post("/", handlers.AddUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.RemoveUser, handlers.GetUsers)

	// Login
	app.Post("/login", handlers.Login)

	log.Fatal(app.Listen(":8080"))

}

// func Test(c *fiber.Ctx) error {
// 	fmt.Println("ok, keep going")
// 	return c.Next()
// }
