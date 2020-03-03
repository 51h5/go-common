package amqp

import (
    "github.com/streadway/amqp"
)

// EncodeRequestFunc encodes the passed request object into
// an AMQP Publishing object. It is designed to be used in AMQP Publishers.
type EncodeRequestFunc func(*amqp.Publishing, interface{}) error