package try

type Method func() error

type Exception interface {
	Catch(fn func(ex Exception))
	Error() error
}

type DefaultException struct {
	err error
}

func (d *DefaultException) Catch(fn func(ex Exception)) {

}

func (d *DefaultException) Error() error {
	return d.err
}

func Do(m Method) Exception {
	var exception DefaultException
	err := m()
	if err != nil {
		exception.err = err
	}
	return &exception
}

func Catch(ex Exception) error {
	return nil
}
