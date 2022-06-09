package utils

import "fmt"

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code:%s, err:%s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

func TableUnavailableError(tableNo string) error {
	return AppError{
		ErrorCode:    "X01",
		ErrorMessage: fmt.Sprintf("Table %s is not available\n", tableNo),
		ErrorType:    0,
	}
}

func DataNotFoundError(info string) error {
	return AppError{
		ErrorCode:    "X02",
		ErrorMessage: fmt.Sprintf("Data [%s] not found\n", info),
		ErrorType:    0,
	}
}

func GeneralError(message string) error {
	return AppError{
		ErrorCode:    "X06",
		ErrorMessage: message,
		ErrorType:    0,
	}
}
