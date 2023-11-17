package main

import (
		"fmt"
		"net/http"
		"os"
		"io"
		"encoding/json"
)


type geoResp struct {

	IP				string	`json:"ip"`
	Country			string	`json:"country_name"`
	City			string	`json:"city"`
	TLD				string	`json:"country_tld"`

}

func main() {

	res, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=" + "REPLACE_WITH_API_KEY")

	if err != nil {
		fmt.Print("Request Error")
		os.Exit(1)
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	var dataGeo geoResp

	_ = json.Unmarshal([]byte(resBody), &dataGeo)


	fmt.Printf("%s (%s, %s)", dataGeo.IP, dataGeo.City, dataGeo.Country)

}