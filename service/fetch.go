package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	Country   string    `json:"country"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	TzId      string    `json:"tz_id"`
	LocalTime time.Time `json:"local_time"`
}

type Current struct {
	LastUpdateEpoch int64     `json:"last_update_epoch"`
	LastUpdated     string    `json:"last_updated"`
	TempC           float64   `json:"temp_c"`
	TempF           float64   `json:"temp_f"`
	IsDay           int8      `json:"is_day"`
	Condition       Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int16  `json:"code"`
}

func GetWheaterService(c *fiber.Ctx) error {
	godotenv.Load()
	APIKEY := os.Getenv("API_KEY_WHATERAPI")
	ENPOINT := os.Getenv("ENDPOINT_WHEATER")

	city := c.Params("city")
	url := fmt.Sprintf("%s/v1/forecast.json?key=%v&q=%s&days=1&aqi=no&alerts=no", ENPOINT, APIKEY, city)
	// http://api.weatherapi.com/v1/forecast.json?key=361cd9ebacfd4d3993d40619222107&q=Yogyakarta&days=1&aqi=no&alerts=no

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		c.Status(500).JSON(fiber.Map{"message": err})
	}

	jsonByte, _ := ioutil.ReadAll(response.Body)
	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
			c.Status(500).JSON(fiber.Map{"message": e})
		}
	}()

	var wheatherResponse WeatherResponse
	er := json.Unmarshal(jsonByte, &wheatherResponse)
	if er != nil {
		log.Fatal(er)
		c.Status(500).JSON(fiber.Map{"message": er})
	}

	return c.Status(200).JSON(wheatherResponse)
}
