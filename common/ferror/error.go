package ferror

import (
	"errors"
	"fmt"

	"github.com/openkaze/flux/common"
)

// Cause wraps an existing error with a custom message, preserving the cause.
func Cause(err error, message ...any) error {
	if err == nil {
		panic("cause on a nil error")
	}

	return fmt.Errorf("%s: %w", fmt.Sprint(message...), err)
}

// WithContext returns a new error with additional context message(s).
func WithContext(cause error, message ...any) error {
	if cause == nil {
		panic("extend on a nil error")
	}

	return fmt.Errorf("%s: %w", fmt.Sprint(message...), cause)
}

func Errors(errs ...error) error {
	errs = common.FilterNotNil(errs)
	errs = common.UniqBy(errs, error.Error)
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return errors.Join(errs...)
}
