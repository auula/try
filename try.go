package try

import "fmt"

type Method func() error

type Exception interface {
	Catch()
	Error() error
	Try(err error)
	Catching() error
}

type DefaultException struct {
	Err error
}

func (d *DefaultException) Try(err error) {
	panic("implement me")
}

func (d *DefaultException) Catching() error {
	fmt.Println(d.Error())
	return nil
}

func (d *DefaultException) Catch() {
	_ = d.Catching()
}

func (d *DefaultException) Error() error {
	return d.Err
}

func Do(m Method, ex Exception) Exception {
	ex.Try(m())
	return ex
}
