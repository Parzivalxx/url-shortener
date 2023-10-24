package main

import (
	"fmt"
	"os"
	"strconv"
	"url-shortener/controllers"
	"url-shortener/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not present")
	}
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	numShortenedURLChars := os.Getenv("NUM_SHORTENED_URL_CHARS")
    numChars, err := strconv.Atoi(numShortenedURLChars)

    if err != nil {
        fmt.Printf("Error converting to integer: %v\n", err)
        return
    }

	app := fiber.New()

	storage, err := storage.Storage(mongoURI, dbName, collectionName)
	if err != nil {
		fmt.Printf("Error initializing MongoDB storage: %v\n", err)
		return
	}

	defer storage.Close()

	urlShortener := controllers.URLShortener(storage, numChars)

	app.Use(logger.New())

	app.Post("/shorten", urlShortener.ShortenURL)
	app.Get("/:shortURL", urlShortener.RedirectURL)

	fmt.Println("Server is running on :3000")
	app.Listen(":3000")
}
