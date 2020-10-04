package events

import "errors"

var EventNotFound = errors.New("not found event")

var MethodNotFunc = errors.New("event method not is func")

var MissingMethodArgs = errors.New("missing args in call event method")