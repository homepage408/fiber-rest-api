package wheater

import "log"

type WheaterService interface {
	GetWheater(api string, method string, city string, forecast bool) (interface{}, error)
}

type wheaterService struct {
	apiKey string
}

func NewWheaterService(apiKey string) *wheaterService {
	return &wheaterService{apiKey: apiKey}
}

func (ws *wheaterService) GetWheater(api string, method string, city string, forecast bool) (interface{}, error) {
	response, err := FetchAPI(ws.apiKey, api, method, city, forecast)
	if err != nil {
		log.Fatal(err)
	}

	return response, err
}
