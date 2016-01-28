# cliwizard

[![Build Status](https://travis-ci.org/robzienert/cliwizard.svg?branch=master)](https://travis-ci.org/robzienert/cliwizard)

A simple library for making CLI wizards in Go.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/robzienert/cliwizard"
)

func main() {
	var favoriteColor string

	w := cliwizard.New()
	w.Ask(cliwizard.Q("What is your favorite color?").
		WithHelp(`
      We're asking this question first because it's the most important.
    `).
		WithCallback(
		func(c cliwizard.Context) error {
			answer := strings.ToLower(c.Raw())
			if answer == "black" || answer == "white" {
				return fmt.Errorf("Pick a new color: %s is a shade.", answer)
			}
			favoriteColor = answer
			return nil
		}))

	fmt.Printf("Your favorite color is: %s\n", favoriteColor)
}
```
