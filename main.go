package main

import "context"
import "github.com/bondanherumurti-transfez/livestream-go/rest-api/database"
import "github.com/gofiber/fiber/v2"
import "github.com/joho/godotenv"
import "go.mongodb.org/mongo-driver/bson"

func main() {
		err := initApp()
		if err != nil {
				panic(err)
		}

		defer database.CloseMongoDB()

		app := fiber.New()

		app.Post("/", func(c *fiber.Ctx) error {
				sampleDoc := bson.M{"name": "Sample Todo"}
				collection := database.GetCollection("todos")
				nDoc, err := collection.InsertOne(context.Background(), sampleDoc)

				if err != nil {
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
								"message": "Failed to insert document",
						})
				}

				return c.JSON(nDoc)
		})

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()
	if err != nil {
			return err
	}

	return nil
}

func loadENV() error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    return nil
}


