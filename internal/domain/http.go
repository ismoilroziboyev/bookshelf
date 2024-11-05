package domain

type Response struct {
	Data    any    `json:"data"`
	IsOK    bool   `json:"isOk"`
	Message string `json:"message"`
}

type R map[string]interface{}
