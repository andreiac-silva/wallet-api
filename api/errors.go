package api

type ErrInvalidPayload struct {
}

func (e ErrInvalidPayload) Error() string {
	return "payload could not be parsed"
}

type ErrInvalidID struct{}

func (e ErrInvalidID) Error() string {
	return "invalid ID"
}

type ErrInvalidAttribute struct {
	Message string
}

func (e ErrInvalidAttribute) Error() string {
	return e.Message
}
