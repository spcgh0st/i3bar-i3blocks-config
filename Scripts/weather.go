package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type conditionObject struct {
			Text string `json:"text"`
		}

type weatherResponse struct {

	Current struct {
		Temp float32 `json:"temp_c"`
		FeelsLike float32 `json:"feelslike_c"`
		Precip float32 `json:"precip_mm"`
		Condition conditionObject `json:"condition"`

	}

	Forecast struct {

		ForecastDay []struct {

			Date string `json:"date"`
			Day struct {
				MaxTemp float32 `json:"maxtemp_c"`
				MinTemp float32 `json:"mintemp_c"`
				TotalPrecip float32 `json:"totalprecip_mm"`
				DailyChanceOfRain int `json:"daily_chance_of_rain"`
				Condition conditionObject `json:"condition"`
			}

		}
	}

}


func main() {

	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=" +
		"REPLACE_WITH_API_KEY" + "&q=Buenos%Aires&days=2&aqi=no&alerts=no")

	if err != nil {
		os.Exit(1)
	} 

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var weather weatherResponse

	_ = json.Unmarshal([]byte(resBody), &weather)

	fmt.Printf("<span foreground=\"cyan\">Current:</span> %s, %.1f°C ᛫ <span foreground=\"cyan\">Today:</span> %s, %.0f/%.0f°C ᛫ <span foreground=\"cyan\">Tomorrow:</span> %s, %.0f/%.0f°C\n",
		weather.Current.Condition.Text, weather.Current.Temp, 
		weather.Forecast.ForecastDay[0].Day.Condition.Text, weather.Forecast.ForecastDay[0].Day.MaxTemp, weather.Forecast.ForecastDay[0].Day.MinTemp,
		weather.Forecast.ForecastDay[1].Day.Condition.Text, weather.Forecast.ForecastDay[1].Day.MaxTemp, weather.Forecast.ForecastDay[1].Day.MinTemp )

}