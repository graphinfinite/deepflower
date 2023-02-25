package repository

type ErrStoreUserNotFound struct {
	msg string
	err error
}

func NewErrUserNotFound(msg string, err error) *ErrStoreUserNotFound {
	return &ErrStoreUserNotFound{msg: msg, err: err}
}

func (err ErrStoreUserNotFound) Error() string {
	return err.msg
}
func (err ErrStoreUserNotFound) Unwrap() error {
	return err.err
}
