package controllers

import (
	"net/http"
	"strings"
	"url-shortener/storage"

	"github.com/gofiber/fiber/v2"
)

type Shortener struct {
    storage *storage.MongoDBStorage
    numShortenedURLChars int
}

func URLShortener(storage *storage.MongoDBStorage, numShortenedURLChars int) *Shortener {
    return &Shortener{storage: storage, numShortenedURLChars: numShortenedURLChars}
}

func (u *Shortener) ShortenURL(c *fiber.Ctx) error {
    var request struct {
        LongURL string `json:"longURL"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
            "details": err.Error(),
        })
    }

    longURL := request.LongURL

    shortURL, err := u.storage.Shorten(longURL, u.numShortenedURLChars)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to shorten URL",
            "details": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "shortURL": shortURL,
    })
}

func (u *Shortener) RedirectURL(c *fiber.Ctx) error {
    shortURL := c.Params("shortURL")

    longURL, err := u.storage.GetOriginal(shortURL)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Short URL not found",
            "details": err.Error(),
        })
    }

    if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
        longURL = "http://" + longURL
    }

    return c.Redirect(longURL, http.StatusSeeOther)
}
