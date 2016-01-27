package authorize

import "errors"

var (
	INVALID_CONTENT       = errors.New("authorize: request content is invalid")
	INVALID_VALUE_LENGTH  = errors.New("authorize: invalid length of field")
	DUPLICATE_TRANSACTION = errors.New("authorize: duplicate transaction. Wait a few seconds.")
	DUPLICATE_RECORD      = errors.New("authorize: duplicate record")
)

func parseError(e *Error) error {
	if e.Messages == nil {
		return nil
	}

	var err error
	if e.TransactionResponse != nil {
		if len(e.TransactionResponse.Errors) > 0 {
			err = errors.New(e.TransactionResponse.Errors[0].ErrorText + ", with code " + e.Messages.Messages[0].Code)
		}
	}

	/*
		errGet, ok := errMap[e.Code]
		if !ok {
			return errors.New(e.Text + ", with code " + e.Code)
		}
	*/

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
