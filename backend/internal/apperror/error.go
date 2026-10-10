package apperror

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code       Code
	Message    string
	HTTPStatus int
	Err        error
}

// error實作go的error interface
func (e *AppError) Error() string {
	return e.Message
}

// 保留底層錯誤，讓errors.Is/errors.As可用
func (e *AppError) Unwrap() error {
	return e.Err
}

// 根據error code建立AppError
func NewWithCode(code Code, err error) *AppError {
	def, ok := definitions[code]

	if !ok {
		err = fmt.Errorf("unknown error code %q: %w", code, err)
		code = ErrInternalServer
		def = definitions[ErrInternalServer]
	}

	return &AppError{
		Code:       code,
		Message:    def.Message,
		HTTPStatus: def.HTTPStatus,
		Err:        err,
	}
}

// 把一般error轉換成AppError
func FromError(err error) *AppError {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr
	}

	return NewWithCode(ErrInternalServer, err)
}
