package main

import (
	"fmt"
	"net/http"
)

// http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au

func main() {

	_, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au")

	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return
	}

}
