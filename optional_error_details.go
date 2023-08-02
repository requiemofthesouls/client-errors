package clienterrors

import (
	clienterrorspb "github.com/requiemofthesouls/client-errors/pb"
)

// OptionalErrorDetails optional parameters for error details
type OptionalErrorDetails interface {
	apply(*clienterrorspb.ErrorDetails)
}

func WithMessage(msg string) OptionalErrorDetails {
	return &withMessage{msg: msg}
}

type withMessage struct {
	msg string
}

func (w *withMessage) apply(details *clienterrorspb.ErrorDetails) {
	details.Msg = w.msg
}

// WithParam creates new ErrorDetailsOption with params
func WithParam(key, value string) OptionalErrorDetails {
	return &withParam{key: key, value: value}
}

type withParam struct {
	key, value string
}

func (w *withParam) apply(details *clienterrorspb.ErrorDetails) {
	if details.Params == nil {
		details.Params = make(map[string]string, 1)
	}

	details.Params[w.key] = w.value
}

func WithFieldViolation(field, msg string) OptionalErrorDetails {
	return &withFieldViolation{
		field: field,
		msg:   msg,
	}
}

type withFieldViolation struct {
	field string
	msg   string
}

func (w *withFieldViolation) apply(details *clienterrorspb.ErrorDetails) {
	if details.ValidationErrors == nil {
		details.ValidationErrors = &clienterrorspb.ErrorDetails_ValidationErrors{
			Items: make([]*clienterrorspb.ErrorDetails_ValidationErrorItem, 0),
		}
	}

	details.ValidationErrors.Items = append(details.ValidationErrors.Items, &clienterrorspb.ErrorDetails_ValidationErrorItem{
		Field: w.field,
		Msg:   w.msg,
	})
}
