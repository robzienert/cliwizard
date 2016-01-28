package cliwizard

// Callback will be called after receiving an answer to a question, and should
// return a non-nil error when validation fails.
type Callback func(c Context) error

// Question defines an individual prompt for the user.
//
// It can optionally include a default value, as well as a help block.
type Question struct {
	Text     string
	Help     string
	Default  string
	Callback Callback
}

// Q is a convenience builder method for creating a new Question.
func Q(text string) *Question {
	return &Question{Text: text}
}

// WithHelp sets up the question to supply a help text block to the user if they
// ask for it.
func (q *Question) WithHelp(text string) *Question {
	q.Help = text
	return q
}

// WithDefault creates a default value if the user submits an empty value.
func (q *Question) WithDefault(text string) *Question {
	q.Default = text
	return q
}

// WithCallback will setup a function to be invoked immediately following input
// from a user.
func (q *Question) WithCallback(cb Callback) *Question {
	q.Callback = cb
	return q
}
