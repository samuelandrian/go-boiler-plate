package errorhelper

var (
	ErrorInvalidRequest error = ErrorStruct{Code: 400, Message: "Bad Request"}
)

type ErrorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements error.
func (e ErrorStruct) Error() string {
	return e.Message
}
