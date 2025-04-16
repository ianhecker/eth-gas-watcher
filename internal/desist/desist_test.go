package desist_test

import (
	"errors"
	"testing"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
	"github.com/stretchr/testify/assert"
)

func TestDesistor_WithError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := errors.New("expected")

		d := desist.NewDesistor()
		d.WithError(expected)

		assert.ErrorIs(t, d.Error(), expected)
	})

	t.Run("happy path - error is nil", func(t *testing.T) {
		d := desist.NewDesistor()
		assert.ErrorIs(t, d.Error(), nil)

		d.WithError(nil)
		assert.ErrorIs(t, d.Error(), nil)
	})
}

func TestDesistor_FatalOnError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := errors.New("expected")

		var called bool
		fatalFunc := func(v ...any) {
			called = true
		}

		d := desist.NewDesistorFromRaw(fatalFunc)
		d.WithError(expected).FatalOnError("message")

		assert.True(t, called)
	})

	t.Run("happy path - error is nil", func(t *testing.T) {
		var called bool
		fatalFunc := func(v ...any) {
			called = true
		}

		d := desist.NewDesistorFromRaw(fatalFunc)
		d.WithError(nil).FatalOnError("message")

		assert.False(t, called)
		assert.Nil(t, d.Error())
	})
}
