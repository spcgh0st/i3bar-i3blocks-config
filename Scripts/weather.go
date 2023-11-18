package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
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

	const apiKey string = "REPLACE_WITH_API_KEY"

	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=" +
		apiKey + "&q=Buenos%Aires&days=2&aqi=no&alerts=no")

	if err != nil {
		os.Exit(1)
	} 

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var weather weatherResponse

	_ = json.Unmarshal([]byte(resBody), &weather)

	fmt.Printf("%s, %.1f°C ᛫ %s ᛫ %s\n",
		weather.Current.Condition.Text, weather.Current.Temp, 
		formatForecastDay(weather, 0),
		formatForecastDay(weather, 1))

}




func formatForecastDay(weather weatherResponse, index int) string {

	condition := weather.Forecast.ForecastDay[index].Day.Condition.Text
	maxTemp := weather.Forecast.ForecastDay[index].Day.MaxTemp
	minTemp := weather.Forecast.ForecastDay[index].Day.MinTemp

	date, _ := time.Parse("2006-01-02", weather.Forecast.ForecastDay[index].Date)

	return fmt.Sprintf("<span foreground=\"cyan\">%s:</span> %s, %.0f/%.0f°C",
		date.Format("Monday"),
		condition,
		maxTemp,
		minTemp)

}