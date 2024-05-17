package db

type ErrCode uint32

const (
	ErrInternal ErrCode = iota
	ErrNotFound
	ErrAlreadyExists
)

type Error struct {
	Code    ErrCode
	Message string
}

func (s Error) Error() string {
	return s.Message
}

func NewError(code ErrCode, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
