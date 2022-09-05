package api

type ErrInvalidPayload struct {
	Message string
}

func (e ErrInvalidPayload) Error() string {
	return e.Message
}

type ErrInvalidID struct {
	Message string
}

func (e ErrInvalidID) Error() string {
	return e.Message
}

type ErrInvalidAttribute struct {
	Message string
}

func (e ErrInvalidAttribute) Error() string {
	return e.Message
}
