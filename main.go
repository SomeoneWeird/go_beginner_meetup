package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au

func main() {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au")

	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading weather: %v", err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response: %s", body)

}
