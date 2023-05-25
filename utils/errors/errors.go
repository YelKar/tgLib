package errors

import "fmt"

const (
	LOW = iota
	MIDDLE
	HIGH
)

type _error struct {
	Code  int
	Text  string
	Level int8
	error
}

func (e *_error) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Text)
}

type Error *_error

var NoUpdates = &_error{
	Code:  1,
	Text:  "no updates",
	Level: LOW,
}

var WebhookIsActive = &_error{
	Code: 409,
	Text: "can't use getUpdates method while webhook is active\n" +
		"\tUse Bot.DeleteWebhook() to delete the webhook",
	Level: HIGH,
}
