package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au

func main() {
	body, err := getWeatherResponseBody()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %s", body)
}

func getWeatherResponseBody() ([]byte, error) {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au")

	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading weather: %v", err)
		return []byte(""), err
	}

	defer resp.Body.Close()

	return body, nil

}

type City struct {
	Weather Weather `json:"main"`
	Name    string  `json:"name"`
}

type Weather struct {
	CurrentTemp float64 `json:"temp"`
	MaxTemp     float64 `json:"mtepm_max"`
}
