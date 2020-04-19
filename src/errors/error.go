package errors

import (
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   http.StatusText(http.StatusInternalServerError),
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   http.StatusText(http.StatusNotFound),
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   http.StatusText(http.StatusBadRequest),
	}
}

func ParsePostgresError(err error) *RestErr {
	pgErr, ok := err.(*pq.Error) // error converted into postgres error
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewNotFoundError("No record matching the given id")
		}
		return NewInternalServerError(
			fmt.Sprintf("Error when parsing database: %s", err.Error()),
		)
	}

	switch pgErr.Code {
	case "23505":
		return NewInternalServerError(
			fmt.Sprintf("Error when saving: %s value already in use", pgErr.Constraint),
		)
	}

	return NewInternalServerError(
		fmt.Sprintf("Error processing request: %s", err.Error()),
	)
}
