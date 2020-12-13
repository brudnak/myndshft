package web

// ErrorResponse is used to respond to clients when something goes wrong.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error adds web info to a request error.
type Error struct {
	Err    error
	Status int
}

func NewRequestError(err error, status int) error {
	return &Error{Err: err, Status: status}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
