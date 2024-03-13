package routes

import ()

import (
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemainig  int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	//implementing ratelimiting by 1 

	//check if the entry is a url 
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid URL",
		})
	}

	//check for domain errors 
	if !helpers.RemoveDomain(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid URL",
		})
	}

	//https cheecker 
	body.URL = helpers.EnforceHTTP(body.URL)
	
}
