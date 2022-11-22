package wrappingerror_test

import (
	"errors"
	"testing"

	"github.com/backdround/go-wrappingerror"
	"github.com/stretchr/testify/assert"
)


func TestWrappedErrorMessage(t *testing.T) {
	t.Run("Without wrapped error", func(t *testing.T) {
		message := "Wrapping error"
		e := wrappingerror.NewWrappingError(message)
		assert.Equal(t, message, e.Error())
	})

	t.Run("With wrapped error", func(t *testing.T) {
		e := wrappingerror.NewWrappingError("external")
		e = e.Wrap(errors.New("internal"))

		expectedMessage := "external:\n  internal"
		assert.Equal(t, expectedMessage, e.Error())
	})
}

func TestWrappingErrorWrapAndUnwrap(t *testing.T) {
	e := wrappingerror.NewWrappingError("wrapping error")

	internalError := errors.New("internal error")
	e = e.Wrap(internalError)

	assert.Equal(t, internalError.Error(), e.Unwrap().Error())
}

func TestWrappingErrorIs(t *testing.T) {
	t.Run("WrappingError type", func(t *testing.T) {
		t.Run("Same value", func(t *testing.T) {
			initialError := wrappingerror.NewWrappingError("initial error")
			wrappingError := initialError.Wrap(errors.New("internal error"))

			assert.True(t, errors.Is(wrappingError, initialError),
				"they are the same errors. Must be true!")
		})

		t.Run("Another value", func(t *testing.T) {
			firstError := wrappingerror.NewWrappingError("wrapping error")
			secondError := wrappingerror.NewWrappingError("wrapping error")

			assert.False(t, errors.Is(firstError, secondError),
				"they are the different errors. Must be false!")
		})
	})

	t.Run("Not WrappedError type", func(t *testing.T) {
		wrappedError := wrappingerror.NewWrappingError("wrapping error")
		simpleError := errors.New("some error")
		assert.False(t, errors.Is(wrappedError, simpleError),
			"they are the different error. Must be false!")
	})
}

