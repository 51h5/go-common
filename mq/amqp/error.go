package amqp

import "errors"

var (
    ErrNotConnected          = errors.New("not connected")
    ErrAlreadyClosed         = errors.New("already closed")
    ErrPublishConfirmFail    = errors.New("publish confirm fail")
    ErrPublishRequestInvalid = errors.New("publish request invalid")
)
