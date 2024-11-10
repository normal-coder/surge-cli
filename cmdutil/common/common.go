package common

type RequestBody struct {
	Method  string
	Url     string
	Payload string
}

type Response struct {
	Request    RequestBody
	Status     string
	StatusCode int
	Context    string
	Error      error
}

type ResponseError struct {
	ErrorStr string `json:"error"`
}
