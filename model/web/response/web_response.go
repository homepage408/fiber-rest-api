package response

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// "status":  msg,
// 		"success": false,
// 		"message": err.Error(),
// 		"data":    d,
