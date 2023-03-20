package usecase

type ErrAuthUserAlreadyExist struct {
	msg string
	err error
}

func NewErrAuthUserAlreadyExist(msg string, err error) *ErrAuthUserAlreadyExist {
	return &ErrAuthUserAlreadyExist{msg: msg, err: err}
}

func (err ErrAuthUserAlreadyExist) Error() string {
	return err.msg
}
func (err ErrAuthUserAlreadyExist) Unwrap() error {
	return err.err
}
