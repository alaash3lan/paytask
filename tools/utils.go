package tools

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Message interface{} `json:"data,omitempty"`
}

func Respond() {

}