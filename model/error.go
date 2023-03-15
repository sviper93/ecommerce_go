package model

import "fmt"

type Error struct {
	Code       string
	Err        error
	Who        string
	StatusHTTP int
	Data       interface{}
	APIMessage string
	UserID     string
}

func NewError() Error {
	return Error{}
}

// Implementando la estructura error
func (e *Error) Error() string {
	return fmt.Sprintf("Código del error: %s, Error: %v, Who: %s, Status: %d, Data: %v, UserID: %s",
		e.Code,
		e.Err,
		e.Who,
		e.StatusHTTP,
		e.Data,
		e.UserID,
	)
}

func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) HasStatusHTTP() bool {
	return e.StatusHTTP > 0
}

func (e *Error) HasData() bool {
	return e.Data != nil
}
