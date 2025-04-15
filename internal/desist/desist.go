package desist

import "fmt"

func Error(msg string, a any) error {
	return fmt.Errorf("%s: '%v'", msg, a)
}
