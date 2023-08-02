package clienterrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	clienterrorspb "github.com/requiemofthesouls/client-errors/pb"
)

var (
	ErrInternalServer  = status.Error(CodeInternalServer, MsgInternalServer)
	CodeInternalServer = codes.Internal // code = 13, httpStatus = 400
	MsgInternalServer  = "internal server error"

	ErrNotAllowed  = status.Error(CodeNotAllowed, MsgNotAllowed)
	CodeNotAllowed = codes.PermissionDenied // code = 7, httpStatus = 403
	MsgNotAllowed  = "not allowed"

	ErrValidation  = status.Error(CodeValidation, MsgValidation)
	CodeValidation = codes.InvalidArgument // code = 3, httpStatus = 400
	MsgValidation  = "validation error"

	ErrServiceUnavailable  = status.Error(CodeServiceUnavailable, MsgServiceUnavailable)
	CodeServiceUnavailable = codes.Unavailable // code = 14, httpStatus = 503
	MsgServiceUnavailable  = "the service is currently unavailable"

	ErrMetadataNotSent  = status.Error(CodeMetadataNotSent, MsgMetadataNotSent)
	CodeMetadataNotSent = codes.InvalidArgument // code = 3, httpStatus = 400
	MsgMetadataNotSent  = "metadata not sent"

	ErrPreconditionFailed  = status.Error(CodePreconditionFailed, MsgPreconditionFailed)
	CodePreconditionFailed = codes.FailedPrecondition // code = 9, httpStatus = 412
	MsgPreconditionFailed  = "precondition failed"

	ErrNotFound  = status.Error(CodeNotFound, MsgNotFound)
	CodeNotFound = codes.NotFound // code = 5, httpStatus = 404
	MsgNotFound  = "not found"
)

// NewValidationError creates new status error with codes.InvalidArgument code
func NewValidationError(opts ...OptionalErrorDetails) error {
	return newErrWithDetails(ErrValidation, opts...)
}

// NewNotAllowedError creates new status error with codes.PermissionDenied
func NewNotAllowedError(opts ...OptionalErrorDetails) error {
	return newErrWithDetails(ErrNotAllowed, opts...)
}

// NewNotFoundError creates new status error with codes.NotFound
func NewNotFoundError(opts ...OptionalErrorDetails) error {
	return newErrWithDetails(ErrNotFound, opts...)
}

func newErrWithDetails(parentErr error, opts ...OptionalErrorDetails) error {
	sErr, _ := status.FromError(parentErr)

	details := &clienterrorspb.ErrorDetails{}
	for _, opt := range opts {
		opt.apply(details)
	}

	if sErr, err := sErr.WithDetails(details); err == nil {
		return sErr.Err()
	}

	return sErr.Err()
}
