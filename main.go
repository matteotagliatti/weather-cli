package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		print("No .env file found")
	}
}

func main() {
	appid := getToken()
	lat, lon := getGeolocation(appid)
	getWheather(appid, lat, lon)

}

func getToken() string {
	appid := os.Getenv("APP_ID")
	if appid == "" {
		panic("APP_ID not set in .env file")
	}

	return appid
}

type Geolocation [1]struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func getGeolocation(appid string) (float64, float64) {
	res, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=Dorno&appid=" + appid + "&lang=en&limit=1")

	if err != nil {
		panic(err) // stop the program and print out error
	}

	defer res.Body.Close() // close the connection

	if res.StatusCode != 200 {
		panic("API request failed")
	}

	body, err := io.ReadAll(res.Body) // read the response body

	if err != nil {
		panic(err)
	}

	var geolocation Geolocation
	err = json.Unmarshal(body, &geolocation) // convert the json to a struct

	if err != nil {
		panic(err)
	}

	return geolocation[0].Lat, geolocation[0].Lon
}

type Wheather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func getWheather(appid string, lat float64, lon float64) {
	// convert float64 to string
	lat_string := strconv.FormatFloat(lat, 'f', -1, 64)
	lon_string := strconv.FormatFloat(lon, 'f', -1, 64)

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + lat_string + "&lon=" + lon_string + "&appid=" + appid + "&lang=en&units=metric")

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

	var wheather Wheather
	err = json.Unmarshal(body, &wheather)

	if err != nil {
		panic(err)
	}

	print("The weather in " + wheather.Name + ", " + wheather.Sys.Country + " is " + wheather.Weather[0].Description)
}
