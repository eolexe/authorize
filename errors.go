package authorize

import (
	"errors"
	"log"
)

var (
	INVALID_CONTENT       = errors.New("authorize: request content is invalid")
	INVALID_VALUE_LENGTH  = errors.New("authorize: invalid length of field")
	DUPLICATE_TRANSACTION = errors.New("authorize: duplicate transaction. Wait a few seconds.")
	DUPLICATE_RECORD      = errors.New("authorize: duplicate record")
)

func parseError(e *Error) error {
	errGet, ok := errMap[e.Code]
	if !ok {
		log.Println("unknown error code", e.Code)
		return errors.New(e.Text)
	}

	err := errors.New(errGet.Error() + ", with code " + e.Code)

	return err
}

var (
	// maps error codes to errors
	errMap = map[string]error{
		"E00015": INVALID_VALUE_LENGTH,
		"E00013": INVALID_CONTENT,
		"E00027": DUPLICATE_TRANSACTION,
		"E00039": DUPLICATE_RECORD,
	}
)
