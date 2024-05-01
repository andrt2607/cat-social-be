package responsedto

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
