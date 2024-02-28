package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		print("No .env file found")
	}
}

type Geolocation [1]struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func main() {
	appid := os.Getenv("APP_ID")
	if appid == "" {
		panic("AUTH_TOKEN not set")
	}

	res, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=Dorno&appid=" + appid + "&lang=it&limit=1")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API request failed")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))

	var geolocation Geolocation
	err = json.Unmarshal(body, &geolocation)

	if err != nil {
		panic(err)
	}

	println(geolocation[0].Lat, geolocation[0].Lon)

	/* res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Dorno&&appid=")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API request failed")
	} */
}
