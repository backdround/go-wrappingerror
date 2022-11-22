// wrappingerror provides error that can wrap other error and returns
// pretty error message.
package wrappingerror

import "sync"

////////////////////////////////////////////////////////
// wrappingErrorCounter

// wrappingErrorIdCounter provides new id by request
type wrappingErrorIdCounter struct {
	currentId int
	mutex     sync.Mutex
}

func (idCounter *wrappingErrorIdCounter) getId() int {
	idCounter.mutex.Lock()
	defer idCounter.mutex.Unlock()

	idCounter.currentId++
	return idCounter.currentId
}

var idCounter = wrappingErrorIdCounter{
	currentId: 0,
	mutex:     sync.Mutex{},
}

// //////////////////////////////////////////////////////
// wrappingError
type wrappingError struct {
	message       string
	internalError error
	id            int
}

// NewWrappingError creates new error with `message` value
func NewWrappingError(message string) wrappingError {
	return wrappingError{
		message:       message,
		internalError: nil,
		id:            idCounter.getId(),
	}
}

// Error generates error message with internal message as a
// clarification.
func (e wrappingError) Error() string {
	if e.internalError != nil {
		internalMessage := indent(e.internalError.Error(), "  ", 1)
		return e.message + ":\n" + internalMessage
	}

	return e.message
}

// Wrap wraps given error inside wrappingError.
func (e wrappingError) Wrap(err error) wrappingError {
	e.internalError = err
	return e
}

// Unwrap gives internal error or nil if it doesn't exist.
func (e wrappingError) Unwrap() error {
	return e.internalError
}

// Is implements interface for `errors.Is`
func (e wrappingError) Is(targetError error) bool {
	t, ok := targetError.(wrappingError)
	if !ok {
		return false
	}
	return e.id == t.id
}
