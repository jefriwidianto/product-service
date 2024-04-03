package Response

type Responses struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}
