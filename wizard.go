package cliwizard

import (
	"fmt"

	"github.com/PhysicalGraph/local-cloud/shared/ioutils"
)

// AskFunc is used to handle prompting a question and receiving the user input.
type AskFunc func(q *Question) Context

func defaultAsk(q *Question) Context {
	text := q.Text
	if q.Default != "" {
		text = fmt.Sprintf("%s [%s]", text, q.Default)
	}
	fmt.Printf("%s > ", text)
	c := &context{}
	{
		v := ioutils.MustReadIn()
		if v == "" {
			v = q.Default
		}
		c.SetRaw(v)
	}
	return c
}

// Engine is the main cliwizard interface.
//
// Ask is a generic question/response prompt.
type Engine interface {
	Ask(q *Question) error
}

// StandardEngine is the default implementation of the Engine interface. It
// will auto-render the help text of a question when "help" is entered as any
// answer value. Further, it will continue to re-prompt for input if the
// Callback does not return a nil error.
type StandardEngine struct {
	askFn AskFunc
}

// New creates a new StandardEngine wizard.
func New() Engine {
	return &StandardEngine{askFn: defaultAsk}
}

// Ask will infinitely prompt a user for an answer until the Callback returns
// a nil error.
func (e *StandardEngine) Ask(q *Question) error {
	if q.Callback == nil {
		return fmt.Errorf("a callback must be defined for question: %s", q.Text)
	}
	for {
		c := e.askFn(q)
		if c.Raw() == "help" {
			fmt.Println(q.Help)
			continue
		}
		if err := q.Callback(c); err != nil {
			fmt.Println(err.Error())
			continue
		}
		return nil
	}
}
