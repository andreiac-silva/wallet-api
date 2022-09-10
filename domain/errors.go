package domain

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrIncompatibleProjectionVersion struct {
	Message string
}

func (e ErrIncompatibleProjectionVersion) Error() string {
	return e.Message
}

type ErrUnprocessable struct {
	Message string
}

func (e ErrUnprocessable) Error() string {
	return e.Message
}

type ErrInsufficientAmount struct{}

func (e ErrInsufficientAmount) Error() string {
	return "The given amount could not be debited from the given wallet because there aren't enough resources"
}
