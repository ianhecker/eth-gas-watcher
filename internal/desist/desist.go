package desist

import (
	"fmt"
	"log"
)

func Error(msg string, a any) error {
	return fmt.Errorf("%s: '%v'", msg, a)
}

type DesistInterface interface {
	WithError(err error) DesistInterface
	FatalOnError(message string)
	Flush() DesistInterface
	Error() error
}

type Desistor struct {
	err   error
	fatal func(format string, v ...any)
}

func NewDesistor() DesistInterface {
	return &Desistor{
		fatal: log.Fatalf,
	}
}

func NewDesistorFromRaw(
	fatalFunc func(format string, v ...any),
) DesistInterface {
	return &Desistor{
		fatal: fatalFunc,
	}
}

func (d *Desistor) WithError(err error) DesistInterface {
	d.err = err
	return d
}

func (d *Desistor) FatalOnError(message string) {
	defer d.Flush()

	if d.err != nil {
		d.fatal("error: %s: %v", message, d.err)
	}
}

func (d *Desistor) Flush() DesistInterface {
	d.err = nil
	return d
}

func (d *Desistor) Error() error {
	return d.err
}
