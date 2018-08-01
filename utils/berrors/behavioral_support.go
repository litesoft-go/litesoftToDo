package berrors

func GetCause(err error) error {
	type causer interface {
		Cause() error
	}

	zCauser, ok := err.(causer)
	if ok {
		return zCauser.Cause()
	}
	return nil
}
