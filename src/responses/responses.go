package responses

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func GenerateResponse(s int, msg string, data map[string]interface{}) Response {
	return Response{Status: s, Message: msg, Data: data}
}
