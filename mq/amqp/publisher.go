package amqp

import (
    "github.com/streadway/amqp"
    "time"
)

type PublishRequest struct {
    Exchange, Key        string
    Mandatory, Immediate bool
    Body                 []byte
}

type Publisher interface {
    Publish(*PublishRequest) error
    UnsafePublish(*PublishRequest) error
    Close() error
}

type publisher struct {
    enc            EncodeRequestFunc
    confirm        bool
    reconnectDelay time.Duration
    resendDelay    time.Duration
    resendTime     int

    conn          *amqp.Connection
    ch            *amqp.Channel
    isConnected   bool
    done          chan bool
    notifyClose   chan *amqp.Error
    notifyConfirm chan amqp.Confirmation
}

type PublisherOption func(*publisher)

func PublisherConfirm(c bool) PublisherOption {
    return func(p *publisher) { p.confirm = c }
}

func PublisherReconnectDelay(t time.Duration) PublisherOption {
    return func(p *publisher) { p.reconnectDelay = t }
}

func PublisherResendDelay(t time.Duration) PublisherOption {
    return func(p *publisher) { p.resendDelay = t }
}

func PublisherResendTime(t int) PublisherOption {
    return func(p *publisher) { p.resendTime = t }
}

func NewPublisher(addr string, enc EncodeRequestFunc, options ...PublisherOption) Publisher {
    p := &publisher{
        enc:            enc,
        confirm:        false,
        reconnectDelay: reconnectDelay,
        resendDelay:    resendDelay,
        resendTime:     resendTime,
        done:           make(chan bool),
    }

    for _, option := range options {
        option(p)
    }

    go p.handleReconnect(addr)
    // go p.handleConfirm()

    return p
}

// func (p *publisher) handleConfirm() {
//     if !p.confirm {
//         return
//     }
//
//     for {
//         if p.notifyConfirm == nil {
//             continue
//         }
//
//         select {
//         case c := <-p.notifyConfirm:
//             if c.Ack {
//                 p.logger.Printf("消息发送确认成功: tag=%d, ack=%v", c.DeliveryTag, c.Ack)
//             }
//         }
//     }
// }

func (p *publisher) handleReconnect(addr string) {
    for {
        if !p.isConnected {
            for !p.connect(addr) {
                time.Sleep(p.reconnectDelay)
            }
        }

        select {
        case <-p.done:
            return
        case <-p.notifyClose:
            p.ch.Close()
            p.conn.Close()
            p.isConnected = false
        }
    }
}

func (p *publisher) connect(addr string) bool {
    var conn *amqp.Connection
    var ch *amqp.Channel
    var err error

    if p.conn == nil || p.conn.IsClosed() {
        conn, err = amqp.Dial(addr)
        if err != nil {
            return false
        }
    } else {
        conn = p.conn
    }

    ch, err = conn.Channel()
    if err != nil {
        return false
    }

    if p.confirm {
        ch.Confirm(false)
    }

    p.changeConnection(conn, ch)
    p.isConnected = true

    return true
}

func (p *publisher) changeConnection(conn *amqp.Connection, ch *amqp.Channel) {
    p.conn = conn
    p.ch = ch

    p.notifyClose = make(chan *amqp.Error)
    p.ch.NotifyClose(p.notifyClose)

    if p.confirm {
        p.notifyConfirm = make(chan amqp.Confirmation, 1)
        p.ch.NotifyPublish(p.notifyConfirm)
    }
}

func (p *publisher) Publish(req *PublishRequest) error {
    if !p.isConnected {
        return ErrNotConnected
    }

    if err := p.UnsafePublish(req); err != nil {
        return err
    }

    if !p.confirm {
        return nil
    }

    if c := <-p.notifyConfirm; c.Ack {
        return nil
    }

    return ErrPublishConfirmFail
}

func (p *publisher) UnsafePublish(req *PublishRequest) error {
    if req == nil {
        return ErrPublishRequestInvalid
    }

    if !p.isConnected {
        return ErrNotConnected
    }

    pub := amqp.Publishing{
        // DeliveryMode: amqp.Transient,
        Body: req.Body,
    }

    if p.enc != nil {
        _ = p.enc(&pub, req)
    }

    return p.ch.Publish(req.Exchange, req.Key, req.Mandatory, req.Immediate, pub)
}

func (p *publisher) Close() error {
    if !p.isConnected {
        return ErrAlreadyClosed
    }

    close(p.done)

    err := p.ch.Close()
    if err != nil {
        return err
    }

    err = p.conn.Close()
    if err != nil {
        return err
    }

    p.isConnected = false

    return nil
}
