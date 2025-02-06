package my_error

type DeliveryError struct {
	error
	Method string
}

func NewDeliveryError(e error, method string) *DeliveryError {
	return &DeliveryError{
		error:  e,
		Method: method,
	}
}
