package wheater

type InputReq struct {
	City     string `json:"city"`
	Forecast bool   `json:"forecast"`
	Method   string `json:"method"`
}
