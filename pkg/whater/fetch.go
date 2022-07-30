package wheater

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchAPI(apiKey string, api string, method string, city string, forecast bool) (interface{}, error) {
	// godotenv.Load()
	// apiKey := os.Getenv("API_KEY_WHATERAPI")
	var WheaterResponse WeatherResponse

	// http://api.weatherapi.com/v1/forecast.json?key=361cd9ebacfd4d3993d4061922-2107&q=London&days=1&aqi=yes&alerts=yes
	var url string
	if forecast {
		url = fmt.Sprintf("http://api.weatherapi.com/%s?key=%s&q=%s&days=1&aqi=yes&alerts=yes", api, apiKey, city)
	} else {
		url = fmt.Sprintf("http://api.weatherapi.com/%s?key=%s&q=%s", api, apiKey, city)
	}
	fmt.Println("Url di fetch => ", url)
	fmt.Println(apiKey)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("REQUEST ", request)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&WheaterResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println()

	return WheaterResponse, nil

}
