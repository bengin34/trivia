// package main

// import (
//     "github.com/bengin34/trivia/database"
//     "github.com/gofiber/fiber/v2"
//     "github.com/gofiber/fiber/v2/middleware/cors"
// )

// func main() {
//     database.ConnectDb()
//     app := fiber.New()

//     setupRoutes(app)

   

//     app.Listen(":3000")
// }

package main

import (
    "github.com/bengin34/trivia/database"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    // Enable CORS with custom configuration
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3001",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    setupRoutes(app)

    app.Listen(":3000")
}
