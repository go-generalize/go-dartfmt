package dartfmt

type unavailableError struct {
	err error
}

var _ error = unavailableError{}

func (e unavailableError) Error() string {
	return e.err.Error()
}

func (e unavailableError) Unwrap() error {
	return e.err
}

func (e unavailableError) Is(target error) bool {
	return target == e.err
}
